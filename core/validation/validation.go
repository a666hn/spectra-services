package validation

import (
	f "fmt"
	rgx "regexp"

	cp "github.com/skinnyguy/spectra-services/core/proto"
)

/* =============== */
// Role Validation //
/* =============== */

// ValidationInputRole ...
func ValidationInputRole(data *cp.InputRole) error {
	if isEmpty(data.Name) {
		return f.Errorf("Role name is required")
	} else if !isValidAlphabetWithWhitespace(data.Name) {
		return f.Errorf("Role name can't have special character : %s", data.Name)
	}

	if len(data.Description) > 0 {
		if !isValidAlphabetWithWhitespace(data.Description) {
			return f.Errorf("Role name can't have special character : %s", data.Description)
		}
	}

	return nil
}

func ValidationInputUpdateRole(data *cp.InputUpdateRole) error {
	if isEmpty(data.ID) {
		return f.Errorf("Role ID is required")
	}

	if isEmpty(data.Name) {
		return f.Errorf("Role name is required")
	} else if !isValidAlphabetWithWhitespace(data.Name) {
		return f.Errorf("Role name can't have special character : %s", data.Name)
	}

	if len(data.Description) > 0 {
		if !isValidAlphabetWithWhitespace(data.Description) {
			return f.Errorf("Role name can't have special character : %s", data.Description)
		}
	}

	return nil
}

func ValidationInputDeleteRole(data *cp.InputDeleteRole) error {
	if isEmpty(data.ID) {
		return f.Errorf("Role ID is required")
	}

	return nil
}

var alphabetPattern = "^[A-Za-z]+$"
var alphabetWithWhitespacePattern = "^[A-Za-z\\s]+$"

var alphabetRegex = rgx.MustCompile(alphabetPattern)
var alphabetWithWhitespaceRegex = rgx.MustCompile(alphabetWithWhitespacePattern)

func isValidAlphabet(alphabet string) bool {
	return alphabetRegex.MatchString(alphabet)
}

func isValidAlphabetWithWhitespace(alphabet string) bool {
	return alphabetWithWhitespaceRegex.MatchString(alphabet)
}

func isEmpty(input string) bool {
	return len(input) < 1
}
