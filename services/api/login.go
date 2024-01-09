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

	account, err := mysql.FetchUser(user.Email)
	if err != nil {
		log.Println("User not found for the user -", user.Email)
		resp := models.Response{Status: "Failed", Message: "User not Found!"}
		respJson, _ := json.Marshal(resp)
		w.WriteHeader(http.StatusAccepted)
		w.Write(respJson)
		return
	}

	salt, err := mysql.FetchSalt(account.SaltId)
	if err != nil {
		log.Println("Salt not found for the user -", user.Email)
		resp := models.Response{Status: "Failed", Message: "Unable to Verify"}
		respJson, _ := json.Marshal(resp)
		w.WriteHeader(http.StatusAccepted)
		w.Write(respJson)
		return
	}

	status, err := helpers.VerifySaltAndPassword(user.Password, salt.Salt, account.PasswordHash)
	if err != nil {
		log.Println("Unable to verify password for the user -", user.Email)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error!"))
		return
	}

	if !status {
		log.Println("Verification unsuccessful! for the user -", user.Email)
		resp := models.Response{Status: "Success", Message: "Unauthorized! Please use correct password"}
		respJson, _ := json.Marshal(resp)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(respJson)
		return
	}

	jwtToken, err := helpers.GenerateJwt(account.Id, account.Username, helpers.Config.Jwt.AuthSessionTime)
	if err != nil {
		log.Println("Error generating Auth Token")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error!"))
		return
	}

	cookie := http.Cookie{
		Name:    "auth",
		Value:   jwtToken,
		Path:    "/",
		Expires: time.Now().Add(time.Hour * time.Duration(helpers.Config.Jwt.AuthSessionTime)),
	}

	refreshToken, err := helpers.GenerateJwt(account.Id, account.Username, helpers.Config.Jwt.RefreshSessionTime)
	if err != nil {
		log.Println("Error generating Refresh Token")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error!"))
		return
	}

	refreshCookie := http.Cookie{
		Name:    "refresh",
		Value:   refreshToken,
		Path:    "/",
		Expires: time.Now().Add(time.Hour * time.Duration(helpers.Config.Jwt.RefreshSessionTime)),
	}

	log.Println("Successfully LoggedIn User - ", account.Username)

	http.SetCookie(w, &cookie)
	http.SetCookie(w, &refreshCookie)

	resp := models.Response{Status: "Success", Message: ""}
	respJson, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error converting response into json"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(respJson)
	return
}
