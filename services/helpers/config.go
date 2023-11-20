package helpers

import (
	"encoding/json"
	"log"
	"os"

	"github.com/vamsikartik01/Authentication_service/api/models"
)

var Config *models.Config

func LoadConfig() error {
	filename := "config.json"
	file, err := os.Open(filename)
	if err != nil {
		log.Println("Error Opening Config file", err)
		return err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&Config)

	if err != nil {
		log.Println("Error Reading as json object", err)
		return err
	}
	return nil
}
