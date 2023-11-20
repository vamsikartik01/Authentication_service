package mysql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vamsikartik01/Authentication_service/api/services/helpers"
)

var db *sql.DB

func InitConnection() {
	const (
		host     = "localhost"
		port     = "3300"
		user     = "root"
		password = "password"
		dbname   = "jack_db"
	)

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", helpers.Config.Mysql.Username, helpers.Config.Mysql.Password, helpers.Config.Mysql.Host, helpers.Config.Mysql.Port, helpers.Config.Mysql.Database)

	var err error
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		log.Println("Connection to mysql failed with error : ", err)
	}

	if err = db.Ping(); err != nil {
		log.Println("Connection unsuccessfull with error :", err)
	}
	log.Println("Successfullt connected to Mysql db.")
}

func CloseConnection() {
	db.Close()
}
