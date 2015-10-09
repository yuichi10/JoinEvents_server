package login

import (
	"D"
	"connectdb"
	"encoding/json"
	"fmt"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	"valid"
)

type errorList struct {
	IsErr          bool   `json:"is_err"`
	Id             string `json:"id_err"`
	Password       string `json:"pass_err"`
	ErrDescription string `json:"err_description"`
}

var mErrList errorList

var Store = sessions.NewCookieStore([]byte(D.S_Key))

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

func isErrorcheckErrlist() bool {
	if mErrList.Id != "" {
		return true
	} else if mErrList.Password != "" {
		return true
	}
	return false
}

func checkIsErr(r *http.Request) bool {
	//とりあえず正規表現チェック
	//check id that is match policy
	mErrList.Id = valid.SetError(r.Form.Get(loginId), valid.Nonzero, valid.Min4, valid.Max12, valid.Ran)
	//check Password that is match policy
	mErrList.Password = valid.SetError(r.Form.Get(loginPassword), valid.Nonzero, valid.Min4, valid.Max12, valid.Ran)
	//とりあえず正規表現エラーチェック
	if mErrList.IsErr = isErrorcheckErrlist(); mErrList.IsErr {
		mErrList.ErrDescription = "something wrong"
		return mErrList.IsErr
	}
	//存在チェック
	var md5Pass string = D.GetMd5Hash(r.Form.Get(loginPassword))
	checkDuplicate := fmt.Sprintf("SELECT user.id FROM user where idUser=\"%s\", password=\"%s\"", r.Form.Get(loginId), md5Pass)
	if mErrList.IsErr = !connectdb.IsDuplicate(checkDuplicate); mErrList.IsErr {
		log.Println("err %v", mErrList.IsErr)
		mErrList.ErrDescription = "something wrong"
	}
	return mErrList.IsErr
}

func Server(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	if checkIsErr(r) {
		jsonObj, err := json.Marshal(mErrList)
		if err != nil {
			fmt.Println("json err:", err)
		}
		//sent json data
		fmt.Fprintf(w, string(jsonObj))
	} else {
		session, err := Store.Get(r, r.Form.Get(loginId))
		if err != nil {
			//panic()
		}
		session.Save(r, w)
		log.Printf("%v", session)
	}
}
