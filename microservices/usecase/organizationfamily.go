package usecase

import (
	c "github.com/skinnyguy/spectra-services/core/proto"
	u "github.com/skinnyguy/spectra-services/core/utils"
)

const FamilyOrganizationUseCaseMessage string = "Family Organization Use Case: "

// AddOrganizationFamily ...
func (uc *UseCase) AddOrganizationFamily(data *c.InputAddFamilyRequest) (*c.FamilyResponseWithMessage, error) {
	u.SendLogInfo(FamilyOrganizationUseCaseMessage + "AddOrganizationFamily...")
	return uc.repository.AddOrganizationFamily(data)
}

// GetOrganizationFamily ...
func (uc *UseCase) GetOrganizationFamily(data *c.InputGetPaginationFamily) ([]*c.FamilyData, error) {
	u.SendLogInfo(FamilyOrganizationUseCaseMessage + "GetOrganizationFamily...")
	return uc.repository.GetOrganizationFamily(data)
}
