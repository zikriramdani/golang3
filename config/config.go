package Config

import (
	"fmt"
	"os"
	"log"
	"strconv"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
  
	return os.Getenv(key)
  }
  

func BuildDBConfig() *DBConfig {
	dbPort, _ := strconv.Atoi(goDotEnvVariable("MYSQL_PORT"))

	dbConfig := DBConfig{
		Host:     goDotEnvVariable("MYSQL_HOST"),
		Port:     dbPort,
		User:     goDotEnvVariable("MYSQL_USER"),
		Password: goDotEnvVariable("MYSQL_PASS"),
		DBName:   goDotEnvVariable("MYSQL_NAME"),
	}
	return &dbConfig
}
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
