package repository

import (
	"efishery-golang/entity"
	"encoding/json"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

type Resources struct {
	db *gorm.DB
}

func NewResources(DB *gorm.DB) Resources {
	return Resources{
		db: DB,
	}
}

func (repo Resources) GetResourceList() ([]entity.Resources, error) {
	var data []entity.Resources
	var client http.Client
	response, err := client.Get(os.Getenv("EFISHERY_BASE_URL") + "storages/5e1edf521073e315924ceab4/list")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		err := json.NewDecoder(response.Body).Decode(&data)
		if err != nil {
			log.Fatal(err)
		}
	}
	return data, err
}
