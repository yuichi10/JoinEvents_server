package connectdb

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
)

var mDb interface{}
var mTableName string
var mInsertPrepare interface{}
var mInsertData []string

func Opendb() bool {
	db, err := sql.Open("mysql", "yuichi:ys084743@/JoinEvents")
	ch := reflect.ValueOf(db)
	fmt.Printf(ch)
	checkErr(err)
	return false
}

func SetTableName(mode int) {

}

func setInsert(prepare string) bool {
	/*
		if mDb != nil {
			mInsertPrepare, err := mDb.Prepare(prepare)
			return *mInsertPrepare
		}
	*/
	return false
}

func SetColum() {

}

func SetInsertData(data ...string) {

}

func AddData() bool {
	return false
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
