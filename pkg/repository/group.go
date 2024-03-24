package repository

import (
	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"gorm.io/gorm"
)

type groupRepository struct {
	database *gorm.DB
}

type GroupRepository interface {
	FindAll() (*[]domain.Group, error)
	Create(group *domain.Group) error
	Update(group *domain.Group) error
	Delete(id uint) error
	FindByID(id uint) (*domain.Group, error)
}

func NewGroupRepository(database *gorm.DB) GroupRepository {
	return &groupRepository{
		database: database,
	}
}

func (repository *groupRepository) FindAll() (*[]domain.Group, error) {
	var categories []domain.Group
	err := repository.database.Find(&categories).Error
	return &categories, err
}

func (repository *groupRepository) Create(group *domain.Group) error {
	return repository.database.Create(group).Error
}

func (repository *groupRepository) Update(group *domain.Group) error {
	return repository.database.Save(&group).Error
}

func (repository *groupRepository) Delete(id uint) error {
	return repository.database.Delete(&domain.Group{}, id).Error
}

func (repository *groupRepository) FindByID(id uint) (*domain.Group, error) {
	var group domain.Group
	err := repository.database.First(&group, id).Error
	return &group, err
}
