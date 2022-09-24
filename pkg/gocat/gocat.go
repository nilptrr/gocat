package gocat

import (
	"github.com/nilptrr/gocat/internal/gocat/validators"
)

// Validates returns the result of validating the address of the corresponding symbol (ticker).
func Validate(symbol string, address string) bool {
	switch symbol {
	case "btc":
		return validators.DefaultValidator(address)
	default:
		return false
	}
}