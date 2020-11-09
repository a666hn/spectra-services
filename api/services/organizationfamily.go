package services

import (
	"context"

	apimodel "github.com/skinnyguy/spectra-services/api/graph/model"
	coreclient "github.com/skinnyguy/spectra-services/core/client"
	coremodel "github.com/skinnyguy/spectra-services/core/proto"
	utils "github.com/skinnyguy/spectra-services/core/utils"
)

// AddOrganizationFamily ...
func AddOrganizationFamily(ctx context.Context, data apimodel.InputAddFamilyRequest) (result *apimodel.FamilyResponseWithMessage, err error) {
	utils.SendLogInfo("Spectra AddOrganizationFamily API Service...")
	request := new(coremodel.InputAddFamilyRequest)

	if err := utils.CopyObject(data, &request); err != nil {
		return nil, err
	}

	service := coreclient.GetOrganizationFamilyService()
	response, err := service.AddOrganizationFamily(
		ctx,
		request,
	)

	if err != nil {
		return nil, err
	}

	result = new(apimodel.FamilyResponseWithMessage)
	err = utils.CopyObject(response, &result)

	return result, err
}

// GetOrganizationFamily ...
func GetOrganizationFamily(ctx context.Context, data apimodel.InputGetPaginationFamily) (results *apimodel.FamilyMultiResponse, err error) {
	utils.SendLogInfo("Spectra GetOrganizationFamily API Service...")
	request := new(coremodel.InputGetPaginationFamily)

	if err := utils.CopyObject(data, &request); err != nil {
		return nil, err
	}

	service := coreclient.GetOrganizationFamilyService()
	response, err := service.GetOrganizationFamily(
		ctx,
		request,
	)

	if err != nil {
		return nil, err
	}

	results = new(apimodel.FamilyMultiResponse)
	err = utils.CopyObject(response, &results)

	return results, err
}
