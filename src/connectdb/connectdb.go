package connectdb

/*
create table user (
id int not null AUTO_INCREMENT,
idUser varchar(255) not null,
password varchar(255) not null,
name varchar(255) not null,
gender tinyint not null,
age int not null,
PRIMARY KEY(id),
UNIQUE(idUser)
)

*/
import (
	"database/sql"
	_ "fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	_ "reflect"
)

var mDb *sql.DB
var mTableName string
var mInsertPrepare interface{}
var mInsertData []string

func Opendb() bool {
	db, err := sql.Open("mysql", "yuichi:ys084743@/JoinEvents")
	mDb = db

	checkErr(err)
	return false
}

func RowInsert(row string) {
	//insert into user value ()
	//insert into user (idUser, password, name, gender, age) values("idtest","passtest","nametest",0,99);
	log.Printf(row)

	result, err := mDb.Exec(row)
	if err != nil {
		log.Println("insert error", err)
		log.Println("insert result", result)
	}
}

func RowRead(row string) *sql.Rows {
	rows, qerr := mDb.Query(row)
	if qerr != nil {
		log.Fatal("query error: %v", qerr)
	}
	for rows.Next() {
		var id int
		if berr := rows.Scan(&id); berr != nil {
			log.Fatal("scan erro: %v", berr)
		}
		log.Println(id)
	}
	return rows
}

/*
	<summary>重複チェック</summary>
	<return>もし重複があったらtrue なかったらfalse</return>
*/
func IsDuplicate(row string) bool {
	rows, qerr := mDb.Query(row)
	if qerr != nil {
		return true
	}
	var count int = 0

	for rows.Next() {
		count++
		break
	}
	if count == 0 {
		return false
	} else {
		return true
	}
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
