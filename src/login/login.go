package login

import (
	"D"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
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
	//response sessionId
	resSessionId string = "SessionId"
	//response Err for password
	resErrPass string = "ErrPass"
	//response Err for id
	resErrId string = "ErrId"
)

func getMd5Hash(text string) string {
	hash := md5.New()
	hash.Write([]byte(text + D.HashKey))
	return hex.EncodeToString(hash.Sum(nil))
}

func Server(w http.ResponseWriter, r *http.Request) {
	//error message for password
	var errPass string = ""
	//error message for id
	var errId string = ""
	//md5 password and md5 id it for sessionID
	var encId string
	var encPass string

	r.ParseForm()
	fmt.Println(r.Form)
	//check id that is match policy
	//The id allowed to use alphabet, capital letter and number
	//the number of chars is 8 <= id <= 16
	if m, _ := regexp.MatchString("^[a-zA-Z0-9]{8,16}$", r.Form.Get(loginId)); !m {
		errId = "alphanumeric only. length is 8 to 16"
	} else {
		encId = getMd5Hash(template.HTMLEscapeString(r.Form.Get(loginId)))
	}
	//check Password that is match policy
	//The Password allowed to use alphabet, capital letter and number.
	//the number of chars is 8 <= passowrd <= 16
	if s, _ := regexp.MatchString("^[a-zA-Z0-9]{8,16}$", r.Form.Get(loginPassword)); !s {
		errPass = "alphanumeric only. length is 8 to 16"
		fmt.Println(errPass)
	} else {
		encPass = getMd5Hash(template.HTMLEscapeString(r.Form.Get(loginPassword)))
		fmt.Println(encPass)
	}

	//set response data as json
	resContents := make(map[string]string)
	//set err of id
	resContents[resErrId] = errId
	//set err of password
	resContents[resErrPass] = errPass
	//set sesion id
	resContents[resSessionId] = encId
	//make json structure
	jsonB, err := json.Marshal(resContents)
	if err != nil {
		fmt.Println("json err:", err)
	}
	//sent json data
	fmt.Fprintf(w, string(jsonB))
}
