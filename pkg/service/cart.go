package service

import (
	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"github.com/alexandrelamberty/macellum-api-server/pkg/repository"
)

type cartService struct {
	repository repository.CartRepository
}

type CartService interface {
	GetAllCarts() (*[]domain.Cart, error)
	CreateCart(*domain.Cart) error
	UpdateCart(cart *domain.Cart) error
	DeleteCart(id uint) error
	GetCartByID(id uint) (*domain.Cart, error)
}

func NewCartService(r repository.CartRepository) CartService {
	return &cartService{
		repository: r,
	}
}

func (s *cartService) GetAllCarts() (*[]domain.Cart, error) {
	return s.repository.FindAll()
}

func (s *cartService) CreateCart(cart *domain.Cart) error {
	return s.repository.Create(cart)
}

func (s *cartService) UpdateCart(cart *domain.Cart) error {
	return s.repository.Update(cart)
}

func (s *cartService) DeleteCart(id uint) error {
	return s.repository.Delete(id)
}

func (s *cartService) GetCartByID(id uint) (*domain.Cart, error) {
	return s.repository.FindByID(id)
}
