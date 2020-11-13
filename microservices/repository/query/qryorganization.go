package query

import (
	cp "github.com/skinnyguy/spectra-services/core/proto"
	ut "github.com/skinnyguy/spectra-services/core/utils"
)

// AddOrganizationQuery ...
func (pr *Repository) AddOrganizationQuery(data *cp.InputOrganization) (string, string, error) {
	var (
		query   string
		message string
	)

	query = `
		INSERT INTO organizations (id, org_name, identitynumber, orgowner, isactive, phone, lat, lng, created_at, updated_at, country)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), NOW(), $9)
	`

	uuid := ut.GenerateUUID()
	db := pr.Connection.DB()

	_, err := db.Exec(query, uuid, data.OrgName, data.Identitynumber, data.Owner, true, data.Phone, data.Lat, data.Lng, data.Country)
	if err != nil {
		message = "Failed add new organization"
		return "", message, err
	}

	message = "Success add new organization with id : " + uuid
	return uuid, message, nil
}
