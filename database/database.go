package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func InitDatabase() {
	var err error
	enverr := godotenv.Load()
	if enverr != nil {
		log.Fatal("Error loading .env file")
	}

	dbpassword := os.Getenv("ROOT_PASSWORD")
	dbusername := os.Getenv("DATABASE_USER")
	dbname := os.Getenv("DATABASE_NAME")
	dbhost := os.Getenv("DATABASE_HOST")

	dsn := dbusername + ":" + dbpassword + "@tcp(" + dbhost + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("Database connected!")
}
