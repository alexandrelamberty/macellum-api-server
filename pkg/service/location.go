package service

import (
	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"github.com/alexandrelamberty/macellum-api-server/pkg/repository"
)

type locationService struct {
	repository repository.LocationRepository
}

type LocationService interface {
	GetAllLocations() (*[]domain.Location, error)
	CreateLocation(*domain.Location) error
	UpdateLocation(location *domain.Location) error
	DeleteLocation(id uint) error
	GetLocationByID(id uint) (*domain.Location, error)
}

func NewLocationService(r repository.LocationRepository) LocationService {
	return &locationService{
		repository: r,
	}
}

func (s *locationService) GetAllLocations() (*[]domain.Location, error) {
	return s.repository.FindAll()
}

func (s *locationService) CreateLocation(location *domain.Location) error {
	return s.repository.Create(location)
}

func (s *locationService) UpdateLocation(location *domain.Location) error {
	return s.repository.Update(location)
}

func (s *locationService) DeleteLocation(id uint) error {
	return s.repository.Delete(id)
}

func (s *locationService) GetLocationByID(id uint) (*domain.Location, error) {
	return s.repository.FindByID(id)
}
