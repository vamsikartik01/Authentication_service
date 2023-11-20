package mysql

import (
	"log"

	"github.com/vamsikartik01/Authentication_service/api/models"
)

func InsertUser(user *models.Account, saltId int) error {
	query := "INSERT INTO Accounts(username, email, passwordHash, saltId) values(?, ?, ?, ?)"
	_, err := db.Exec(query, user.Username, user.Email, user.PasswordHash, saltId)
	if err != nil {
		log.Println("Error inserting Salt", err)
		return err
	}
	return nil
}

func InsertSalt(salt string) (int, error) {
	query := "INSERT INTO Salts(salt) VALUES(?)"
	result, err := db.Exec(query, salt)
	if err != nil {
		log.Println("Error inserting Salt", err)
		return 0, err
	}
	rowId, err := result.LastInsertId()
	if err != nil {
		log.Println("Error Fetching Salt Id", err)
		return 0, err
	}
	return int(rowId), nil
}

func FetchUser(user *models.SigninForm) (*models.Account, error) {
	query := "SELECT * FROM Accounts WHERE email = ?"

	account := models.Account{}
	err := db.QueryRow(query, user.Email).Scan(&account.Id, &account.Username, &account.Email, &account.PasswordHash, &account.SaltId, &account.CreatedAt, &account.PasswordChangedAt)
	if err != nil {
		log.Println("Error Fetching User Account", err)
		return nil, err
	}

	return &account, nil
}

func FetchSalt(saltId int) (*models.Salt, error) {
	query := "SELECT * FROM Salts WHERE id = ?"

	salt := models.Salt{}
	err := db.QueryRow(query, saltId).Scan(&salt.Id, &salt.Salt, &salt.CreatedAt)
	if err != nil {
		log.Println("Error Fetching User Salt", err)
		return nil, err
	}

	return &salt, nil
}
