package service

import (
	"errors"

	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"github.com/alexandrelamberty/macellum-api-server/pkg/repository"
)

type userService struct {
	repository repository.UserRepository
}

type UserService interface {
	GetAllUsers() (*[]domain.User, error)
	InviteUser(*domain.User) error
	UpdateUser(user *domain.User) error
	DeleteUser(id uint) error
	GetUserByID(id uint) (*domain.User, error)
	GetUserByUsername(username string) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{
		repository: r,
	}
}

func (s *userService) GetAllUsers() (*[]domain.User, error) {
	return s.repository.FindAll()
}

func (s *userService) InviteUser(user *domain.User) error {
	// Check if email already exists
	existingUser, err := s.repository.FindByEmail(user.Email)
	if err == nil && existingUser != nil {
		return errors.New("email already exists")
	}

	return s.repository.Create(user)
}

func (s *userService) UpdateUser(user *domain.User) error {
	return s.repository.Update(user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.repository.Delete(id)
}

func (s *userService) GetUserByID(id uint) (*domain.User, error) {
	return s.repository.FindByID(id)
}

func (s *userService) GetUserByUsername(username string) (*domain.User, error) {
	return s.repository.FindByUsername(username)
}

func (s *userService) GetUserByEmail(email string) (*domain.User, error) {
	return s.repository.FindByEmail(email)
}
