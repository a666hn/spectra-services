package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/skinnyguy/spectra-services/api/graph"
	"github.com/skinnyguy/spectra-services/api/graph/model"
	"github.com/skinnyguy/spectra-services/api/services"
)

func (r *familyMutationResolver) AddOrganizationFamily(ctx context.Context, obj *model.AbstractModel, request model.InputAddFamilyRequest) (*model.FamilyResponseWithMessage, error) {
	organization, err := services.AddOrganizationFamily(ctx, request)
	if err != nil {
		return nil, err
	}

	return organization, nil
}

func (r *familyQueryResolver) GetOrganizationFamily(ctx context.Context, obj *model.AbstractModel, request model.InputGetPaginationFamily) (*model.FamilyMultiResponse, error) {
	organizations, err := services.GetOrganizationFamily(ctx, request)
	if err != nil {
		return nil, err
	}

	return organizations, nil
}

// FamilyMutation returns graph.FamilyMutationResolver implementation.
func (r *Resolver) FamilyMutation() graph.FamilyMutationResolver { return &familyMutationResolver{r} }

// FamilyQuery returns graph.FamilyQueryResolver implementation.
func (r *Resolver) FamilyQuery() graph.FamilyQueryResolver { return &familyQueryResolver{r} }

type familyMutationResolver struct{ *Resolver }
type familyQueryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *familyMutationResolver) AddOrgFamily(ctx context.Context, obj *model.AbstractModel, request model.InputAddFamilyRequest) (*model.FamilyResponseWithMessage, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *familyQueryResolver) GetOrgFamily(ctx context.Context, obj *model.AbstractModel, request model.InputGetPaginationFamily) (*model.FamilyMultiResponse, error) {
	panic(fmt.Errorf("not implemented"))
}
