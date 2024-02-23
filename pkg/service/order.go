package service

import (
	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"github.com/alexandrelamberty/macellum-api-server/pkg/repository"
)

type orderService struct {
	repository repository.OrderRepository
}

type OrderService interface {
	GetAllOrders() (*[]domain.Order, error)
	CreateOrder(*domain.Order) error
	UpdateOrder(order *domain.Order) error
	DeleteOrder(id uint) error
	GetOrderByID(id uint) (*domain.Order, error)
}

func NewOrderService(r repository.OrderRepository) OrderService {
	return &orderService{
		repository: r,
	}
}

func (s *orderService) GetAllOrders() (*[]domain.Order, error) {
	return s.repository.FindAll()
}

func (s *orderService) CreateOrder(order *domain.Order) error {
	return s.repository.Create(order)
}

func (s *orderService) UpdateOrder(order *domain.Order) error {
	return s.repository.Update(order)
}

func (s *orderService) DeleteOrder(id uint) error {
	return s.repository.Delete(id)
}

func (s *orderService) GetOrderByID(id uint) (*domain.Order, error) {
	return s.repository.FindByID(id)
}
