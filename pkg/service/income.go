package service

import (
	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"github.com/alexandrelamberty/macellum-api-server/pkg/repository"
)

type incomeService struct {
	repository repository.IncomeRepository
}

type IncomeService interface {
	GetAllIncomes() (*[]domain.Income, error)
	CreateIncome(*domain.Income) error
	UpdateIncome(income *domain.Income) error
	DeleteIncome(id uint) error
	GetIncomeByID(id uint) (*domain.Income, error)
}

func NewIncomeService(r repository.IncomeRepository) IncomeService {
	return &incomeService{
		repository: r,
	}
}

func (s *incomeService) GetAllIncomes() (*[]domain.Income, error) {
	return s.repository.FindAll()
}

func (s *incomeService) CreateIncome(income *domain.Income) error {
	return s.repository.Create(income)
}

func (s *incomeService) UpdateIncome(income *domain.Income) error {
	return s.repository.Update(income)
}

func (s *incomeService) DeleteIncome(id uint) error {
	return s.repository.Delete(id)
}

func (s *incomeService) GetIncomeByID(id uint) (*domain.Income, error) {
	return s.repository.FindByID(id)
}
