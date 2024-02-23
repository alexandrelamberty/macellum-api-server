package service

import (
	"errors"

	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"github.com/alexandrelamberty/macellum-api-server/pkg/repository"
)

type customerService struct {
	repository repository.CustomerRepository
}

type CustomerService interface {
	GetAllCustomers() (*[]domain.Customer, error)
	CreateCustomer(*domain.Customer) error
	UpdateCustomer(customer *domain.Customer) error
	DeleteCustomer(id uint) error
	GetCustomerByID(id uint) (*domain.Customer, error)
	GetCustomerByName(name string) (*domain.Customer, error)
	GetCustomerByEmail(email string) (*domain.Customer, error)
}

func NewCustomerService(r repository.CustomerRepository) CustomerService {
	return &customerService{
		repository: r,
	}
}

func (s *customerService) GetAllCustomers() (*[]domain.Customer, error) {
	return s.repository.FindAll()
}

func (s *customerService) CreateCustomer(customer *domain.Customer) error {
	// Check if customername or email already exists
	existingCustomer, err := s.repository.FindByName(customer.FirstName + " " + customer.LastName)
	if err == nil && existingCustomer != nil {
		return errors.New("customername already exists")
	}

	// TODO: check phone and address ?

	existingCustomer, err = s.repository.FindByEmail(customer.Email)
	if err == nil && existingCustomer != nil {
		return errors.New("email already exists")
	}

	return s.repository.Create(customer)
}

func (s *customerService) UpdateCustomer(customer *domain.Customer) error {
	return s.repository.Update(customer)
}

func (s *customerService) DeleteCustomer(id uint) error {
	return s.repository.Delete(id)
}

func (s *customerService) GetCustomerByID(id uint) (*domain.Customer, error) {
	return s.repository.FindByID(id)
}

func (s *customerService) GetCustomerByName(name string) (*domain.Customer, error) {
	return s.repository.FindByName(name)
}

func (s *customerService) GetCustomerByEmail(email string) (*domain.Customer, error) {
	return s.repository.FindByEmail(email)
}
