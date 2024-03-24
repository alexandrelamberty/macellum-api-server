package service

import (
	"github.com/alexandrelamberty/macellum-api-server/pkg/domain"
	"github.com/alexandrelamberty/macellum-api-server/pkg/repository"
)

type groupService struct {
	repository repository.GroupRepository
}

type GroupService interface {
	GetAllGroups() (*[]domain.Group, error)
	CreateGroup(*domain.Group) error
	UpdateGroup(group *domain.Group) error
	DeleteGroup(id uint) error
	GetGroupByID(id uint) (*domain.Group, error)
}

func NewGroupService(r repository.GroupRepository) GroupService {
	return &groupService{
		repository: r,
	}
}

func (s *groupService) GetAllGroups() (*[]domain.Group, error) {
	return s.repository.FindAll()
}

func (s *groupService) CreateGroup(group *domain.Group) error {
	return s.repository.Create(group)
}

func (s *groupService) UpdateGroup(group *domain.Group) error {
	return s.repository.Update(group)
}

func (s *groupService) DeleteGroup(id uint) error {
	return s.repository.Delete(id)
}

func (s *groupService) GetGroupByID(id uint) (*domain.Group, error) {
	return s.repository.FindByID(id)
}
