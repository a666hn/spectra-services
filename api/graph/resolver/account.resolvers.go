package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	generated "github.com/skinnyguy/spectra-services/api/graph"
	"github.com/skinnyguy/spectra-services/api/graph/model"
)

func (r *accountMutationResolver) AddAccount(ctx context.Context, obj *model.AbstractModel, in model.AddAccountRequest) (*model.AccountMessageResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *accountMutationResolver) UpdateAccount(ctx context.Context, obj *model.AbstractModel, in model.UpdateAccountRequest) (*model.AccountMessageResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *accountMutationResolver) SoftDeleteAccount(ctx context.Context, obj *model.AbstractModel, in model.DeleteAccountRequest) (*model.AccountMessageResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *accountMutationResolver) HardDeleteAccount(ctx context.Context, obj *model.AbstractModel, in model.DeleteAccountRequest) (*model.AccountMessageResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *accountQueryResolver) GetAccount(ctx context.Context, obj *model.AbstractModel, filter model.GetAccountRequest) (*model.AccountResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *accountQueryResolver) GetPaginationAccounts(ctx context.Context, obj *model.AbstractModel, filter model.GetPaginationAccount) (*model.AccountPaginationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

// AccountMutation returns generated.AccountMutationResolver implementation.
func (r *Resolver) AccountMutation() generated.AccountMutationResolver {
	return &accountMutationResolver{r}
}

// AccountQuery returns generated.AccountQueryResolver implementation.
func (r *Resolver) AccountQuery() generated.AccountQueryResolver { return &accountQueryResolver{r} }

type accountMutationResolver struct{ *Resolver }
type accountQueryResolver struct{ *Resolver }
