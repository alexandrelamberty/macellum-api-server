package service

import (
	"errors"
	"time"

	"github.com/alexandrelamberty/macellum-api-server/internal/config"
	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"github.com/alexandrelamberty/macellum-api-server/pkg/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	repository repository.UserRepository
}

type AuthService interface {
	RegisterUser(*domain.User) error

	LoginUser(email string, pass string) (*domain.User, error)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func NewAuthService(r repository.UserRepository) AuthService {
	return &authService{
		repository: r,
	}
}

func (s *authService) RegisterUser(user *domain.User) error {
	// Check if username or email already exists
	existingAuth, err := s.repository.FindByUsername(user.FirstName + " " + user.LastName)
	if err == nil && existingAuth != nil {
		return errors.New("username already exists")
	}

	existingAuth, err = s.repository.FindByEmail(user.Email)
	if err == nil && existingAuth != nil {
		return errors.New("email already exists")
	}

	// Hash password
	password, err := hashPassword(user.Password)
	if err != nil {
		return errors.New("failed to hash password")
	}
	user.Password = password

	return s.repository.Create(user)
}

func (s *authService) LoginUser(email string, pass string) (*domain.User, error) {
	var user *domain.User
	user, err := s.repository.FindByEmail(email)
	// Check if user exists, if not return error
	if err == nil && user == nil {
		return nil, errors.New("user not found")
	}

	// Check password hash if user exists
	if !CheckPasswordHash(pass, user.Password) {
		return nil, errors.New("invalid password")
	}

	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	_, err = token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return nil, err
	}

	return user, nil
}
