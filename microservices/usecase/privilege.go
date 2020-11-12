package usecase

import (
	cp "github.com/skinnyguy/spectra-services/core/proto"
	u "github.com/skinnyguy/spectra-services/core/utils"
	v "github.com/skinnyguy/spectra-services/core/validation"
)

var (
	priviligeUseCasePrefix = "Privileges UseCase:"
)

// AddRole ...
func (uc *UseCase) AddRole(data *cp.InputRole) (*cp.ResponseWithMessage, error) {
	u.SendLogInfo(priviligeUseCasePrefix + "AddRole...")
	if err := v.ValidationInputRole(data); err != nil {
		return nil, err
	}

	role, err := uc.repository.AddRole(data)
	if err != nil {
		return nil, err
	}

	return role, nil
}

// UpdateRole ...
func (uc *UseCase) UpdateRole(data *cp.InputUpdateRole) (*cp.ResponseWithMessage, error) {
	u.SendLogInfo(priviligeUseCasePrefix + "UpdateRole...")
	if err := v.ValidationInputUpdateRole(data); err != nil {
		return nil, err
	}

	role, err := uc.repository.UpdateRole(data)
	if err != nil {
		return nil, err
	}

	return role, nil
}

// DeleteRole ...
func (uc *UseCase) DeleteRole(data *cp.InputDeleteRole) (*cp.ResponseWithMessage, error) {
	u.SendLogInfo(priviligeUseCasePrefix + "DeleteRole...")
	if err := v.ValidationInputDeleteRole(data); err != nil {
		return nil, err
	}

	role, err := uc.repository.DeleteRole(data)
	if err != nil {
		return nil, err
	}

	return role, nil
}

// GetListRole ...
func (uc *UseCase) GetListRole(data *cp.GetPaginationRoleRequest) ([]*cp.RoleData, error) {
	u.SendLogInfo(priviligeUseCasePrefix + "GetListRole")

	roles, err := uc.repository.GetListRole(data)
	if err != nil {
		return nil, err
	}

	return roles, nil
}
