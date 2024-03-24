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
		&domain.Category{},
		&domain.Group{},
		&domain.Limit{},
		&domain.PaymentMethod{},
		&domain.Income{},
		&domain.Location{},
		&domain.User{},
		&domain.Customer{},
		&domain.Provider{},
		&domain.Product{},
		&domain.Price{},
		&domain.Cart{},
		&domain.CartProduct{},
		&domain.Order{},
		&domain.OrderProduct{},
		&domain.Calendar{},
		&domain.Event{},
	)

	// Setup join table
	err = db.SetupJoinTable(&domain.Cart{}, "Products", &domain.CartProduct{})
	if err != nil {
		panic("Failed to setup join table")
	}

	err = db.SetupJoinTable(&domain.Order{}, "Products", &domain.OrderProduct{})
	if err != nil {
		panic("Failed to setup join table")
	}

	// Set database
	DB = db
}
