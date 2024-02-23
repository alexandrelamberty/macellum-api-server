package service

import (
	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"github.com/alexandrelamberty/macellum-api-server/pkg/repository"
)

type providerService struct {
	repository repository.ProviderRepository
}

type ProviderService interface {
	GetAllProviders() (*[]domain.Provider, error)
	CreateProvider(*domain.Provider) error
	UpdateProvider(provider *domain.Provider) error
	DeleteProvider(id uint) error
	GetProviderByID(id uint) (*domain.Provider, error)
	GetProviderByName(name string) (*domain.Provider, error)
}

func NewProviderService(r repository.ProviderRepository) ProviderService {
	return &providerService{
		repository: r,
	}
}

func (s *providerService) GetAllProviders() (*[]domain.Provider, error) {
	return s.repository.FindAll()
}

func (s *providerService) CreateProvider(provider *domain.Provider) error {
	// Check if provider name already exists

	return s.repository.Create(provider)
}

func (s *providerService) UpdateProvider(provider *domain.Provider) error {
	return s.repository.Update(provider)
}

func (s *providerService) DeleteProvider(id uint) error {
	return s.repository.Delete(id)
}

func (s *providerService) GetProviderByID(id uint) (*domain.Provider, error) {
	return s.repository.FindByID(id)
}

func (s *providerService) GetProviderByName(name string) (*domain.Provider, error) {
	return s.repository.FindByName(name)
}
