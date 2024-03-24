package repository

import (
	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"gorm.io/gorm"
)

type calendarRepository struct {
	database *gorm.DB
}

type CalendarRepository interface {
	FindAll() (*[]domain.Calendar, error)
	Create(calendar *domain.Calendar) error
	Update(calendar *domain.Calendar) error
	Delete(id uint) error
	FindByID(id uint) (*domain.Calendar, error)
	FindByName(name string) (*domain.Calendar, error)
}

func NewCalendarRepository(database *gorm.DB) CalendarRepository {
	return &calendarRepository{
		database: database,
	}
}

func (repository *calendarRepository) FindAll() (*[]domain.Calendar, error) {
	var calendars []domain.Calendar
	err := repository.database.Find(&calendars).Error
	return &calendars, err
}

func (repository *calendarRepository) Create(calendar *domain.Calendar) error {
	return repository.database.Create(calendar).Error
}

func (repository *calendarRepository) Update(calendar *domain.Calendar) error {
	return repository.database.Save(&calendar).Error
}

func (repository *calendarRepository) Delete(id uint) error {
	return repository.database.Delete(&domain.Calendar{}, id).Error
}

func (repository *calendarRepository) FindByID(id uint) (*domain.Calendar, error) {
	var calendar domain.Calendar
	err := repository.database.First(&calendar, id).Error
	return &calendar, err
}

func (repository *calendarRepository) FindByName(name string) (*domain.Calendar, error) {
	var calendar domain.Calendar
	err := repository.database.Model(domain.User{}).Where("name = ?", name).First(&calendar).Error
	return &calendar, err
}
