package services

import (
	"context"

	am "github.com/skinnyguy/spectra-services/api/graph/model"
	cl "github.com/skinnyguy/spectra-services/core/client"
	cp "github.com/skinnyguy/spectra-services/core/proto"
	u "github.com/skinnyguy/spectra-services/core/utils"
)

var (
	spectraRoleService = "Spectra.api.services."
)

// AddRole ...
func AddRole(ctx context.Context, data am.InputRole) (result *am.RoleResponseWithMessage, err error) {
	u.SendLogInfo(spectraRoleService + "AddRole")

	request := new(cp.InputRole)
	if err := u.CopyObject(data, &request); err != nil {
		return nil, err
	}

	service := cl.GetRoleService()
	response, err := service.AddRole(
		ctx,
		request,
	)

	if err != nil {
		return nil, err
	}

	result = new(am.RoleResponseWithMessage)
	err = u.CopyObject(response, &result)

	return result, err
}

// UpdateRole ...
func UpdateRole(ctx context.Context, id, name, description string) (result *am.RoleResponseWithMessage, err error) {
	u.SendLogInfo(spectraRoleService + "UpdateRole")

	request := new(cp.InputUpdateRole)
	request.ID = id
	request.Name = name
	request.Description = description

	service := cl.GetRoleService()
	response, err := service.UpdateRole(
		ctx,
		request,
	)

	if err != nil {
		return nil, err
	}

	result = new(am.RoleResponseWithMessage)
	err = u.CopyObject(response, &result)

	return result, err
}

// DeleteRole ...
func DeleteRole(ctx context.Context, id string) (result *am.RoleResponseWithMessage, err error) {
	u.SendLogInfo(spectraRoleService + "DeleteRole")

	request := new(cp.InputDeleteRole)
	request.ID = id

	service := cl.GetRoleService()
	response, err := service.DeleteRole(
		ctx,
		request,
	)

	if err != nil {
		return nil, err
	}

	result = new(am.RoleResponseWithMessage)
	err = u.CopyObject(response, &result)

	return result, err
}

// GetListRole ...
func GetListRole(ctx context.Context, data am.GetPaginationRoleRequest) (result *am.RoleMultiResponse, err error) {
	u.SendLogInfo(spectraRoleService + "GetListRole")

	request := new(cp.GetPaginationRoleRequest)
	if err := u.CopyObject(data, &request); err != nil {
		return nil, err
	}

	service := cl.GetRoleService()
	response, err := service.GetListRole(
		ctx,
		request,
	)

	if err != nil {
		return nil, err
	}

	result = new(am.RoleMultiResponse)
	err = u.CopyObject(response.Results, &result.Results)

	return result, err
}
