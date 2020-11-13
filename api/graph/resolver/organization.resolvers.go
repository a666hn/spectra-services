package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/skinnyguy/spectra-services/api/graph"
	"github.com/skinnyguy/spectra-services/api/graph/model"
)

func (r *organizationMutationResolver) AddOrganization(ctx context.Context, obj *model.AbstractModel, request *model.InputOrganization) (*model.OrganizationResponseWithMessage, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *organizationMutationResolver) UpdateOrganization(ctx context.Context, obj *model.AbstractModel, orgID string) (*model.OrganizationResponseWithMessage, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *organizationQueryResolver) GetOrganization(ctx context.Context, obj *model.AbstractModel, request model.InputOrganizationPagination) (*model.OrganizationMultiResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *organizationQueryResolver) GetDetailOrganization(ctx context.Context, obj *model.AbstractModel, orgID string) (*model.OrganizationData, error) {
	panic(fmt.Errorf("not implemented"))
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
