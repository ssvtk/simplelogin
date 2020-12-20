package serv

import (
	"net/http"
	"simplelogin/database"
)

//SignIn ...
func SignUp(w http.ResponseWriter, r *http.Request) {
	creds := ParseCredentials(w, r)
	hashedPassword := HashPassword(creds)
	database.InsertToDB(creds.Username, hashedPassword, w)
}

//SignUp ...
func SignIn() {

}
