package handler

import (
	"context"

	cp "github.com/skinnyguy/spectra-services/core/proto"
)

type PrivilegeHandler struct{}

// AddRole ...
func (*PrivilegeHandler) AddRole(
	ctx context.Context, in *cp.InputRole, out *cp.RoleResponseWithMessage,
) error {
	role, err := GetUsecase().AddRole(in)
	out.Result = role

	return err
}

// UpdateRole ...
func (*PrivilegeHandler) UpdateRole(
	ctx context.Context, in *cp.InputUpdateRole, out *cp.RoleResponseWithMessage,
) error {
	role, err := GetUsecase().UpdateRole(in)
	out.Result = role

	return err
}

// DeleteRole ...
func (*PrivilegeHandler) DeleteRole(
	ctx context.Context, in *cp.InputDeleteRole, out *cp.RoleResponseWithMessage,
) error {
	role, err := GetUsecase().DeleteRole(in)
	out.Result = role

	return err
}

func (*PrivilegeHandler) GetListRole(
	ctx context.Context, in *cp.GetPaginationRoleRequest, out *cp.RoleMultiResponse,
) error {
	roles, err := GetUsecase().GetListRole(in)
	out.Results = roles

	return err
}
