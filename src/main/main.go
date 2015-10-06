package main

import (
	"fmt"
	_ "github.com/gorilla/mux"
	_ "html/template"
	"log"
	"login"
	"net/http"
	"signup"
	_ "strings"
)

func routing(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()
	//fmt.Println(r.Form)
	//fmt.Println("path", r.URL.Path)
	//fmt.Println("scheme", r.URL.Scheme)
	//fmt.Println(r.Form["url_long"])
	//for k, v := range r.Form {
	//	fmt.Println("kwy:", k)
	//	fmt.Println("val:", strings.Join(v, ""))
	//}
	fmt.Fprintf(w, "Hello go world!!!!")
}

/*
func login(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()
		if r.Method == "GET" {
			t, _ := template.ParseFiles("login.gtpl")
			w.Header().Set("Content-Type", "text/html")
			t.Execute(w, sess.Get("username"))
		} else {
			sess.Set("username", r.Form["username"])
			http.Redirect(w, r, "/", 302)
		}

}*/

func main() {
	//r := mux.NewRouter()
	http.HandleFunc("/", routing)
	http.HandleFunc("/login", login.Server)
	http.HandleFunc("/signup", signup.Server)
	err := http.ListenAndServe(":8989", nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
	//fmt.Fprintf(w, "hellow")
}
