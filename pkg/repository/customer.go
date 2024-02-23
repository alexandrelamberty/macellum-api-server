package repository

import (
	"fmt"

	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"gorm.io/gorm"
)

type customerRepository struct {
	database *gorm.DB
}

type CustomerRepository interface {
	FindAll() (*[]domain.Customer, error)
	Create(customer *domain.Customer) error
	Update(customer *domain.Customer) error
	Delete(id uint) error
	FindByID(id uint) (*domain.Customer, error)
	FindByName(name string) (*domain.Customer, error)
	FindByEmail(email string) (*domain.Customer, error)
}

func NewCustomerRepository(database *gorm.DB) CustomerRepository {
	return &customerRepository{
		database: database,
	}
}

func (ur *customerRepository) FindAll() (*[]domain.Customer, error) {
	var customers []domain.Customer
	err := ur.database.Find(&customers).Error
	return &customers, err
}

func (ur *customerRepository) Create(customer *domain.Customer) error {
	return ur.database.Create(customer).Error
}

func (ur *customerRepository) Update(customer *domain.Customer) error {
	return ur.database.Save(&customer).Error
}

func (r *customerRepository) Delete(id uint) error {
	return r.database.Delete(&domain.Customer{}, id).Error
}

func (r *customerRepository) FindByID(id uint) (*domain.Customer, error) {
	var customer domain.Customer
	err := r.database.First(&customer, id).Error
	fmt.Println(customer)
	return &customer, err
}

func (r *customerRepository) FindByEmail(email string) (*domain.Customer, error) {
	var customer domain.Customer
	err := r.database.Where("email = ?", email).First(&customer).Error
	return &customer, err
}

func (r *customerRepository) FindByName(name string) (*domain.Customer, error) {
	var customer domain.Customer
	err := r.database.Where("name = ?", name).First(&customer).Error
	return &customer, err
}
