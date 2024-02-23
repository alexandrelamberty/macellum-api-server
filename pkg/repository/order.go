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

func (ur *orderRepository) FindAll() (*[]domain.Order, error) {
	var orders []domain.Order
	err := ur.database.Find(&orders).Error
	return &orders, err
}

func (ur *orderRepository) Create(order *domain.Order) error {
	return ur.database.Create(order).Error
}

func (ur *orderRepository) Update(order *domain.Order) error {
	return ur.database.Save(&order).Error
}

func (r *orderRepository) Delete(id uint) error {
	return r.database.Delete(&domain.Order{}, id).Error
}

func (r *orderRepository) FindByID(id uint) (*domain.Order, error) {
	var order domain.Order
	err := r.database.First(&order, id).Error
	return &order, err
}
