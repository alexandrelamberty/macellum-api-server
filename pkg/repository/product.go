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

func (repository *productRepository) FindAll() (*[]domain.Product, error) {
	var products []domain.Product
	err := repository.database.Preload("Provider").
		Preload("Category").Preload("Group").Find(&products).Error
	return &products, err
}

func (repository *productRepository) Create(product *domain.Product) error {
	return repository.database.Create(product).Error
}

func (repository *productRepository) Update(product *domain.Product) error {
	return repository.database.Save(&product).Error
}

func (repository *productRepository) Delete(id uint) error {
	return repository.database.Delete(&domain.Product{}, id).Error
}

func (repository *productRepository) FindByID(id uint) (*domain.Product, error) {
	var product domain.Product
	err := repository.database.First(&product, id).Error
	return &product, err
}

func (repository *productRepository) FindBySKU(sku string) (*domain.Product, error) {
	var product domain.Product
	err := repository.database.Model(domain.Product{}).Where("sku = ?", sku).First(&product).Error
	return &product, err
}
