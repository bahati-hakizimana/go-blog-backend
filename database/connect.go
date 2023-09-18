package database

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/bahati-hakizimana/blogue-backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	error :=godotenv.Load()

	if  error !=nil {
		log.Fatal("error load .env file")
		
	}
	dsn:=os.Getenv("DSN")
	database,error:=gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if error !=nil{
		panic("could not connect to the database")
	}

	DB=database
	database.AutoMigrate(
		&models.User{},
		&models.Blog{},
	)
}