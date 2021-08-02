package main

import (
	"fmt"
	config "golang3/config"
	models "golang3/models"
	"golang3/routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer config.DB.Close()
	//	Config.DB.DropTableIfExists(&Model.Article{}, &Model.Account{})
	config.DB.AutoMigrate(&models.Article{})
	r := routes.SetUpRouter()

	r.Run()
}
