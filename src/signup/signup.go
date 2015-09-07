package signup

import (
	"fmt"
	"net/http"
)

const (
	httpId        string = "Id"
	httpPassword  string = "Password"
	httpPassword2 string = "Password2"
	httpName      string = "Name"
	httpGender    string = "Gender"
)

func Server(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello sign up go")
	r.ParseForm()
	id := r.Form.Get(httpId)
	password1 := r.Form.Get(httpPassword)
	password2 := r.Form.Get(httpPassword2)
	name := r.Form.Get(httpName)
	gender := r.Form.Get(httpGender)
	fmt.Println(w, id)
	fmt.Println(w, password1)
	fmt.Println(w, password2)
	fmt.Println(w, name)
	fmt.Println(w, gender)
}
