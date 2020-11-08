package usecase

import (
	core "github.com/skinnyguy/spectra-services/core/proto"
	utils "github.com/skinnyguy/spectra-services/core/utils"
)

const AccountUseCaseMessage string = "Account Use Case: "

func (uc *UseCase) AddAccount(data *core.AddAccountRequest) (*core.AccountMessageResponse, error) {
	utils.SendLogInfo(AccountUseCaseMessage + "AddAccount...")
	return uc.repository.AddAccount(data)
}
