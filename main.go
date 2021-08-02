package main

import (
	"golang1/controllers"
	"golang1/database"
	"golang1/entity"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on http://localhost:8080")

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Router API
func initaliseHandlers(router *mux.Router) {
	// Router Article
	router.HandleFunc("/api/v1/article/add", controllers.CreateArticle).Methods("POST")        // Create
	router.HandleFunc("/api/v1/articleList", controllers.GetAllArticle).Methods("POST")        // Read
	router.HandleFunc("/api/v1/articleList/{id}", controllers.GetArticleByID).Methods("GET")   // Read ByID
	router.HandleFunc("/api/v1/article/{id}", controllers.UpdateArticleByID).Methods("PUT")    // Update
	router.HandleFunc("/api/v1/article/{id}", controllers.DeleteArticleByID).Methods("DELETE") // Delete
}

func initDB() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "12345678",
			DB:         "db_golang3",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&entity.Article{})
}
