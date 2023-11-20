package main

import (
	"log"
	"net/http"

	"github.com/vamsikartik01/Authentication_service/api/services/api"
	"github.com/vamsikartik01/Authentication_service/api/services/helpers"
	"github.com/vamsikartik01/Authentication_service/api/services/mysql"
)

func main() {
	log.Println("Hello Auth! API is listening at port :2200")

	err := helpers.LoadConfig()
	if err != nil {
		log.Println("Error Loading Config.", err)
	}

	mysql.InitConnection()
	defer mysql.CloseConnection()

	http.HandleFunc("/auth/v1/signin", api.SigninUser)
	http.HandleFunc("/auth/v1/signup", api.SignupUser)
	http.HandleFunc("/auth/v1/verify", api.Verify)

	err = http.ListenAndServe(":2200", nil)
	if err != nil {
		log.Println("closing with error ", err)
	}
}
