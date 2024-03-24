package repository

import (
	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	database *gorm.DB
}

type UserRepository interface {
	FindAll() (*[]domain.User, error)
	Create(user *domain.User) error
	Update(user *domain.User) error
	Delete(id uint) error
	FindByID(id uint) (*domain.User, error)
	FindByUsername(username string) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
}

func NewUserRepository(database *gorm.DB) UserRepository {
	return &userRepository{
		database: database,
	}
}

func (repository *userRepository) FindAll() (*[]domain.User, error) {
	var users []domain.User
	err := repository.database.Find(&users).Error
	return &users, err
}

func (repository *userRepository) Create(user *domain.User) error {
	return repository.database.Create(user).Error
}

func (repository *userRepository) Update(user *domain.User) error {
	return repository.database.Save(&user).Error
}

func (repository *userRepository) Delete(id uint) error {
	return repository.database.Delete(&domain.User{}, id).Error
}

func (repository *userRepository) FindByID(id uint) (*domain.User, error) {
	var user domain.User
	err := repository.database.First(&user, id).Error
	return &user, err
}

func (repository *userRepository) FindByUsername(username string) (*domain.User, error) {
	var user domain.User
	err := repository.database.Model(domain.User{}).Where("username = ?", username).First(&user).Error
	return &user, err
}

func (repository *userRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := repository.database.Where("email = ?", email).First(&user).Error
	return &user, err
}
