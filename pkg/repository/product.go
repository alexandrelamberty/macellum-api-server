package repository

import (
	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"gorm.io/gorm"
)

type productRepository struct {
	database *gorm.DB
}

type ProductRepository interface {
	FindAll() (*[]domain.Product, error)
	Create(product *domain.Product) error
	Update(product *domain.Product) error
	Delete(id uint) error
	FindByID(id uint) (*domain.Product, error)
	FindBySKU(sku string) (*domain.Product, error)
}

func NewProductRepository(database *gorm.DB) ProductRepository {
	return &productRepository{
		database: database,
	}
}

func (ur *productRepository) FindAll() (*[]domain.Product, error) {
	var products []domain.Product
	err := ur.database.Find(&products).Error
	return &products, err
}

func (ur *productRepository) Create(product *domain.Product) error {
	return ur.database.Create(product).Error
}

func (ur *productRepository) Update(product *domain.Product) error {
	return ur.database.Save(&product).Error
}

func (r *productRepository) Delete(id uint) error {
	return r.database.Delete(&domain.Product{}, id).Error
}

func (r *productRepository) FindByID(id uint) (*domain.Product, error) {
	var product domain.Product
	err := r.database.First(&product, id).Error
	return &product, err
}

func (r *productRepository) FindBySKU(sku string) (*domain.Product, error) {
	var product domain.Product
	err := r.database.Model(domain.Product{}).Where("sku = ?", sku).First(&product).Error
	return &product, err
}
