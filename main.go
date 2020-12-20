package main

import (
	"log"
	"net/http"
	"simplelogin/serv"
)

func main() {
	http.HandleFunc("/signup", serv.SignUp)
	http.HandleFunc("/signin", serv.SignIn)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
