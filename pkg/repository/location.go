package repository

import (
	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"gorm.io/gorm"
)

type locationRepository struct {
	database *gorm.DB
}

type LocationRepository interface {
	FindAll() (*[]domain.Location, error)
	Create(location *domain.Location) error
	Update(location *domain.Location) error
	Delete(id uint) error
	FindByID(id uint) (*domain.Location, error)
}

func NewLocationRepository(database *gorm.DB) LocationRepository {
	return &locationRepository{
		database: database,
	}
}

func (repository *locationRepository) FindAll() (*[]domain.Location, error) {
	var categories []domain.Location
	err := repository.database.Find(&categories).Error
	return &categories, err
}

func (repository *locationRepository) Create(location *domain.Location) error {
	return repository.database.Create(location).Error
}

func (repository *locationRepository) Update(location *domain.Location) error {
	return repository.database.Save(&location).Error
}

func (repository *locationRepository) Delete(id uint) error {
	return repository.database.Delete(&domain.Location{}, id).Error
}

func (repository *locationRepository) FindByID(id uint) (*domain.Location, error) {
	var location domain.Location
	err := repository.database.First(&location, id).Error
	return &location, err
}
