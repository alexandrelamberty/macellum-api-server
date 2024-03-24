package repository

import (
	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"gorm.io/gorm"
)

type orderRepository struct {
	database *gorm.DB
}

type OrderRepository interface {
	FindAll() (*[]domain.Order, error)
	Create(order *domain.Order) error
	Update(order *domain.Order) error
	Delete(id uint) error
	FindByID(id uint) (*domain.Order, error)
}

func NewOrderRepository(database *gorm.DB) OrderRepository {
	return &orderRepository{
		database: database,
	}
}

func (repository *orderRepository) FindAll() (*[]domain.Order, error) {
	var orders []domain.Order
	err := repository.database.Find(&orders).Error
	return &orders, err
}

func (repository *orderRepository) Create(order *domain.Order) error {
	return repository.database.Create(order).Error
}

func (repository *orderRepository) Update(order *domain.Order) error {
	return repository.database.Save(&order).Error
}

func (repository *orderRepository) Delete(id uint) error {
	return repository.database.Delete(&domain.Order{}, id).Error
}

func (repository *orderRepository) FindByID(id uint) (*domain.Order, error) {
	var order domain.Order
	err := repository.database.First(&order, id).Error
	return &order, err
}
