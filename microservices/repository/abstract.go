package repository

import (
	cp "github.com/skinnyguy/spectra-services/core/proto"
	ut "github.com/skinnyguy/spectra-services/core/utils"
	rq "github.com/skinnyguy/spectra-services/microservices/repository/query"
)

type Repository interface {
	// Role Function Services
	AddRoleQuery(data *cp.InputRole) (string, error)
	UpdateRoleQuery(data *cp.InputUpdateRole) (string, error)
	DeleteRoleQuery(data *cp.InputDeleteRole) (string, error)
	GetListRoleQuery(data *cp.GetPaginationRoleRequest) ([][]string, error)

	// Organization Function Services
	AddOrganizationQuery(data *cp.InputOrganization) (string, string, error)
}

type AbstractRepository struct {
	Repository
}

func NewRepository() *AbstractRepository {
	newRepo := &AbstractRepository{}

	if ut.GetDatasourceInfo() == ut.Postgres {
		newRepo.Repository = &rq.Repository{
			Connection: ut.GetPostgresHandler(),
		}
	} else {
		newRepo.Repository = nil
	}

	return newRepo
}
