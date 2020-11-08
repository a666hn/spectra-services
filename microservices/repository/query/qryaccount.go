package query

import (
	core "github.com/skinnyguy/spectra-services/core/proto"
	utils "github.com/skinnyguy/spectra-services/core/utils"
)

func (pr *Repository) AddAccountQuery(in *core.AddAccountRequest) (string, error) {
	query := `
		INSERT INTO accounts (id, firstname, lastname, fullname, nickname, username, password, email, phone_number, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW(), NOW())
	`

	id := utils.GenerateUUID()
	fullName := in.Firstname + in.Lastname
	passwordAfterHash := utils.GetHashPassword(in.Password)

	if err := pr.Connection.ExecuteDBInsert(query, id, in.Firstname, in.Lastname, fullName, in.Nickname, in.Username, passwordAfterHash, in.Email, in.PhoneNumber); err != nil {
		return "", err
	}

	return id, nil
}
