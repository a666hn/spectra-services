package repository

import (
	f "fmt"

	cp "github.com/skinnyguy/spectra-services/core/proto"
	ut "github.com/skinnyguy/spectra-services/core/utils"
)

// GetOrganization ...
func (ar *AbstractRepository) GetOrganization(data *cp.InputOraganizationPagination) (int32, []*cp.OrganizationData, error) {
	var (
		total   int32
		records [][]string
		results []*cp.OrganizationData
		err     error
		errText = "Spectra.service.organization.err"
	)

	if records, err = ar.GetOrganizationQuery(data); err != nil {
		return 0, nil, ut.SendLogError(errText+"GetOrganizationQuery", err)
	}

	f.Println("Document Length -> ", len(records))
	for i := range records {
		if i != 0 {
			result := new(cp.OrganizationData)
			orgStatus, _ := ut.StringToBoolean(records[i][4])

			result.OrgID = records[i][0]
			result.OrgName = records[i][1]
			result.Identitynumber = records[i][2]
			result.Owner = records[i][3]
			result.IsActive = orgStatus
			result.Phone = records[i][5]
			result.Lat = records[i][6]
			result.Lng = records[i][7]
			result.CreatedAt = records[i][8]
			result.UpdatedAt = records[i][9]
			result.Country = records[i][10]

			results = append(results, result)
		}
	}

	total = int32(len(results))
	return total, results, nil
}

// GetDetailOrganization ...
func (ar *AbstractRepository) GetDetailOrganization(id string) (*cp.OrganizationData, error) {
	var (
		records [][]string
		result  *cp.OrganizationData
		err     error
		errText = "Spectra.service.organization.err"
	)

	if records, err = ar.GetDetailOrganizationQuery(id); err != nil {
		return nil, ut.SendLogError(errText+"GetDetailOrganizationQuery", err)
	}

	for i := range records {
		if i != 0 {
			result = new(cp.OrganizationData)

			orgStatus, _ := ut.StringToBoolean(records[i][4])

			result.OrgID = records[i][0]
			result.OrgName = records[i][1]
			result.Identitynumber = records[i][2]
			result.Owner = records[i][3]
			result.IsActive = orgStatus
			result.Phone = records[i][5]
			result.Lat = records[i][6]
			result.Lng = records[i][7]
			result.CreatedAt = records[i][8]
			result.UpdatedAt = records[i][9]
			result.Country = records[i][10]
		}
	}

	return result, nil
}

// AddOrganization ...
func (ar *AbstractRepository) AddOrganization(data *cp.InputOrganization) (string, string, error) {
	var (
		message string
		uuid    string
		err     error
		errText = "service.organization.err"
	)

	if uuid, message, err = ar.AddOrganizationQuery(data); err != nil {
		return "", message, ut.SendLogError(errText+"AddOrganizationQuery", err)
	}

	return uuid, message, nil
}

// UpdateOrganization ...
func (ar *AbstractRepository) UpdateOrganization(data *cp.InputUpdateOrganization, id string) (string, string, error) {
	var (
		uuid, message string
		err           error
		errText       = "Spectra.service.organization.err"
	)

	if uuid, err = ar.UpdateOrganizationQuery(id, data); err != nil {
		message = "Failed to update organization"
		return "", message, ut.SendLogError(errText+"UpdateOrganizationQuery", err)
	}

	message = "Success update organization : " + uuid
	return uuid, message, nil
}
