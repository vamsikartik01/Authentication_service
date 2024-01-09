package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/vamsikartik01/Authentication_service/api/models"
	"github.com/vamsikartik01/Authentication_service/api/services/helpers"
)

func Verify(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
		return
	}

	// resp1 := models.Response{Status: "Success", Message: "mars"}
	// respJson1, err := json.Marshal(resp1)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte("Internal Server Error converting response into json"))
	// 	return
	// }
	// w.WriteHeader(http.StatusCreated)
	// w.Write(respJson1)
	// return

	cookie, err := r.Cookie("auth")
	if err == http.ErrNoCookie {
		refreshCookie, err := r.Cookie("refresh")
		if err == http.ErrNoCookie {
			resp := models.Response{Status: "Failed", Message: "Unauthorized"}
			respJson, _ := json.Marshal(resp)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(respJson)
			return
		} else if err != nil {
			resp := models.Response{Status: "Failed", Message: "Internal Error!"}
			respJson, _ := json.Marshal(resp)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(respJson)
			return
		} else {
			claimsR, err := helpers.VerifyToken(refreshCookie.Value)
			if err != nil {
				resp := models.Response{Status: "Failed", Message: "Unauthorized, Invalid Token"}
				respJson, _ := json.Marshal(resp)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write(respJson)
				return
			}
			jwtToken, err := helpers.GenerateJwt(claimsR.UserId, claimsR.Username, helpers.Config.Jwt.AuthSessionTime)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal Server Error!"))
				return
			}

			redreshToken, err := helpers.GenerateJwt(claimsR.UserId, claimsR.Username, helpers.Config.Jwt.RefreshSessionTime)
			if err != nil {
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

			cookieR := http.Cookie{
				Name:    "refresh",
				Value:   redreshToken,
				Path:    "/",
				Expires: time.Now().Add(time.Hour * time.Duration(helpers.Config.Jwt.RefreshSessionTime)),
			}

			http.SetCookie(w, &cookie)
			http.SetCookie(w, &cookieR)
			log.Println("Successfully Verified User - ", claimsR.Username)

			resp := models.Response{Status: "Success", Message: claimsR.Username}
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

	} else if err != nil {
		resp := models.Response{Status: "Failed", Message: "Internal Error!"}
		respJson, _ := json.Marshal(resp)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(respJson)
		return
	}

	claims, err := helpers.VerifyToken(cookie.Value)
	if err != nil {
		resp := models.Response{Status: "Failed", Message: "Unauthorized, Invalid Token"}
		respJson, _ := json.Marshal(resp)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(respJson)
		return
	}

	log.Println("Successfully Verified User - ", claims.Username)

	resp := models.Response{Status: "Success", Message: claims.Username}
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
