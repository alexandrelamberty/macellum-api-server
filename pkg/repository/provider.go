package repository

import (
	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"gorm.io/gorm"
)

type providerRepository struct {
	database *gorm.DB
}

type ProviderRepository interface {
	FindAll() (*[]domain.Provider, error)
	Create(provider *domain.Provider) error
	Update(provider *domain.Provider) error
	Delete(id uint) error
	FindByID(id uint) (*domain.Provider, error)
	FindByName(name string) (*domain.Provider, error)
}

func NewProviderRepository(database *gorm.DB) ProviderRepository {
	return &providerRepository{
		database: database,
	}
}

func (repository *providerRepository) FindAll() (*[]domain.Provider, error) {
	var providers []domain.Provider
	err := repository.database.Find(&providers).Error
	return &providers, err
}

func (repository *providerRepository) Create(provider *domain.Provider) error {
	return repository.database.Create(provider).Error
}

func (repository *providerRepository) Update(provider *domain.Provider) error {
	return repository.database.Save(&provider).Error
}

func (repository *providerRepository) Delete(id uint) error {
	return repository.database.Delete(&domain.Provider{}, id).Error
}

func (repository *providerRepository) FindByID(id uint) (*domain.Provider, error) {
	var provider domain.Provider
	err := repository.database.First(&provider, id).Error
	return &provider, err
}

func (repository *providerRepository) FindByName(name string) (*domain.Provider, error) {
	var provider domain.Provider
	err := repository.database.First(&provider, name).Error
	return &provider, err
}
