package repository

import (
	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"gorm.io/gorm"
)

type cartRepository struct {
	database *gorm.DB
}

type CartRepository interface {
	FindAll() (*[]domain.Cart, error)
	Create(cart *domain.Cart) error
	Update(cart *domain.Cart) error
	Delete(id uint) error
	FindByID(id uint) (*domain.Cart, error)
}

func NewCartRepository(database *gorm.DB) CartRepository {
	return &cartRepository{
		database: database,
	}
}

func (repository *cartRepository) FindAll() (*[]domain.Cart, error) {
	var carts []domain.Cart
	err := repository.database.Find(&carts).Error
	return &carts, err
}

func (repository *cartRepository) Create(cart *domain.Cart) error {
	return repository.database.Create(cart).Error
}

func (repository *cartRepository) Update(cart *domain.Cart) error {
	return repository.database.Save(&cart).Error
}

func (repository *cartRepository) Delete(id uint) error {
	return repository.database.Delete(&domain.Cart{}, id).Error
}

func (repository *cartRepository) FindByID(id uint) (*domain.Cart, error) {
	var cart domain.Cart
	err := repository.database.First(&cart, id).Error
	return &cart, err
}
