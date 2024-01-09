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
		w.Write([]byte("Method Not Allowed!"))
		return
	}

	user := &models.SignupForm{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		log.Println("Error decoding request body")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Internal Server Error!"))
		return
	}

	account, err := mysql.FetchUser(user.Email)
	if err == nil {
		log.Println("A user Already exists with email - ", account.Email, "and username -", account.Username)
		resp := models.Response{Status: "Failed", Message: "User Already Exists!"}
		respJson, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusAlreadyReported)
			w.Write([]byte("Internal Server Error converting response into json"))
			return
		}
		w.WriteHeader(http.StatusAccepted)
		w.Write(respJson)
		return
	}

	salt, err := helpers.GenerateSalt(user.Password)
	if err != nil {
		log.Println("Unable to generate Salt for request - ", user.Email)
		resp := models.Response{Status: "Failed", Message: "Unable to SignUp!"}
		respJson, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error converting response into json"))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(respJson)
		return
	}

	hash, err := helpers.GeneratePasswordHash(user.Password, salt)
	if err != nil {
		log.Println("Unable to generate Password hash! for request - ", user.Email)
		resp := models.Response{Status: "Failed", Message: "Internal Server Error"}
		respJson, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error converting response into json"))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(respJson)
		return
	}

	saltId, err := mysql.InsertSalt(salt)
	if err != nil {
		resp := models.Response{Status: "Failed", Message: "Internal Server Error db error"}
		respJson, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error converting response into json"))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(respJson)
		return
	}

	account = &models.Account{Username: user.Username, PasswordHash: hash, Email: user.Email}
	err = mysql.InsertUser(account, saltId)
	if err != nil {
		resp := models.Response{Status: "Failed", Message: "Internal Server Error db error"}
		respJson, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error converting response into json"))
			return
		}
		w.WriteHeader(http.StatusAlreadyReported)
		w.Write(respJson)
		return
	}

	defer r.Body.Close()

	log.Println("Successfully Created User - ", user.Username)

	resp := models.Response{Status: "Success", Message: "Succesfully Created User Account"}
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
