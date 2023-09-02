package dao

import (
	"db_to_csv_go/model"
	"fmt"
	"log"
	"strings"
	"sync"
)

const queryLen = 500

type ResDetail struct {
	ResultSn   string
	FieldName  string
	FieldValue string
}

type DetailHelper struct {
	detailChan chan ResDetail
}

type QueryStat struct {
	F   func(int)
	Arg int
}

type WorkerPool struct {
	Jobs      chan QueryStat
	Wg        *sync.WaitGroup
	JobCounts int
	Helper    DetailHelper
	Done      chan bool
	Mu        *sync.Mutex
}

func (w *WorkerPool) PushJob(task QueryStat) {
	w.Jobs <- task
}

func (w *WorkerPool) Task(offset int) {
	queryString := fmt.Sprintf(`SELECT result_sn,
																		field_name,
																		field_value
																	FROM result_detail LIMIT %d OFFSET %d `, queryLen, offset)
	res, err := GetInstance().Queryx(queryString)
	if err != nil {
		log.Fatalf("result %v\n", err)
	}
	for res.Rows.Next() {
		var detail ResDetail
		res.Scan(&detail.ResultSn, &detail.FieldName, &detail.FieldValue)
		w.Helper.detailChan <- detail
	}
	log.Println("offset", offset, "querying...")
}

func (w *WorkerPool) Run() {
	for i := 0; i < w.JobCounts; i++ {
		w.Wg.Add(1)
		go func() {
			defer w.Wg.Done()
			select {
			case f, ok := <-w.Jobs:
				if !ok {
					close(w.Jobs)
					w.Done <- true
					break
				}
				f.F(f.Arg)
			default:
			}
		}()
	}
}

func (w *WorkerPool) NewWokerPool(total, chanCap int) WorkerPool {
	return WorkerPool{
		Jobs:      make(chan QueryStat, chanCap),
		Wg:        &sync.WaitGroup{},
		JobCounts: total,
		Helper: DetailHelper{
			detailChan: make(chan ResDetail, total),
		},
	}
}

func GetUsersMap(resultDetailLen int) map[string]any {

	usersMap := make(map[string]any)
	var w WorkerPool
	w = w.NewWokerPool(resultDetailLen, 50)
	go func() {
		w.Run()
	}()
	for i := 0; i < resultDetailLen/queryLen; i++ {
		w.PushJob(QueryStat{F: w.Task, Arg: i * queryLen})
	}
	w.Wg.Wait()

	for v := range w.Helper.detailChan {
		if len(w.Helper.detailChan) == 0 {
			break
		}
		if _, ok := usersMap[v.ResultSn]; !ok {
			inside := make(map[string]string)
			inside[v.FieldName] = v.FieldValue
			usersMap[v.ResultSn] = inside
		} else {
			u := usersMap[v.ResultSn].(map[string]string)
			u[v.FieldName] = v.FieldValue
			usersMap[v.ResultSn] = u
		}
	}
	return usersMap
}

func GetResultDetailLen() int {
	tx := GetInstance().QueryRowx(`SELECT
																		COUNT(*)
																		FROM result_detail`)
	var resultDetailLen int
	err := tx.Scan(&resultDetailLen)
	if err != nil {
		log.Panicf("get count err=%v\n", err)
	}
	return resultDetailLen
}

func InsertAnalysisDetail(list []model.InfoDetail, insertCols []string) {
	listLen := len(list)
	listCnt := 0
	const addPage = 50
	for listCnt < listLen {

		right := listCnt + addPage
		if right > listLen {
			right = listLen
		}
		insertStr := fmt.Sprintf(`INSERT INTO analysis_detail (%s)
															VALUES (%s) `, strings.Join(model.ColNames, ","), strings.TrimRight(strings.Join(insertCols, ","), ","))
		tx, err := GetInstance().Beginx()
		if err != nil {
			tx.Rollback()
			break
		}
		_, err = tx.NamedQuery(insertStr, list[listCnt:right])
		if err != nil {
			log.Fatalf("err=%v\n", err)
			tx.Rollback()
			break
		}
		err = tx.Commit()
		if err != nil {
			fmt.Printf("commit err=%v\n", err)
			tx.Rollback()
			break
		}
		log.Printf("listCnt=%v, right=%v done..\n", listCnt, right)
		listCnt = right
	}
}

func GetResultSnCount() ([]string, bool) {
	res, err := GetInstance().Queryx(`SELECT DISTINCT result_sn FROM result_detail GROUP BY result_sn`)
	if err != nil {
		log.Panicf("get result_sn count err=%v\n", err)
	}
	var distinctUserSn string
	resultSNList := make([]string, 10)
	for res.Rows.Next() {
		res.Scan(&distinctUserSn)
		resultSNList = append(resultSNList, distinctUserSn)
	}
	if len(resultSNList) == 0 {
		return nil, true
	}
	return resultSNList, false
}
