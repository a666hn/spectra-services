package query

import (
	f "fmt"

	cp "github.com/skinnyguy/spectra-services/core/proto"
	ut "github.com/skinnyguy/spectra-services/core/utils"
)

// GetOrganizationQuery ...
func (pr *Repository) GetOrganizationQuery(filter *cp.InputOraganizationPagination) ([][]string, error) {
	qFilter := "WHERE "
	query := `
		SELECT
			o.id, -- 0
			o.org_name, -- 1
			o.identitynumber, -- 2
			o.orgowner, -- 3
			o.isactive, -- 4
			o.phone, -- 5
			o.lat, -- 6
			o.lng, -- 7
			o.created_at, -- 8
			o.updated_at, -- 9
			o.country -- 10
		FROM organizations o
	`

	if filter.OrgID != "" {
		if qFilter != "WHERE " {
			qFilter += " AND "
		}

		qFilter += f.Sprintf(`o.id = '%s'`, filter.OrgID)
		query += qFilter
	}

	if filter.OrgName != "" {
		if qFilter != "WHERE " {
			qFilter += " AND "
		}

		qFilter += f.Sprintf(`o.org_name ILIKE '%%%%%s%%%%'`, filter.OrgName)
		query += qFilter
	}

	if filter.Owner != "" {
		if qFilter != "WHERE " {
			qFilter += " AND "
		}

		qFilter += f.Sprintf(`o.orgowner ILIKE '%%%%%s%%%%'`, filter.Owner)
		query += qFilter
	}

	if filter.Identitynumber != "" {
		if qFilter != "WHERE " {
			qFilter += " AND "
		}

		qFilter += f.Sprintf(`o.identitynumber = %s'`, filter.Identitynumber)
		query += qFilter
	}

	if filter.Country != "" {
		if qFilter != "WHERE " {
			qFilter += " AND "
		}

		qFilter += f.Sprintf(`o.country ILIKE '%%%%%s%%%%'`, filter.Country)
		query += qFilter
	}

	if filter.Page > 0 && filter.Perpage > 0 {
		offset := (filter.Page * filter.Perpage) - filter.Perpage
		fetch := filter.Perpage

		qPagination := f.Sprintf(`
			OFFSET '%d' ROWS
			FETCH FIRST '%d' ROWS ONLY
		`, offset, fetch)

		query = query + " " + qPagination
	}

	return pr.Connection.ExecuteDBQuery(query)
}

// GetDetailOrganizationQuery ...
func (pr *Repository) GetDetailOrganizationQuery(id string) ([][]string, error) {
	query := f.Sprintf(`
		SELECT
			o.id, -- 0
			o.org_name, -- 1
			o.identitynumber, -- 2
			o.orgowner, -- 3
			o.isactive, -- 4
			o.phone, -- 5
			o.lat, -- 6
			o.lng, -- 7
			o.created_at, -- 8
			o.updated_at, -- 9
			o.country -- 10
		FROM organizations o
		WHERE o.id = '%s'
	`, id)

	return pr.Connection.ExecuteDBQuery(query)
}

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

	orgName := ut.Uppercase(data.OrgName)

	_, err := db.Exec(query, uuid, orgName, data.Identitynumber, data.Owner, true, data.Phone, data.Lat, data.Lng, data.Country)
	if err != nil {
		message = "Failed add new organization"
		return "", message, err
	}

	message = "Success add new organization with id : " + uuid
	return uuid, message, nil
}

// UpdateOrganizationQuery ...
func (pr *Repository) UpdateOrganizationQuery(id string, data *cp.InputUpdateOrganization) (string, error) {
	query := `
		UPDATE organizations
		SET org_name = $2,
			identitynumber = $3,
			phone = $4,
			lat = $5,
			lng = $6,
			country = $7,
			updated_at = NOW()
		WHERE id = $1
	`

	db := pr.Connection.DB()
	_, err := db.Exec(query, id, data.OrgName, data.Identitynumber, data.Phone, data.Lat, data.Lng, data.Country)
	if err != nil {
		return "", err
	}

	return id, nil
}
