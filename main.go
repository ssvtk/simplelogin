package main

import (
	"fmt"
	"simplelogin/database"
)

func main() {
	conn := database.Connect()
	fmt.Println(conn)
}
