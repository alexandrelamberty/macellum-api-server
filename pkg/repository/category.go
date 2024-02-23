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

func (ur *categoryRepository) FindAll() (*[]domain.Category, error) {
	var categories []domain.Category
	err := ur.database.Find(&categories).Error
	return &categories, err
}

func (ur *categoryRepository) Create(category *domain.Category) error {
	return ur.database.Create(category).Error
}

func (ur *categoryRepository) Update(category *domain.Category) error {
	return ur.database.Save(&category).Error
}

func (r *categoryRepository) Delete(id uint) error {
	return r.database.Delete(&domain.Category{}, id).Error
}

func (r *categoryRepository) FindByID(id uint) (*domain.Category, error) {
	var category domain.Category
	err := r.database.First(&category, id).Error
	return &category, err
}
