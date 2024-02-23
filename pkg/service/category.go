package service

import (
	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"github.com/alexandrelamberty/macellum-api-server/pkg/repository"
)

type categoryService struct {
	repository repository.CategoryRepository
}

type CategoryService interface {
	GetAllCategories() (*[]domain.Category, error)
	CreateCategory(*domain.Category) error
	UpdateCategory(category *domain.Category) error
	DeleteCategory(id uint) error
	GetCategoryByID(id uint) (*domain.Category, error)
}

func NewCategoryService(r repository.CategoryRepository) CategoryService {
	return &categoryService{
		repository: r,
	}
}

func (s *categoryService) GetAllCategories() (*[]domain.Category, error) {
	return s.repository.FindAll()
}

func (s *categoryService) CreateCategory(category *domain.Category) error {
	return s.repository.Create(category)
}

func (s *categoryService) UpdateCategory(category *domain.Category) error {
	return s.repository.Update(category)
}

func (s *categoryService) DeleteCategory(id uint) error {
	return s.repository.Delete(id)
}

func (s *categoryService) GetCategoryByID(id uint) (*domain.Category, error) {
	return s.repository.FindByID(id)
}
