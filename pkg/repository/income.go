package repository

import (
	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"gorm.io/gorm"
)

type incomeRepository struct {
	database *gorm.DB
}

type IncomeRepository interface {
	FindAll() (*[]domain.Income, error)
	Create(income *domain.Income) error
	Update(income *domain.Income) error
	Delete(id uint) error
	FindByID(id uint) (*domain.Income, error)
}

func NewIncomeRepository(database *gorm.DB) IncomeRepository {
	return &incomeRepository{
		database: database,
	}
}

func (repository *incomeRepository) FindAll() (*[]domain.Income, error) {
	var incomes []domain.Income
	err := repository.database.Find(&incomes).Error
	return &incomes, err
}

func (repository *incomeRepository) Create(income *domain.Income) error {
	return repository.database.Create(income).Error
}

func (repository *incomeRepository) Update(income *domain.Income) error {
	return repository.database.Save(&income).Error
}

func (repository *incomeRepository) Delete(id uint) error {
	return repository.database.Delete(&domain.Income{}, id).Error
}

func (repository *incomeRepository) FindByID(id uint) (*domain.Income, error) {
	var income domain.Income
	err := repository.database.First(&income, id).Error
	return &income, err
}
