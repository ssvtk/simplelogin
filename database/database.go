package database

import (
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx"
	"io/ioutil"
	"log"
)

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


