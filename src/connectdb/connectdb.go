package connectdb

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db interface{}
var mTableName string
var mColums []string
var mInsertData []string

func opendb() bool {
	db, err := sql.Open("mysql", "root:@/my_database")
	checkErr(err)
}

func setTableName() {

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
