package mysql

import (
	"fmt"

	"github.com/vamsikartik01/Authentication_service/api/models"
)

func InsertUser(user *models.Account, saltId int) error {
	query := "INSERT INTO Accounts(username, email, passwordHash, saltId) values(?, ?, ?, ?)"
	_, err := db.Exec(query, user.Username, user.Email, user.PasswordHash, saltId)
	if err != nil {
		fmt.Println("Error inserting Salt", err)
		return err
	}
	return nil
}

func InsertSalt(salt string) (int, error) {
	query := "INSERT INTO Salts(salt) VALUES(?)"
	result, err := db.Exec(query, salt)
	if err != nil {
		fmt.Println("Error inserting Salt", err)
		return 0, err
	}
	rowId, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Error Fetching Salt Id", err)
		return 0, err
	}
	return int(rowId), nil
}
