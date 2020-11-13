package handler

import (
	"context"

	cp "github.com/skinnyguy/spectra-services/core/proto"
)

type OrganizationHandler struct{}

// GetOrganization ...
func (*OrganizationHandler) GetOrganization(
	ctx context.Context, in *cp.InputOraganizationPagination, out *cp.OrganizationMultiResponse,
) error {
	total, organizations, err := GetUsecase().GetOrganization(in)
	out.Page = in.Page
	out.Perpage = in.Perpage
	out.TotalDocument = total
	out.Results = organizations

	return err
}

// GetDetailOrganization ...
func (*OrganizationHandler) GetDetailOrganization(
	ctx context.Context, in *cp.OrganizationDefaultRequest, out *cp.OrganizationSingleResponse,
) error {
	organization, err := GetUsecase().GetDetailOrganization(in.OrgID)
	out.Result = organization

	return err
}

// AddOrganization ...
func (*OrganizationHandler) AddOrganization(
	ctx context.Context, in *cp.InputOrganization, out *cp.ResponseMessage,
) error {
	org, err := GetUsecase().AddOrganization(in)
	out.Result = org

	return err
}

// UpdateOrganization ...
func (*OrganizationHandler) UpdateOrganization(
	ctx context.Context, in *cp.InputUpdateOrganization, out *cp.ResponseMessage,
) error {
	organization, err := GetUsecase().UpdateOrganization(in, in.OrgID)
	out.Result = organization
	return err
}
