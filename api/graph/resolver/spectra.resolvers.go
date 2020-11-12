package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	generated "github.com/skinnyguy/spectra-services/api/graph"
	"github.com/skinnyguy/spectra-services/api/graph/model"
)

func (r *mutationResolver) Account(ctx context.Context) (*model.AbstractModel, error) {
	return new(model.AbstractModel), nil
}

func (r *mutationResolver) Family(ctx context.Context) (*model.AbstractModel, error) {
	return new(model.AbstractModel), nil
}

func (r *mutationResolver) Privileges(ctx context.Context) (*model.AbstractModel, error) {
	return new(model.AbstractModel), nil
}

func (r *queryResolver) Account(ctx context.Context) (*model.AbstractModel, error) {
	return new(model.AbstractModel), nil
}

func (r *queryResolver) Family(ctx context.Context) (*model.AbstractModel, error) {
	return new(model.AbstractModel), nil
}

func (r *queryResolver) Privileges(ctx context.Context) (*model.AbstractModel, error) {
	return new(model.AbstractModel), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) Role(ctx context.Context) (*model.AbstractModel, error) {
	return new(model.AbstractModel), nil
}
func (r *queryResolver) Role(ctx context.Context) (*model.AbstractModel, error) {
	return new(model.AbstractModel), nil
}
