package usecase

import (
	cp "github.com/skinnyguy/spectra-services/core/proto"
	u "github.com/skinnyguy/spectra-services/core/utils"
	v "github.com/skinnyguy/spectra-services/core/validation"
)

var roleUseCasePrefix = "Role UseCase: "

// AddRole ...
func (uc *UseCase) AddRole(data *cp.InputRole) (*cp.ResponseWithMessage, error) {
	u.SendLogInfo(roleUseCasePrefix + "AddRole...")
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
	u.SendLogInfo(roleUseCasePrefix + "UpdateRole...")
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
	u.SendLogInfo(roleUseCasePrefix + "DeleteRole...")
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
	u.SendLogInfo(roleUseCasePrefix + "GetListRole")

	roles, err := uc.repository.GetListRole(data)
	if err != nil {
		return nil, err
	}

	return roles, nil
}
