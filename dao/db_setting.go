package dao

import (
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

var SingletonDB *sqlx.DB

func GetInstance() *sqlx.DB {
	// if singletonDB != nil {
	return SingletonDB
	// }
	// InitDB()
	// return new(sqlx.DB)
}

func GetDBConnString() string {
	DB_ACCOUNT := viper.GetString("db.user_account")
	DB_PASSWORD := viper.GetString("db.user_password")
	DB_URI := viper.GetString("db.uri")
	DB_NAME := viper.GetString("db.db_name")
	rtn := fmt.Sprintf("%s:%s@(%s)/%s", DB_ACCOUNT, DB_PASSWORD, DB_URI, DB_NAME)
	// fmt.Printf("rtn=%v\n", rtn)
	if rtn == "" {
		log.Panicf("empty db conn string\n")
	}
	return rtn
}

func GetDBType() string {
	DB_TYPE := viper.GetString("db.type")
	switch DB_TYPE {
	case "mysql":
		DB_TYPE = "mysql"
	default:
		DB_TYPE = "mysql"
	}
	if DB_TYPE == "" {
		log.Panicf("empty DB_TYPE\n")
	}
	return DB_TYPE
}

func InitDB() {
	once := sync.Once{}
	var err error
	once.Do(func() {
		SingletonDB, err = sqlx.Connect(GetDBType(), GetDBConnString())
		if err != nil {
			log.Panicf(" init db failed=%v\n", err)
		}
		SingletonDB.SetMaxOpenConns(10)
	})
}
