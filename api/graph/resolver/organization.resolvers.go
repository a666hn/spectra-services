package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/skinnyguy/spectra-services/api/graph"
	"github.com/skinnyguy/spectra-services/api/graph/model"
	"github.com/skinnyguy/spectra-services/api/services"
)

func (r *organizationMutationResolver) AddOrganization(ctx context.Context, obj *model.AbstractModel, request *model.InputOrganization) (*model.OrganizationResponseWithMessage, error) {
	organization, err := services.AddOrganization(ctx, request)
	if err != nil {
		return nil, err
	}

	return organization, nil
}

func (r *organizationMutationResolver) UpdateOrganization(ctx context.Context, obj *model.AbstractModel, orgID string, request model.UpdateOrganization) (*model.OrganizationResponseWithMessage, error) {
	organization, err := services.UpdateOrganization(ctx, orgID, request)
	if err != nil {
		return nil, err
	}

	return organization, nil
}

func (r *organizationQueryResolver) GetOrganization(ctx context.Context, obj *model.AbstractModel, request model.InputOrganizationPagination) (*model.OrganizationMultiResponse, error) {
	organization, err := services.GetOrganization(ctx, request)
	if err != nil {
		return nil, err
	}

	return organization, nil
}

func (r *organizationQueryResolver) GetDetailOrganization(ctx context.Context, obj *model.AbstractModel, orgID string) (*model.OrganizationData, error) {
	organization, err := services.GetDetailOraganization(ctx, orgID)
	if err != nil {
		return nil, err
	}

	return organization, nil
}

// OrganizationMutation returns graph.OrganizationMutationResolver implementation.
func (r *Resolver) OrganizationMutation() graph.OrganizationMutationResolver {
	return &organizationMutationResolver{r}
}

// OrganizationQuery returns graph.OrganizationQueryResolver implementation.
func (r *Resolver) OrganizationQuery() graph.OrganizationQueryResolver {
	return &organizationQueryResolver{r}
}

type organizationMutationResolver struct{ *Resolver }
type organizationQueryResolver struct{ *Resolver }
