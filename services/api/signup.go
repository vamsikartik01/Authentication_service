package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/vamsikartik01/Authentication_service/api/models"
	"github.com/vamsikartik01/Authentication_service/api/services/helpers"
	"github.com/vamsikartik01/Authentication_service/api/services/mysql"
)

func SignupUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
		return
	}

	user := &models.SignupForm{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	salt, err := helpers.GenerateSalt(user.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	hash, err := helpers.GeneratePasswordHash(user.Password, salt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	saltId, err := mysql.InsertSalt(salt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error mysql error"))
		return
	}

	account := &models.Account{Username: user.Username, PasswordHash: hash, Email: user.Email}
	err = mysql.InsertUser(account, saltId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error mysql error"))
		return
	}

	defer r.Body.Close()

	log.Println("Successfully Created User - ", user.Username)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User Created Successfully!"))
	return
}
