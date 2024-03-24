package repository

import (
	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"gorm.io/gorm"
)

type payementMethodRepository struct {
	database *gorm.DB
}

type PaymentMethodRepository interface {
	FindAll() (*[]domain.PaymentMethod, error)
	Create(payementMethod *domain.PaymentMethod) error
	Update(payementMethod *domain.PaymentMethod) error
	Delete(id uint) error
	FindByID(id uint) (*domain.PaymentMethod, error)
}

func NewPaymentMethodRepository(database *gorm.DB) PaymentMethodRepository {
	return &payementMethodRepository{
		database: database,
	}
}

func (repository *payementMethodRepository) FindAll() (*[]domain.PaymentMethod, error) {
	var categories []domain.PaymentMethod
	err := repository.database.Find(&categories).Error
	return &categories, err
}

func (repository *payementMethodRepository) Create(payementMethod *domain.PaymentMethod) error {
	return repository.database.Create(payementMethod).Error
}

func (repository *payementMethodRepository) Update(payementMethod *domain.PaymentMethod) error {
	return repository.database.Save(&payementMethod).Error
}

func (repository *payementMethodRepository) Delete(id uint) error {
	return repository.database.Delete(&domain.PaymentMethod{}, id).Error
}

func (repository *payementMethodRepository) FindByID(id uint) (*domain.PaymentMethod, error) {
	var payementMethod domain.PaymentMethod
	err := repository.database.First(&payementMethod, id).Error
	return &payementMethod, err
}
