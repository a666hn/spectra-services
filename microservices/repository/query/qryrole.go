package query

/*	==========================	*/
//	Part of Privilege Services	//
/*	==========================	*/

import (
	f "fmt"

	c "github.com/skinnyguy/spectra-services/core/proto"
	u "github.com/skinnyguy/spectra-services/core/utils"
)

// AddRoleQuery ...
func (pr *Repository) AddRoleQuery(data *c.InputRole) (string, error) {
	query := `
		INSERT INTO roles(id, role, description)
		VALUES($1, $2, $3)
		RETURNING id;
	`

	uuid := u.GenerateUUID()
	roleName := u.Uppercase(data.Name)
	roleName = u.ReplaceCharacter(roleName, " ", "_")

	err := pr.Connection.ExecuteDBInsert(query, uuid, roleName, data.Description)
	if err != nil {
		return "", err
	}

	return uuid, nil
}

// UpdateRoleQuery ...
func (pr *Repository) UpdateRoleQuery(data *c.InputUpdateRole) (string, error) {
	query := `
		UPDATE roles
		SET role = $2, description = $3
		WHERE id = $1
		RETURNING id;
	`

	roleName := u.Uppercase(data.Name)
	roleName = u.ReplaceCharacter(roleName, " ", "_")

	_, err := pr.Connection.DB().Exec(query, data.ID, roleName, data.Description)
	if err != nil {
		return "", err
	}

	return data.ID, nil
}

// DeleteRoleQuery ...
func (pr *Repository) DeleteRoleQuery(data *c.InputDeleteRole) (string, error) {
	query := `
		DELETE FROM roles
		WHERE id = $1
		RETURNING id;
	`

	_, err := pr.Connection.DB().Exec(query, data.ID)
	if err != nil {
		return "", err
	}

	return data.ID, nil
}

// GetListRoleQuery ...
func (pr *Repository) GetListRoleQuery(data *c.GetPaginationRoleRequest) ([][]string, error) {
	qFilter := " WHERE "
	query := `
		SELECT
			r.id, -- 0
			r."role", -- 1
			r.description -- 2
		FROM roles r
	`

	if data.Name != "" {
		qFilter += f.Sprintf(`r."role" ILIKE '%%%%%s%%%%'`, data.Name)
		query = query + qFilter
	}

	if data.Page > 0 && data.Perpage > 0 {
		offset := (data.Page * data.Perpage) - data.Perpage
		fetch := data.Perpage

		qPagination := f.Sprintf(`
			OFFSET '%d' ROWS
			FETCH FIRST '%d' ROWS ONLY
		`, offset, fetch)

		query = query + " " + qPagination
	}

	return pr.Connection.ExecuteDBQuery(query)
}
