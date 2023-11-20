package api

import (
	"log"
	"net/http"

	"github.com/vamsikartik01/Authentication_service/api/services/helpers"
)

func Verify(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
		return
	}

	cookie, err := r.Cookie("auth")
	if err == http.ErrNoCookie {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Please Login"))
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	// type VerifyForm struct {
	// 	Token string `json:"token"`
	// }

	// token := &VerifyForm{}
	// err = json.NewDecoder(r.Body).Decode(token)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte("Bad Request"))
	// 	return
	// }

	claims, err := helpers.VerifyToken(cookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid token!"))
		return
	}

	log.Println("Successfully Verified User - ", claims.Username)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(claims.Username))
	return
}
