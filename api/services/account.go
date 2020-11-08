package services

import (
	"context"

	apimodel "github.com/skinnyguy/spectra-services/api/graph/model"
	coreclient "github.com/skinnyguy/spectra-services/core/client"
	coremodel "github.com/skinnyguy/spectra-services/core/proto"
	utils "github.com/skinnyguy/spectra-services/core/utils"
)

// AddAccount API Service ...
func AddAccount(ctx context.Context, obj apimodel.AddAccountRequest) (result *apimodel.AccountMessageResponse, err error) {
	utils.SendLogInfo("Spectra AddAccount API Service...")
	request := new(coremodel.AddAccountRequest)

	if err := utils.CopyObject(obj, &request); err != nil {
		return nil, err
	}

	service := coreclient.GetAccountService()
	response, err := service.AddAccount(
		ctx,
		request,
	)

	if err != nil {
		return nil, err
	}

	err = utils.CopyObject(response, &result)
	return result, err
}
