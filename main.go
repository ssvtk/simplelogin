package main

import (
	"log"
	"net/http"
	"simplelogin/serv"
)

func main() {
	http.HandleFunc("/signup", serv.SignUp)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
