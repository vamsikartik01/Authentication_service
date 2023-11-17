package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
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

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)

	var err error
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		fmt.Println("Connection to postgres failed with error : ", err)
	}

	if err = db.Ping(); err != nil {
		fmt.Println("Connection unsuccessfull with error :", err)
	}
	fmt.Println("Successfullt connected to Mysql db.")
}

func CloseConnection() {
	db.Close()
}
