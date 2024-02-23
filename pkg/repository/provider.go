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

func (ur *providerRepository) FindAll() (*[]domain.Provider, error) {
	var providers []domain.Provider
	err := ur.database.Find(&providers).Error
	return &providers, err
}

func (ur *providerRepository) Create(provider *domain.Provider) error {
	return ur.database.Create(provider).Error
}

func (ur *providerRepository) Update(provider *domain.Provider) error {
	return ur.database.Save(&provider).Error
}

func (r *providerRepository) Delete(id uint) error {
	return r.database.Delete(&domain.Provider{}, id).Error
}

func (r *providerRepository) FindByID(id uint) (*domain.Provider, error) {
	var provider domain.Provider
	err := r.database.First(&provider, id).Error
	return &provider, err
}

func (r *providerRepository) FindByName(name string) (*domain.Provider, error) {
	var provider domain.Provider
	err := r.database.First(&provider, name).Error
	return &provider, err
}
