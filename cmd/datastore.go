package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/ErwinSalas/inventory-service/models"
	"github.com/jinzhu/gorm"

	"github.com/joho/godotenv"
)

func InitDatastore() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	Dbdriver := os.Getenv("DB_DRIVER")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbPort := os.Getenv("DB_PORT")
	DbHost := os.Getenv("DB_HOST")
	DbName := os.Getenv("DB_NAME")

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	datastore, err := gorm.Open(Dbdriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", Dbdriver)
	}

	datastore.Debug().AutoMigrate(&models.Item{}) //database migration
	return datastore

}
