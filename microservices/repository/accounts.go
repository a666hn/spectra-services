package repository

import (
	core "github.com/skinnyguy/spectra-services/core/proto"
	"github.com/skinnyguy/spectra-services/core/utils"
)

const RepositoryKey string = "Account."
const errText string = "Service.Account.err"

// AddAccount ...
func (ar *AbstractRepository) AddAccount(data *core.AddAccountRequest) (*core.AccountMessageResponse, error) {
	utils.SendLogDebug(RepositoryKey + "AddAccount Query ...")
	var (
		accountId string
		err       error
		message   string
		response  *core.AccountMessageResponse
	)

	if accountId, err = ar.AddAccountQuery(data); err != nil {
		utils.SendLogError(errText+"AddAccountQuery", err)
		message = "Query Failed"
	}

	message = "Success"
	response = new(core.AccountMessageResponse)

	response.ID = accountId
	response.Message = message

	return response, nil
}
