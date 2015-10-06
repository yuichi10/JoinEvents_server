package signup

import (
	"D"
	"encoding/json"
	_ "errors"
	"fmt"
	_ "gopkg.in/validator.v2"
	"net/http"
	"valid"
)

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

var (
	id        string
	password1 string
	password2 string
	name      string
	gender    string
	age       string
)

var errList checkError

func insertData() {
	//insertPrepare := "INSERT INTO users SET id=? password=? name=? gender=? age=? update_time=?"
	//connectdb.Opendb()
	//stmt = connectdb.setInsertPrepare(insertPrepare)
}

func initErrorList() {

}

func showErrorList() {
	fmt.Println(errList.IsErr)
	fmt.Println(errList.Id)
	fmt.Println(errList.Password)
	fmt.Println(errList.Password2)
	fmt.Println(errList.Name)
	fmt.Println(errList.Gender)
	fmt.Println(errList.Age)
}

func isErrorcheckErrlist() bool {
	if errList.Id == "" {
		return true
	} else if errList.Password == "" {
		return true
	} else if errList.Password2 == "" {
		return true
	} else if errList.Name == "" {
		return true
	} else if errList.Gender == "" {
		return true
	} else if errList.Age == "" {
		return true
	}
	return false
}

func checkIsError() bool {
	errList.Id = valid.SetError(id, valid.Nonzero, valid.Min4, valid.Max12, valid.Ran)
	errList.Password = valid.SetError(password1, valid.Nonzero, valid.Min4, valid.Max12, valid.Ran)
	errList.Password2 = valid.SetError(password2, valid.Nonzero, valid.Min4, valid.Max12, valid.Ran)
	errList.Name = valid.SetError(name, valid.Nonzero, valid.Max18, valid.Ra)
	errList.Gender = valid.SetError(gender, valid.Rn)
	errList.Age = valid.SetError(age, valid.Rn, valid.Min2, valid.Max2)
	if password1 != password2 {
		errList.Password2 = "Err : the password is not same"
	}
	errList.IsErr = isErrorcheckErrlist()
	return errList.IsErr
}

func Server(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id = r.Form.Get(httpId)
	password1 = r.Form.Get(httpPassword)
	password2 = r.Form.Get(httpPassword2)
	name = r.Form.Get(httpName)
	gender = r.Form.Get(httpGender)
	age = r.Form.Get(httpAge)

	var is_err = checkIsError()
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
	fmt.Println(is_err)
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
		fmt.Fprintf(w, "No ERROR!!!!"+md5Pass)
	}
}
