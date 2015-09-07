package connectdb

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var mDb interface{}
var mTableName string
var mColums []string
var mInsertData []string

func opendb() bool {
	db, err := sql.Open("mysql", "userName:password@/databaseName")
	mDb = db
	checkErr(err)
}

func setTableName(mode int) {

}

func setColum(colums ...string) {
	mColums = colums
}

func setInsertData(data ...string) {
	mInsertData = data
}

func addData() bool {

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
