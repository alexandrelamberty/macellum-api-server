package repository

import (
	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"gorm.io/gorm"
)

type categoryRepository struct {
	database *gorm.DB
}

type CategoryRepository interface {
	FindAll() (*[]domain.Category, error)
	Create(category *domain.Category) error
	Update(category *domain.Category) error
	Delete(id uint) error
	FindByID(id uint) (*domain.Category, error)
}

func NewCategoryRepository(database *gorm.DB) CategoryRepository {
	return &categoryRepository{
		database: database,
	}
}

func (repository *categoryRepository) FindAll() (*[]domain.Category, error) {
	var categories []domain.Category
	err := repository.database.Find(&categories).Error
	return &categories, err
}

func (repository *categoryRepository) Create(category *domain.Category) error {
	return repository.database.Create(category).Error
}

func (repository *categoryRepository) Update(category *domain.Category) error {
	return repository.database.Save(&category).Error
}

func (repository *categoryRepository) Delete(id uint) error {
	return repository.database.Delete(&domain.Category{}, id).Error
}

func (repository *categoryRepository) FindByID(id uint) (*domain.Category, error) {
	var category domain.Category
	err := repository.database.First(&category, id).Error
	return &category, err
}
