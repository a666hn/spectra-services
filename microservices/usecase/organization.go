package usecase

import (
	cp "github.com/skinnyguy/spectra-services/core/proto"
	ut "github.com/skinnyguy/spectra-services/core/utils"
)

var prefixOrganizationUseCase = "Organization Usecase "

// GetOrganization ...
func (uc *UseCase) GetOrganization(filter *cp.InputOraganizationPagination) (total int32, results []*cp.OrganizationData, err error) {
	ut.SendLogInfo(prefixOrganizationUseCase + "GetOrganization...")
	total, results, err = uc.repository.GetOrganization(filter)
	if err != nil {
		return 0, nil, err
	}

	return total, results, nil
}

// GetDetailOrganization ...
func (uc *UseCase) GetDetailOrganization(id string) (result *cp.OrganizationData, err error) {
	ut.SendLogInfo(prefixOrganizationUseCase + "GetDetailOrganization...")
	result, err = uc.repository.GetDetailOrganization(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// AddOrganization ...
func (uc *UseCase) AddOrganization(data *cp.InputOrganization) (*cp.MessageResponse, error) {
	ut.SendLogInfo(prefixOrganizationUseCase + "AddOrganization...")
	uuid, message, err := uc.repository.AddOrganization(data)
	if err != nil {
		return nil, err
	}

	result := new(cp.MessageResponse)
	result.ID = uuid
	result.Message = message

	return result, nil
}

// UpdateOrganization ...
func (uc *UseCase) UpdateOrganization(data *cp.InputUpdateOrganization, id string) (*cp.MessageResponse, error) {
	ut.SendLogInfo(prefixOrganizationUseCase + "UpdateOrganization...")
	uuid, message, err := uc.repository.UpdateOrganization(data, id)
	if err != nil {
		return nil, err
	}

	result := new(cp.MessageResponse)
	result.ID = uuid
	result.Message = message

	return result, nil
}
