package service

import (
	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"github.com/alexandrelamberty/macellum-api-server/pkg/repository"
)

type limitService struct {
	repository repository.LimitRepository
}

type LimitService interface {
	GetAllLimits() (*[]domain.Limit, error)
	CreateLimit(*domain.Limit) error
	UpdateLimit(limit *domain.Limit) error
	DeleteLimit(id uint) error
	GetLimitByID(id uint) (*domain.Limit, error)
}

func NewLimitService(r repository.LimitRepository) LimitService {
	return &limitService{
		repository: r,
	}
}

func (s *limitService) GetAllLimits() (*[]domain.Limit, error) {
	return s.repository.FindAll()
}

func (s *limitService) CreateLimit(limit *domain.Limit) error {
	return s.repository.Create(limit)
}

func (s *limitService) UpdateLimit(limit *domain.Limit) error {
	return s.repository.Update(limit)
}

func (s *limitService) DeleteLimit(id uint) error {
	return s.repository.Delete(id)
}

func (s *limitService) GetLimitByID(id uint) (*domain.Limit, error) {
	return s.repository.FindByID(id)
}
