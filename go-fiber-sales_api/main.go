package main

import (
	"fmt"
	"os"

	"github.com/Valgard/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	fmt.Println("deve code app running...")

	godotenv.Load()

	dbhost = os.Getenv("MYSQL_HOST")
	dbpassword = os.Getenv("MYSQL_PASSWORD")
	dbuser = os.Getenv("MYSQL_USER")
	dbname = os.Getenv("MYSQL_DBNAME")

	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpassword)
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	DB = db
	fmt.Println("connection successful")
}
