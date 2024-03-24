package service

import (
	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"github.com/alexandrelamberty/macellum-api-server/pkg/repository"
)

type payementMethodService struct {
	repository repository.PaymentMethodRepository
}

type PaymentMethodService interface {
	GetAllPaymentMethods() (*[]domain.PaymentMethod, error)
	CreatePaymentMethod(*domain.PaymentMethod) error
	UpdatePaymentMethod(payementMethod *domain.PaymentMethod) error
	DeletePaymentMethod(id uint) error
	GetPaymentMethodByID(id uint) (*domain.PaymentMethod, error)
}

func NewPaymentMethodService(r repository.PaymentMethodRepository) PaymentMethodService {
	return &payementMethodService{
		repository: r,
	}
}

func (s *payementMethodService) GetAllPaymentMethods() (*[]domain.PaymentMethod, error) {
	return s.repository.FindAll()
}

func (s *payementMethodService) CreatePaymentMethod(payementMethod *domain.PaymentMethod) error {
	return s.repository.Create(payementMethod)
}

func (s *payementMethodService) UpdatePaymentMethod(payementMethod *domain.PaymentMethod) error {
	return s.repository.Update(payementMethod)
}

func (s *payementMethodService) DeletePaymentMethod(id uint) error {
	return s.repository.Delete(id)
}

func (s *payementMethodService) GetPaymentMethodByID(id uint) (*domain.PaymentMethod, error) {
	return s.repository.FindByID(id)
}
