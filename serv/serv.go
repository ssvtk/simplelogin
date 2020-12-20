package serv

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type Credentials struct {
	Username string `db:"username" ,json:"username"`
	Password string `db:"password" ,json:"password"`
}

//ParseCredentials parses and decodes the request body into a new `Credentials` instance
func ParseCredentials(w http.ResponseWriter, r *http.Request) *Credentials {
	creds := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err)
	}
	return creds
}

//HashPassword encrypt password with bcrypt algorythm and returns a hash value
func HashPassword(c *Credentials) []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(c.Password), 8)
	if err != nil {
		log.Fatal(err)
	}
	return hashedPassword
}
