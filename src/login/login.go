package login

import (
	"D"
	"crypto/md5"

	"fmt"
	"net/http"
	"regexp"
	"text/template"
)

const (
	//login api for id
	loginId string = "Id"
	//login api for password
	loginPassword string = "Password"
)

func Server(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("login success")
	fmt.Println(r.Form)
	//check id that is match policy
	//The id allowed to use alphabet, capital letter and number
	//the number of chars is 8 <= id <= 16
	if m, _ := regexp.MatchString("^[a-zA-Z0-9]{8,16}$", r.Form.Get(loginId)); !m {
		fmt.Println("this id is not good")
	}
	//check Password that is match policy
	//The Password allowed to use alphabet, capital letter and number.
	//the number of chars is 8 <= passowrd <= 16
	if s, _ := regexp.MatchString("^[a-zA-Z0-9]{8,16}$", r.Form.Get(loginPassword)); !s {
		fmt.Println("this password is not good")
	} else {
		md5Pass := md5.Sum([]byte(template.HTMLEscapeString(r.Form.Get(loginPassword)) + D.PasswordHash))
		fmt.Println(md5Pass)
	}
	fmt.Println("login   ID: ", template.HTMLEscapeString(r.Form.Get(loginId)))
	fmt.Println("login PASS: ", template.HTMLEscapeString(r.Form.Get(loginPassword)))
	fmt.Println("login   ID: ", r.FormValue(loginId))
	fmt.Println("login PASS: ", r.FormValue(loginPassword))
	fmt.Fprintf(w, "login Id", r.FormValue(loginId))
	fmt.Fprintf(w, "Password", r.FormValue(loginPassword))
}
