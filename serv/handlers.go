package serv

import (
	"database/sql"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"simplelogin/database"
)

//SignUp handles incomming logic as a handler function
func SignUp(w http.ResponseWriter, r *http.Request) {
	creds := ParseCredentials(w, r)
	hashedPassword := HashPassword(creds)
	database.InsertToDB(creds.Username, hashedPassword, w)
}

//SignIn ...
func SignIn(w http.ResponseWriter, r *http.Request) {
	database.DB = database.Connect()
	creds := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result := database.DB.QueryRow("select password from users where username=$1", creds.Username)
	storedCreds := &Credentials{}

	err = result.Scan(&storedCreds.Password)
	if err != nil {

		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(storedCreds.Password), []byte(creds.Password)); err != nil {

		w.WriteHeader(http.StatusUnauthorized)
	}
	database.DB.Close()
}
