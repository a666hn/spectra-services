package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/skinnyguy/spectra-services/api/graph"
	"github.com/skinnyguy/spectra-services/api/graph/model"
	srv "github.com/skinnyguy/spectra-services/api/services"
)

func (r *privilegeMutationResolver) AddRole(ctx context.Context, obj *model.AbstractModel, request model.InputRole) (*model.RoleResponseWithMessage, error) {
	roles, err := srv.AddRole(ctx, request)
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *privilegeMutationResolver) UpdateRole(ctx context.Context, obj *model.AbstractModel, id string, name string, description string) (*model.RoleResponseWithMessage, error) {
	roles, err := srv.UpdateRole(ctx, id, name, description)
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *privilegeMutationResolver) DeleteRole(ctx context.Context, obj *model.AbstractModel, id string) (*model.RoleResponseWithMessage, error) {
	roles, err := srv.DeleteRole(ctx, id)
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *privilegeQueryResolver) GetListRole(ctx context.Context, obj *model.AbstractModel, request model.GetPaginationRoleRequest) (*model.RoleMultiResponse, error) {
	roles, err := srv.GetListRole(ctx, request)
	if err != nil {
		return nil, err
	}

	return roles, nil
}

// PrivilegeMutation returns graph.PrivilegeMutationResolver implementation.
func (r *Resolver) PrivilegeMutation() graph.PrivilegeMutationResolver {
	return &privilegeMutationResolver{r}
}

// PrivilegeQuery returns graph.PrivilegeQueryResolver implementation.
func (r *Resolver) PrivilegeQuery() graph.PrivilegeQueryResolver { return &privilegeQueryResolver{r} }

type privilegeMutationResolver struct{ *Resolver }
type privilegeQueryResolver struct{ *Resolver }
