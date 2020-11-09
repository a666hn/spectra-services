package handler

import (
	"context"

	c "github.com/skinnyguy/spectra-services/core/proto"
)

type OrganizationFamilyHandler struct{}

// AddOrganizationFamily ...
func (*OrganizationFamilyHandler) AddOrganizationFamily(
	ctx context.Context, in *c.InputAddFamilyRequest, out *c.FamilyResponseWithMessage,
) error {
	result, err := GetUsecase().AddOrganizationFamily(in)

	out.ID = result.ID
	out.Message = result.Message

	return err
}

// GetOrganizationFamily ...
func (*OrganizationFamilyHandler) GetOrganizationFamily(
	ctx context.Context, in *c.InputGetPaginationFamily, out *c.FamilyMultiResponse,
) error {
	var results []*c.FamilyData

	orgResult, err := GetUsecase().GetOrganizationFamily(in)
	results = append(results, orgResult...)

	out.OrgFamilys = results

	return err
}
