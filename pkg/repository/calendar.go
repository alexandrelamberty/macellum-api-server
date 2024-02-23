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

func (ur *calendarRepository) FindAll() (*[]domain.Calendar, error) {
	var calendars []domain.Calendar
	err := ur.database.Find(&calendars).Error
	return &calendars, err
}

func (ur *calendarRepository) Create(calendar *domain.Calendar) error {
	return ur.database.Create(calendar).Error
}

func (ur *calendarRepository) Update(calendar *domain.Calendar) error {
	return ur.database.Save(&calendar).Error
}

func (r *calendarRepository) Delete(id uint) error {
	return r.database.Delete(&domain.Calendar{}, id).Error
}

func (r *calendarRepository) FindByID(id uint) (*domain.Calendar, error) {
	var calendar domain.Calendar
	err := r.database.First(&calendar, id).Error
	return &calendar, err
}

func (r *calendarRepository) FindByName(name string) (*domain.Calendar, error) {
	var calendar domain.Calendar
	err := r.database.Model(domain.User{}).Where("name = ?", name).First(&calendar).Error
	return &calendar, err
}
