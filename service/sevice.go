package service

import (
	"db_to_csv_go/model"
	"encoding/json"
	"fmt"
	"strconv"
)

func ProcessInsertColsAndInfos(resultSNList []string, usersMap map[string]any) ([]model.InfoDetail, []string) {
	list := make([]model.InfoDetail, 0)
	for _, v := range resultSNList {

		var detail model.InfoDetail

		jsonStr, err := json.Marshal(usersMap[v])
		if err != nil {
			fmt.Printf("marshal usersMap[v] err= %+v\n", err)
		}

		err = json.Unmarshal(jsonStr, &detail)
		if err != nil && v != "" {
			fmt.Printf("unmarshal usersMap[v] err= %+v\n", err)
			continue
		}
		resultSn, err := strconv.Atoi(v)
		if err != nil && v != "" {
			fmt.Printf("resultSn v=%v, err= %+v\n", v, err)
			continue
		}
		detail.ResultSN = resultSn
		if detail.ResultSN == 0 {
			continue
		}
		list = append(list, detail)
	}

	insertCols := make([]string, len(model.ColNames))
	copy(insertCols, model.ColNames)
	for key := range insertCols {
		insertCols[key] = ":" + insertCols[key]
	}
	return list, insertCols
}
