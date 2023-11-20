package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/vamsikartik01/Authentication_service/api/models"
	"github.com/vamsikartik01/Authentication_service/api/services/helpers"
	"github.com/vamsikartik01/Authentication_service/api/services/mysql"
)

func SigninUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Methos not allowed!"))
		return
	}

	user := &models.SigninForm{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	account, err := mysql.FetchUser(user)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("User Not Found!"))
		return
	}

	salt, err := mysql.FetchSalt(account.SaltId)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("Salt Not Found!"))
		return
	}

	status, err := helpers.VerifySaltAndPassword(user.Password, salt.Salt, account.PasswordHash)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error!"))
		return
	}

	if !status {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized! Please use correct password"))
		return
	}

	jwtToken, err := helpers.GenerateJwt(account.Id, account.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error!"))
		return
	}

	cookie := http.Cookie{
		Name:    "auth",
		Value:   jwtToken,
		Expires: time.Now().Add(time.Hour * 2),
	}

	log.Println("Successfully LoggedIn User - ", account.Username)

	http.SetCookie(w, &cookie)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login Successfull!"))
	return
}
