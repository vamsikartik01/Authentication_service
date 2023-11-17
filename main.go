package main

import (
	"fmt"
	"net/http"

	"github.com/vamsikartik01/Authentication_service/api/services/api"
	"github.com/vamsikartik01/Authentication_service/api/services/mysql"
)

func main() {
	fmt.Println("Hello Auth!")

	mysql.InitConnection()
	defer mysql.CloseConnection()

	http.HandleFunc("/login", api.LoginUser)
	http.HandleFunc("/signup", api.SignupUser)

	err := http.ListenAndServe(":2200", nil)
	if err != nil {
		fmt.Println("closing")
	}
}
