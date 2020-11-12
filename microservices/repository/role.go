package repository

import (
	cp "github.com/skinnyguy/spectra-services/core/proto"
	u "github.com/skinnyguy/spectra-services/core/utils"
)

// AddRole ...
func (ar *AbstractRepository) AddRole(data *cp.InputRole) (*cp.ResponseWithMessage, error) {
	var (
		result *cp.ResponseWithMessage
		uuid   string
		err    error

		errText = "Service.Role.err"
	)

	if uuid, err = ar.AddRoleQuery(data); err != nil {
		return nil, u.SendLogError(errText+"AddRoleQuery", err)
	}

	result = new(cp.ResponseWithMessage)

	result.ID = uuid
	result.Message = "Success adding new Role"

	return result, nil
}

// UpdateRole ...
func (ar *AbstractRepository) UpdateRole(data *cp.InputUpdateRole) (*cp.ResponseWithMessage, error) {
	var (
		result *cp.ResponseWithMessage
		uuid   string
		err    error

		errText = "Service.Role.err"
	)

	if uuid, err = ar.UpdateRoleQuery(data); err != nil {
		return nil, u.SendLogError(errText+"UpdateRoleQuery", err)
	}

	result = new(cp.ResponseWithMessage)

	result.ID = uuid
	result.Message = "Success updating Role"

	return result, nil
}

// DeleteRole ...
func (ar *AbstractRepository) DeleteRole(data *cp.InputDeleteRole) (*cp.ResponseWithMessage, error) {
	var (
		result  *cp.ResponseWithMessage
		uuid string
		err     error

		errText = "Service.Role.err"
	)

	if uuid, err = ar.DeleteRoleQuery(data); err != nil {
		return nil, u.SendLogError(errText+"DeleteRoleQuery", err)
	}

	result = new(cp.ResponseWithMessage)

	result.ID = uuid
	result.Message = "Success deleting Role"

	return result, nil
}

// GetListRole ...
func (ar *AbstractRepository) GetListRole(data *cp.GetPaginationRoleRequest) ([]*cp.RoleData, error) {
	var (
		records [][]string
		results []*cp.RoleData
		err     error
		errText = "Service.Role.err"
	)

	if records, err = ar.GetListRoleQuery(data); err != nil {
		return nil, u.SendLogError(errText+"GetListRoleQuery", err)
	}

	for i := range records {
		if i != 0 {
			result := new(cp.RoleData)

			result.ID = records[i][0]
			result.Role = records[i][1]
			result.Description = records[i][2]

			results = append(results, result)
		}
	}

	return results, nil
}
