package repository

import (
	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"gorm.io/gorm"
)

type limitRepository struct {
	database *gorm.DB
}

type LimitRepository interface {
	FindAll() (*[]domain.Limit, error)
	Create(limit *domain.Limit) error
	Update(limit *domain.Limit) error
	Delete(id uint) error
	FindByID(id uint) (*domain.Limit, error)
}

func NewLimitRepository(database *gorm.DB) LimitRepository {
	return &limitRepository{
		database: database,
	}
}

func (repository *limitRepository) FindAll() (*[]domain.Limit, error) {
	var categories []domain.Limit
	err := repository.database.Find(&categories).Error
	return &categories, err
}

func (repository *limitRepository) Create(limit *domain.Limit) error {
	return repository.database.Create(limit).Error
}

func (repository *limitRepository) Update(limit *domain.Limit) error {
	return repository.database.Save(&limit).Error
}

func (repository *limitRepository) Delete(id uint) error {
	return repository.database.Delete(&domain.Limit{}, id).Error
}

func (repository *limitRepository) FindByID(id uint) (*domain.Limit, error) {
	var limit domain.Limit
	err := repository.database.First(&limit, id).Error
	return &limit, err
}
