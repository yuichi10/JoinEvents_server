package signup

import (
	"D"
	"connectdb"
	"encoding/json"
	_ "errors"
	"fmt"
	_ "gopkg.in/validator.v2"
	_ "log"
	"net/http"
	"valid"
)

//<summary>sign up　の値を取り出すときに使うkey</summary>
const (
	httpId        string = "Id"
	httpPassword  string = "Password"
	httpPassword2 string = "Password2"
	httpName      string = "Name"
	httpGender    string = "Gender"
	httpAge       string = "Age"
	httpInfoErr   string = "InfoErr"
)

//the var name should be start from chapital letter for json
type checkError struct {
	IsErr     bool   `json:"isErr"`
	Id        string `json:"id_err"`
	Password  string `json:"pass1_err"`
	Password2 string `json:"pass2_err"`
	Name      string `json:"name_err"`
	Gender    string `json:"gender_err"`
	Age       string `json:"age_err"`
}

//<summary>とりあえず取得したあたいをこれにセット</summary>
var (
	id        string
	password1 string
	password2 string
	name      string
	gender    string
	age       string
)

//エラーのリストをここにいれる
var errList checkError

func insertData() {
	//insertPrepare := "INSERT INTO users SET id=? password=? name=? gender=? age=? update_time=?"
	//connectdb.Opendb()
	//stmt = connectdb.setInsertPrepare(insertPrepare)
}

func initErrorList() {

}

//エラーの一覧を表示　デバッグ用
func showErrorList() {
	fmt.Print(id + " -> ")
	fmt.Println(errList.Id)
	fmt.Print(password1 + " -> ")
	fmt.Println(errList.Password)
	fmt.Print(password2 + " -> ")
	fmt.Println(errList.Password2)
	fmt.Print(name + " -> ")
	fmt.Println(errList.Name)
	fmt.Print(gender + " -> ")
	fmt.Println(errList.Gender)
	fmt.Print(age + " -> ")
	fmt.Println(errList.Age)
}

//エラーが存在するか
func isErrorcheckErrlist() bool {
	if errList.Id != "" {
		return true
	} else if errList.Password != "" {
		return true
	} else if errList.Password2 != "" {
		return true
	} else if errList.Name != "" {
		return true
	} else if errList.Gender != "" {
		return true
	} else if errList.Age != "" {
		return true
	}
	return false
}

//<summary>実際にエラーがあるかどうか</summary>
func checkIsError() bool {
	errList.Id = valid.SetError(id, valid.Nonzero, valid.Min4, valid.Max12, valid.Ran)
	errList.Password = valid.SetError(password1, valid.Nonzero, valid.Min4, valid.Max12, valid.Ran)
	errList.Password2 = valid.SetError(password2, valid.Nonzero, valid.Min4, valid.Max12, valid.Ran)
	errList.Name = valid.SetError(name, valid.Nonzero, valid.Max18, valid.Ra)
	errList.Gender = valid.SetError(gender, valid.Rn)
	errList.Age = valid.SetError(age, valid.Rn, valid.Min2, valid.Max2)
	//パスワードが一緒かどうか
	if password1 != password2 {
		errList.Password2 = "Err : the password is not same"
	}
	//同じidが存在しないかどうか
	checkDoplicate := fmt.Sprintf("SELECT user.id FROM user where idUser=\"%v\"", id)
	if connectdb.IsDuplicate(checkDoplicate) {
		errList.Id = "There are duplicate"
	}
	errList.IsErr = isErrorcheckErrlist()
	return errList.IsErr
}

//送られたデータの取得
func getData(r *http.Request) {
	r.ParseForm()
	id = r.Form.Get(httpId)
	password1 = r.Form.Get(httpPassword)
	password2 = r.Form.Get(httpPassword2)
	name = r.Form.Get(httpName)
	gender = r.Form.Get(httpGender)
	age = r.Form.Get(httpAge)
}

func Server(w http.ResponseWriter, r *http.Request) {
	//データのセット
	getData(r)
	//エラーが有るかどうか
	is_err := checkIsError()
	//えらーあったらjson 形式でエラーを返す
	if is_err {
		jsonReturn, err := json.Marshal(&errList)
		if err != nil {
			fmt.Print("Can not make json data")
			return
		}
		fmt.Println("Json data")
		fmt.Println(string(jsonReturn))
		fmt.Fprintf(w, string(jsonReturn))
	} else {
		var md5Pass string = D.GetMd5Hash(password1)
		insertSentence := fmt.Sprintf("insert into user (idUser,password,name,gender,age) values(\"%s\",\"%s\",\"%s\",\"%s\",\"%s\")", id, md5Pass, name, gender, age)
		connectdb.RowInsert(insertSentence)
	}
}
