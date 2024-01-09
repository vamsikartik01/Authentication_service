package main

import (
	"log"
	"net/http"

	"github.com/vamsikartik01/Authentication_service/api/services/api"
	"github.com/vamsikartik01/Authentication_service/api/services/helpers"
	"github.com/vamsikartik01/Authentication_service/api/services/mysql"
)

func CorsMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	log.Println("Hello Auth! API is listening at port :2200")

	err := helpers.LoadConfig()
	if err != nil {
		log.Println("Error Loading Config.", err)
	}
	log.Println("Successfully Loaded Configs")

	mysql.InitConnection()
	defer mysql.CloseConnection()

	mux := http.NewServeMux()

	http.Handle("/", CorsMiddleWare(mux))

	mux.HandleFunc("/v1/signin", api.SigninUser)
	mux.HandleFunc("/v1/signup", api.SignupUser)
	mux.HandleFunc("/v1/verify", api.Verify)

	err = http.ListenAndServe(":2200", nil)
	if err != nil {
		log.Println("closing with error ", err)
	}
}
