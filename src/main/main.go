package main

import (
	"connectdb"
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
	fmt.Fprintf(w, "Hello go world!!!!")
}

func main() {
	//r := mux.NewRouter()
	connectdb.Opendb()
	http.HandleFunc("/", routing)
	http.HandleFunc("/login", login.Server)
	http.HandleFunc("/signup", signup.Server)
	err := http.ListenAndServe(":8989", nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
