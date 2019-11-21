package cla_groups

import (
	"github.com/communitybridge/easycla-api/gen/models"
	"github.com/communitybridge/easycla-api/gen/restapi/operations/cla_groups"
)

// Service interface defines methods of cla_groups service
type Service interface {
	CreateCLAGroup(in *cla_groups.CreateCLAGroupParams) (*models.ClaGroup, error)
	DeleteCLAGroup(in *cla_groups.DeleteCLAGroupParams) error
	UpdateCLAGroup(in *cla_groups.UpdateCLAGroupParams) error
	ListCLAGroups(in *cla_groups.ListCLAGroupsParams) (*models.ClaGroupList, error)
}

type service struct {
	repo Repository
}

// NewService creates new instance of event service
func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) CreateCLAGroup(in *cla_groups.CreateCLAGroupParams) (*models.ClaGroup, error) {
	return s.repo.CreateCLAGroup(in.ClaGroup)
}

func (s *service) DeleteCLAGroup(in *cla_groups.DeleteCLAGroupParams) error {
	return s.repo.DeleteCLAGroup(in.ClaGroupID)
}

func (s *service) UpdateCLAGroup(in *cla_groups.UpdateCLAGroupParams) error {
	return s.repo.UpdateCLAGroup(in.ClaGroupID, in.ClaGroup)
}

func (s *service) ListCLAGroups(in *cla_groups.ListCLAGroupsParams) (*models.ClaGroupList, error) {
	return s.repo.ListCLAGroups(in)
}
