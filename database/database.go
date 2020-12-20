package database

import (
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx"
	"io/ioutil"
	"log"
	"net/http"
)

var DB *pgx.Conn

//DBConfig unmarshalls the JSON file to pgx.ConnConfig struct and returns it.
func DBConfig() *pgx.ConnConfig {
	cfg := new(pgx.ConnConfig)
	buffer, err := ioutil.ReadFile("database/config.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(buffer, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}

//Connect performs a connection to DB according to DBConfig.
func Connect() *pgx.Conn {
	conn, err := pgx.Connect(*DBConfig())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to DB has been successfully done")
	return conn
}

//InsertToDB inserts login and password into the connected db
func InsertToDB(login string, password []byte, w http.ResponseWriter) {
	DB = Connect()
	_, err := DB.Query("INSERT INTO users values ($1, $2)", login, string(password))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	DB.Close()
}
