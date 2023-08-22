package db

import (
	"fmt"
	"log"
	"os"

	"github.com/Valgard/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"go-fiber/model"
)

var DB *gorm.DB

func Connect() {

	godotenv.Load(".env")

	dbhost := os.Getenv("MYSQL_HOST")
	dbpassword := os.Getenv("MYSQL_PASSWORD")
	dbuser := os.Getenv("MYSQL_USER")
	dbname := os.Getenv("MYSQL_DBNAME")

	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpassword, dbhost, dbname)
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	DB = db
	fmt.Println("connection successful")

	AutoMigrate(db)
	if err := DB.AutoMigrate(&model.Cashier{}, &model.Category{}, &model.Payment{}, &model.PaymentType{}, &model.Product{}, &model.Discount{}, &model.Order{}).Error; err != nil {
		log.Fatalf("Migration failed %v", err)
	}

}

func AutoMigrate(connection *gorm.DB) {
	connection.Debug().AutoMigrate(
		&model.Cashier{},
		&model.Category{},
		&model.Payment{},
		&model.PaymentType{},
		&model.Product{},
		&model.Discount{},
		&model.Order{},
	)
}
