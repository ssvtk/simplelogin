package serv

import (
	"net/http"
	"simplelogin/database"
)

//SignIn handles incomming logic as a handler function
func SignUp(w http.ResponseWriter, r *http.Request) {
	creds := ParseCredentials(w, r)
	hashedPassword := HashPassword(creds)
	database.InsertToDB(creds.Username, hashedPassword, w)
}

//SignUp ...
func SignIn() {

}
