package database

import (
	"golang1/models"
	"log"

	"github.com/jinzhu/gorm"
)

// Connector variable used for CRUD operation's
var Connector *gorm.DB

// Connect creates MySQL connection
func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	log.Println("Connection was successful!!")
	return nil
}

// Migrate create/updates database table
func Migrate(table *models.Article) {
	Connector.AutoMigrate(&table)
	log.Println("Table migrated")
}
