package infrastructure

import (
	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDatabase() {
	dsn := "macellum:macellum@tcp(127.0.0.1:3306)/macellum?charset=utf8mb4&parseTime=True&loc=Local"

	// Connect to the database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Migrate the schemas
	db.AutoMigrate(
		&domain.Calendar{},
		&domain.Category{},
		&domain.Customer{},
		&domain.Event{},
		&domain.Cart{},
		&domain.Order{},
		&domain.OrderProduct{},
		&domain.Product{},
		&domain.Provider{},
		&domain.ProductProvider{},
		&domain.User{})

	DB = db
}
