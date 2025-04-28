package repository

import (
	"github.com/cukhoaimon/khoainats/internal/repository/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrganizationRepository interface {
	ById(id uuid.UUID) (model.Organization, error)
	Create(org model.Organization) (model.Organization, error)
}

func NewOrganizationRepository(db *gorm.DB) OrganizationRepository {
	return organizationRepository{db: db}
}

type organizationRepository struct {
	db *gorm.DB
}

func (r organizationRepository) ById(id uuid.UUID) (model.Organization, error) {
	var org model.Organization
	if result := r.db.First(&org, id); result.Error != nil {
		return model.Organization{}, result.Error
	}

	return org, nil
}

func (r organizationRepository) Create(org model.Organization) (model.Organization, error) {
	if result := r.db.Create(org); result.Error != nil {
		return model.Organization{}, result.Error
	}

	return org, nil
}
