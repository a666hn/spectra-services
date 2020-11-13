package services

import (
	"context"

	am "github.com/skinnyguy/spectra-services/api/graph/model"
	cl "github.com/skinnyguy/spectra-services/core/client"
	cp "github.com/skinnyguy/spectra-services/core/proto"
	u "github.com/skinnyguy/spectra-services/core/utils"
)

var (
	spectraOrganization = "Spectra.api.Organization.services."
)

// GetOrganization ...
func GetOrganization(ctx context.Context, filter am.InputOrganizationPagination) (result *am.OrganizationMultiResponse, err error) {
	u.SendLogInfo(spectraOrganization + "GetOrganization")

	reqFilter := new(cp.InputOraganizationPagination)
	if err := u.CopyObject(filter, &reqFilter); err != nil {
		return nil, err
	}

	service := cl.GetOrganizationService()
	response, err := service.GetOrganization(
		ctx,
		reqFilter,
	)

	if err != nil {
		return nil, err
	}

	result = new(am.OrganizationMultiResponse)
	err = u.CopyObject(response, &result)

	return result, err
}

// GetDetailOraganization ...
func GetDetailOraganization(ctx context.Context, id string) (result *am.OrganizationData, err error) {
	u.SendLogInfo(spectraOrganization + "GetDetailOraganization")

	request := &cp.OrganizationDefaultRequest{
		OrgID: id,
	}

	service := cl.GetOrganizationService()
	response, err := service.GetDetailOrganization(
		ctx,
		request,
	)

	if err != nil {
		return nil, err
	}

	result = new(am.OrganizationData)
	err = u.CopyObject(response.Result, &result)

	return result, err
}

// AddOrganization ...
func AddOrganization(ctx context.Context, data *am.InputOrganization) (result *am.OrganizationResponseWithMessage, err error) {
	u.SendLogInfo(spectraOrganization + "AddOrganization")

	request := new(cp.InputOrganization)
	if err := u.CopyObject(data, &request); err != nil {
		return nil, err
	}

	service := cl.GetOrganizationService()
	response, err := service.AddOrganization(
		ctx,
		request,
	)

	if err != nil {
		return nil, err
	}

	result = new(am.OrganizationResponseWithMessage)
	err = u.CopyObject(response, &result)

	return result, err
}

// UpdateOrganization ...
func UpdateOrganization(ctx context.Context, id string, data am.UpdateOrganization) (result *am.OrganizationResponseWithMessage, err error) {
	u.SendLogInfo(spectraOrganization + "UpdateOrganization")

	request := new(cp.InputUpdateOrganization)
	if err := u.CopyObject(data, &request); err != nil {
		return nil, err
	}

	request.OrgID = id

	service := cl.GetOrganizationService()
	response, err := service.UpdateOrganization(
		ctx,
		request,
	)

	if err != nil {
		return nil, err
	}

	result = new(am.OrganizationResponseWithMessage)
	err = u.CopyObject(response, &result)

	return result, err
}
