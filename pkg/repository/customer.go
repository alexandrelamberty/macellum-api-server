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

func (repository *customerRepository) FindAll() (*[]domain.Customer, error) {
	var customers []domain.Customer
	err := repository.database.Find(&customers).Error
	return &customers, err
}

func (repository *customerRepository) Create(customer *domain.Customer) error {
	return repository.database.Create(customer).Error
}

func (repository *customerRepository) Update(customer *domain.Customer) error {
	return repository.database.Save(&customer).Error
}

func (repository *customerRepository) Delete(id uint) error {
	return repository.database.Delete(&domain.Customer{}, id).Error
}

func (repository *customerRepository) FindByID(id uint) (*domain.Customer, error) {
	var customer domain.Customer
	err := repository.database.First(&customer, id).Error
	fmt.Println(customer)
	return &customer, err
}

func (repository *customerRepository) FindByEmail(email string) (*domain.Customer, error) {
	var customer domain.Customer
	err := repository.database.Where("email = ?", email).First(&customer).Error
	return &customer, err
}

func (repository *customerRepository) FindByName(name string) (*domain.Customer, error) {
	var customer domain.Customer
	err := repository.database.Where("name = ?", name).First(&customer).Error
	return &customer, err
}
