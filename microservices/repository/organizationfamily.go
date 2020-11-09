package repository

import (
	c "github.com/skinnyguy/spectra-services/core/proto"
	ut "github.com/skinnyguy/spectra-services/core/utils"
)

const RepositoryKeyOrg string = "OrganizationFamily."
const errTextOrg string = "Service.OrganizationFamily.err"

// AddOrganizationFamily ...
func (ar *AbstractRepository) AddOrganizationFamily(data *c.InputAddFamilyRequest) (*c.FamilyResponseWithMessage, error) {
	ut.SendLogDebug(RepositoryKeyOrg + "AddOrganizationFamily...")
	var (
		id      string
		message string
		result  *c.FamilyResponseWithMessage
		err     error
	)

	id, err = ar.AddOrganizationFamilyQuery(data)
	if err != nil {
		message = "Failed to save Organization Family"
		return nil, ut.SendLogError(errTextOrg+"AddOrganizationFamilyQuery", err)
	}

	message = "Success Save New Organization Family"
	result = new(c.FamilyResponseWithMessage)

	result.ID = id
	result.Message = message

	return result, nil
}

// GetOrganizationFamily ...
func (ar *AbstractRepository) GetOrganizationFamily(data *c.InputGetPaginationFamily) ([]*c.FamilyData, error) {
	ut.SendLogDebug(RepositoryKeyOrg + "GetOrganizationFamily...")
	var (
		records [][]string
		results []*c.FamilyData
		err     error
	)

	if records, err = ar.GetOrganizationFamilyQuery(data); err != nil {
		return nil, ut.SendLogError(errTextOrg+"GetOrganizationFamilyQuery", err)
	}

	for i := range records {
		if i != 0 {
			result := new(c.FamilyData)

			status, _ := ut.StringToBoolean(records[i][3])

			result.ID = records[i][0]
			result.Name = records[i][1]
			result.IdentityNumber = records[i][2]
			result.Status = status
			result.CreatedAt = records[i][4]
			result.UpdatedAt = records[i][5]

			results = append(results, result)
		}
	}

	return results, nil
}
