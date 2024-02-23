package service

import (
	"errors"

	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"github.com/alexandrelamberty/macellum-api-server/pkg/repository"
)

type calendarService struct {
	repository repository.CalendarRepository
}

type CalendarService interface {
	GetAllCalendars() (*[]domain.Calendar, error)
	CreateCalendar(*domain.Calendar) error
	UpdateCalendar(calendar *domain.Calendar) error
	DeleteCalendar(id uint) error
	GetCalendarByID(id uint) (*domain.Calendar, error)
}

func NewCalendarService(r repository.CalendarRepository) CalendarService {
	return &calendarService{
		repository: r,
	}
}

func (s *calendarService) GetAllCalendars() (*[]domain.Calendar, error) {
	return s.repository.FindAll()
}

func (s *calendarService) CreateCalendar(calendar *domain.Calendar) error {
	// Check if calendar name already exists
	existingCalendar, err := s.repository.FindByName(calendar.Name)
	if err == nil && existingCalendar != nil {
		return errors.New("calendar name already exists")
	}

	return s.repository.Create(calendar)
}

func (s *calendarService) UpdateCalendar(calendar *domain.Calendar) error {
	return s.repository.Update(calendar)
}

func (s *calendarService) DeleteCalendar(id uint) error {
	return s.repository.Delete(id)
}

func (s *calendarService) GetCalendarByID(id uint) (*domain.Calendar, error) {
	return s.repository.FindByID(id)
}
