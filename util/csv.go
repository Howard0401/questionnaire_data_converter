package util

import (
	"db_to_csv_go/model"
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

func OutputCSV(list []model.InfoDetail) {
	clientsFile, err := os.OpenFile("clients.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Panicf("clientsFile err=%v\n", list)
	}
	err = gocsv.MarshalFile(&list, clientsFile)
	if err != nil {
		log.Panicf("gocsv err=%v\n", list)
	}
}
