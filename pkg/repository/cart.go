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

func (ur *cartRepository) FindAll() (*[]domain.Cart, error) {
	var carts []domain.Cart
	err := ur.database.Find(&carts).Error
	return &carts, err
}

func (ur *cartRepository) Create(cart *domain.Cart) error {
	return ur.database.Create(cart).Error
}

func (ur *cartRepository) Update(cart *domain.Cart) error {
	return ur.database.Save(&cart).Error
}

func (r *cartRepository) Delete(id uint) error {
	return r.database.Delete(&domain.Cart{}, id).Error
}

func (r *cartRepository) FindByID(id uint) (*domain.Cart, error) {
	var cart domain.Cart
	err := r.database.First(&cart, id).Error
	return &cart, err
}
