package query

import (
	core "github.com/skinnyguy/spectra-services/core/proto"
	utils "github.com/skinnyguy/spectra-services/core/utils"
)

// AddOrganizationFamilyQuery ...
func (pr *Repository) AddOrganizationFamilyQuery(data *core.InputAddFamilyRequest) (string, error) {
	query := `
		INSERT INTO family_org (id, name, identitynumber, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
	`

	id := utils.GenerateUUID()
	defaultStatus := true

	if err := pr.Connection.ExecuteDBInsert(query, id, data.Name, data.IdentityNumber, defaultStatus); err != nil {
		return "", err
	}

	return id, nil
}

// GetOrganizationFamilyQuery ...
func (pr *Repository) GetOrganizationFamilyQuery(data *core.InputGetPaginationFamily) ([][]string, error) {
	// var pagination *utils.QueryPagination

	// Get pagination parameters from request
	// pagination = new(utils.QueryPagination)
	// pagination.Page = data.Page
	// pagination.Perpage = data.Perpage

	// qFilter := "WHERE "

	// if data.ID != "" {
	// 	qFilter += f.Sprintf(`fo.id = '%s'`, data.ID)
	// }

	// if data.IdentityNumber != "" {
	// 	if qFilter != "WHERE " {
	// 		qFilter += " AND "
	// 	}

	// 	qFilter += f.Sprintf(`fo.identitynumber = '%s'`, data.IdentityNumber)
	// }

	// if data.Name != "" {
	// 	if qFilter != "WHERE " {
	// 		qFilter += " AND "
	// 	}

	// 	qFilter += f.Sprintf(`fo."name" ILIKE '%%%%%s%%%%'`, data.Name)
	// }

	// if utils.IsUsePagination(pagination) {
	// }

	// Create select query
	query := `
		SELECT
			fo.id, -- 0
			fo."name", -- 1
			fo.identitynumber, -- 2
			fo.status, -- 3
			fo.created_at, -- 4
			fo.updated_at -- 5
		FROM family_org fo
	`

	return pr.Connection.ExecuteDBQuery(query)
}
