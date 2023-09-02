package main

import (
	"db_to_csv_go/dao"
	"db_to_csv_go/service"
	"db_to_csv_go/util"
)

func main() {
	InitSetting()
	dao.InitDB()
	defer dao.SingletonDB.Close()

	// 1. 先查有幾個 user 有資料，以及取 userList
	resultSNList, shouldReturn := dao.GetResultSnCount()
	if shouldReturn {
		return
	}

	// 2. GetResultDetailLen 先拿總共有幾筆 result_detail
	// 3. GetUsersMap 分段查詢 result_detail 組成 usersMap(資料筆數少應該不會爆)
	usersMap := dao.GetUsersMap(dao.GetResultDetailLen())

	// 4. 準備插入資料
	// 5. 準備 insert analysis_detail 的 column
	list, insertCols := service.ProcessInsertColsAndInfos(resultSNList, usersMap)

	// 6. 實際插入的格式
	// 分頁最後一筆
	dao.InsertAnalysisDetail(list, insertCols)

	// ˙7. output csv file
	util.OutputCSV(list)
}

// 轉換工具
// func main2() {
// 	file, err := os.Open("./tt4.txt") //File("./tt.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fileScanner := bufio.NewScanner(file)
// 	fileScanner.Split(bufio.ScanLines)

// 	editSlice := make([]string, 10)
// 	for fileScanner.Scan() {
// 		// fmt.Println(fileScanner.Text())
// 		// fmt.Printf("fileScanner.Text()=%v\n", fileScanner.Text())
// 		// tt3 := strings.Split(fileScanner.Text(), ",")
// 		// editSlice = append(editSlice, strings.Trim(tt3[0], " "))
// 		// editSlice = append(editSlice, strings.Trim(fileScanner.Text(), " ")+",")
// 		str := strings.Trim(fileScanner.Text(), " ")
// 		// editSlice = append(editSlice, "`"+str+"` TEXT NOT NULL,")
// 		editSlice = append(editSlice, "\""+str+"\",")
// 		// editSlice = append(editSlice, "\""+strings.Trim(fileScanner.Text(), " ")+"\",")
// 	}
// 	// fmt.Printf("editSlice=%v\n", editSlice)

// 	wFile, _ := os.OpenFile("./tt4.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
// 	w := bufio.NewWriter(wFile)
// 	for _, data := range editSlice {
// 		_, err = w.WriteString(data + "\n")
// 		if err != nil {
// 			fmt.Printf("err=%v\n", err)
// 		}
// 	}
// 	w.Flush()
// 	file.Close()
// }
