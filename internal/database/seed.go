package database

import (
	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"gorm.io/gorm"
)

func SeedDatabase(db *gorm.DB) {
	// Insert initial definitions
	var categories = []domain.Category{
		{Name: "Biscuits"},
		{Name: "Cereals"},
		{Name: "Cleaning products"},
		{Name: "Coffee & Tea"},
		{Name: "Hygiene"},
		{Name: "Jars & Cans"},
		{Name: "Laundry"},
		{Name: "Milk"},
		{Name: "Misc"},
		{Name: "Pasta"},
		{Name: "Rice"},
		{Name: "Sauce, Oil & Condiment"},
		{Name: "Sweets"},
		{Name: "Syrup & Juice"},
		{Name: "Toothpaste"},
		{Name: "Unknown"},
	}
	db.Create(&categories)

	var groups = []domain.Group{
		{Name: "Cleaning products"},
		{Name: "Creamery"},
		{Name: "Fruits"},
		{Name: "Grocery Store"},
		{Name: "Hygiene"},
		{Name: "Non food"},
		{Name: "Unknown"},
		{Name: "Vegetables"},
	}
	db.Create(&groups)

	var paymentMethods = []domain.PaymentMethod{
		{Name: "Cash"},
		{Name: "Payconic"},
	}
	db.Create(&paymentMethods)

	var limits = []domain.Limit{
		{Name: "Individual", Peoples: 1, Days: 30, Amount: 10},
		{Name: "2", Peoples: 2, Days: 30, Amount: 10},
		{Name: "3", Peoples: 3, Days: 30, Amount: 10},
		{Name: "4", Peoples: 4, Days: 30, Amount: 10},
		{Name: "5", Peoples: 5, Days: 30, Amount: 10},
		{Name: "6", Peoples: 6, Days: 30, Amount: 10},
		{Name: "7", Peoples: 7, Days: 30, Amount: 10},
		{Name: "8", Peoples: 8, Days: 30, Amount: 10},
		{Name: "9", Peoples: 9, Days: 30, Amount: 10},
	}
	db.Create(&limits)

	var incomes = []domain.Income{
		{Name: "Disability allowance"},
		{Name: "Mutual insurance"},
		{Name: "Pension"},
		{Name: "RIS"},
		{Name: "Unemployment benefit"},
		{Name: "Without income"},
		{Name: "Work"},
	}
	db.Create(&incomes)

	var locations = []domain.Location{
		{Name: "Chastre", Address: domain.Address{Street: "123 Main St", Number: "New York", ZipCode: "NY", City: "Chastre"}, Phone: "123-456-7890", Email: "5HJ9z@example.com"},
		{Name: "Mont-St-Guibert", Address: domain.Address{Street: "123 Main St", Number: "New York", ZipCode: "NY", City: "Mont-St-Guibert"}, Phone: "123-456-7890", Email: "5HJ9z@example.com"},
		{Name: "Walhain", Address: domain.Address{Street: "123 Main St", Number: "New York", ZipCode: "NY", City: "Walhain"}, Phone: "123-456-7890", Email: "5HJ9z@example.com"},
	}
	db.Create(&locations)

	var providers = []domain.Provider{
		{Name: "Colruyt", Address: domain.Address{Street: "123 Main St", Number: "New York", ZipCode: "NY", City: "Chastre"}},
		{Name: "Bio-Planet", Address: domain.Address{Street: "123 Main St", Number: "New York", ZipCode: "NY", City: "Mont-St-Guibert"}},
		{Name: "Carrefour Market", Address: domain.Address{Street: "123 Main St", Number: "New York", ZipCode: "NY", City: "Walhain"}},
	}
	db.Create(&providers)

	var customers = []domain.Customer{
		{Code: "C001", FirstName: "John", LastName: "Doe", BirthDate: "01-01-1990", Address: domain.Address{Street: "123 Main St", Number: "New York", ZipCode: "NY", City: "Chastre"}},
		{Code: "C002", FirstName: "Jane", LastName: "Doe", BirthDate: "01-01-1990", Address: domain.Address{Street: "123 Main St", Number: "New York", ZipCode: "NY", City: "Mont-St-Guibert"}},
		{Code: "C003", FirstName: "Bob", LastName: "Smith", BirthDate: "01-01-1990", Address: domain.Address{Street: "123 Main St", Number: "New York", ZipCode: "NY", City: "Walhain"}},
	}
	db.Create(&customers)

	var users = []domain.User{
		{FirstName: "John", LastName: "Doe", Email: "a@a.com", Password: "$2a$12$alXkQB/xKfoS6/eJOQg4G.LpRW9HaSuJ47o8htplVWP6PvjxUB5.S"},
		{FirstName: "Jane", LastName: "Doe", Email: "b@b.com", Password: "$2a$12$alXkQB/xKfoS6/eJOQg4G.LpRW9HaSuJ47o8htplVWP6PvjxUB5.S"},
		{FirstName: "Bob", LastName: "Smith", Email: "c@c.com", Password: "$2a$12$alXkQB/xKfoS6/eJOQg4G.LpRW9HaSuJ47o8htplVWP6PvjxUB5.S"},
	}
	db.Create(&users)

	var products = []domain.Product{
		{
			Code:  "8978796896",
			Name:  "Peach Syrup",
			Brand: "Everyday",
			Prices: []domain.Price{
				{ProductID: 1, Price: 1.00, Date: "2022-01-01"},
			},
			Weight: domain.Weight{
				Value: 1,
				Unit:  "kg",
			},
			Stock: domain.Stock{
				Shelf:    "A",
				Quantity: 10,
				Minimum:  5,
				Maximum:  20,
				InStock:  true,
			},
			Store: domain.Store{
				Shelf:    "A",
				Quantity: 10,
				Minimum:  5,
				Maximum:  20,
				InStore:  true,
			},
			ProviderID: 1,
			CategoryID: 1,
			GroupID:    1,
		},
	}
	db.Create(&products)

}
