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

func (ur *userRepository) FindAll() (*[]domain.User, error) {
	var users []domain.User
	err := ur.database.Find(&users).Error
	return &users, err
}

func (ur *userRepository) Create(user *domain.User) error {
	return ur.database.Create(user).Error
}

func (ur *userRepository) Update(user *domain.User) error {
	return ur.database.Save(&user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.database.Delete(&domain.User{}, id).Error
}

func (r *userRepository) FindByID(id uint) (*domain.User, error) {
	var user domain.User
	err := r.database.First(&user, id).Error
	return &user, err
}

func (r *userRepository) FindByUsername(username string) (*domain.User, error) {
	var user domain.User
	err := r.database.Model(domain.User{}).Where("username = ?", username).First(&user).Error
	return &user, err
}

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.database.Where("email = ?", email).First(&user).Error
	return &user, err
}
