package service

import (
	"errors"

	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"github.com/alexandrelamberty/macellum-api-server/pkg/repository"
)

type productService struct {
	repository repository.ProductRepository
}

type ProductService interface {
	GetAllProducts() (*[]domain.Product, error)
	CreateProduct(*domain.Product) error
	UpdateProduct(product *domain.Product) error
	DeleteProduct(id uint) error
	GetProductByID(id uint) (*domain.Product, error)
	GetProductBySKU(sku string) (*domain.Product, error)
}

func NewProductService(r repository.ProductRepository) ProductService {
	return &productService{
		repository: r,
	}
}

func (s *productService) GetAllProducts() (*[]domain.Product, error) {
	return s.repository.FindAll()
}

func (s *productService) CreateProduct(product *domain.Product) error {
	// Check if product SKU already exists
	existingProduct, err := s.repository.FindBySKU(product.SKU)
	if err == nil && existingProduct != nil {
		return errors.New("productname already exists")
	}

	return s.repository.Create(product)
}

func (s *productService) UpdateProduct(product *domain.Product) error {
	return s.repository.Update(product)
}

func (s *productService) DeleteProduct(id uint) error {
	return s.repository.Delete(id)
}

func (s *productService) GetProductByID(id uint) (*domain.Product, error) {
	return s.repository.FindByID(id)
}

func (s *productService) GetProductBySKU(sku string) (*domain.Product, error) {
	return s.repository.FindBySKU(sku)
}
