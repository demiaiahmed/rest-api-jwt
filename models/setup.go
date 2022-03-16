package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error Loading .env file")
	}

	DbDriver := os.Getenv("DB_DRIVER")
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	DB, err = gorm.Open(DbDriver, DBURL)

	sqlDB := DB.DB()
	defer sqlDB.Close()

	if err != nil {
		fmt.Println("Cannot Connect to Database", DbDriver)
		log.Fatal("Connection Error ", err)
	} else {
		fmt.Println("We are connected to database", DbDriver)
	}
	DB.AutoMigrate(&User{})
}
