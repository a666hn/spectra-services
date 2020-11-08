package handler

import (
	"context"

	core "github.com/skinnyguy/spectra-services/core/proto"
)

type AccountHandler struct{}

func (*AccountHandler) AddAccount(
	ctx context.Context, in *core.AddAccountRequest, out *core.AccountMessageResponse,
) error {
	account, err := GetUsecase().AddAccount(in)
	out.ID = account.GetID()
	out.Message = account.GetMessage()

	return err
}
