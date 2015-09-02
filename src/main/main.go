package main

import (
	"fmt"
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

func main() {
	http.HandleFunc("/", routing)
	http.HandleFunc("/login", login.Server)
	http.HandleFunc("/signup", signup.Server)
	err := http.ListenAndServe(":9898", nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
	//fmt.Fprintf(w, "hellow")
}
