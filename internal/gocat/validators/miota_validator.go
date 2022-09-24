package validators

import (
	"regexp"
	"strconv"
)

func MIOTAValidator(address string) (result bool) {
	// SEE https://github.com/iotaledger/iota.crypto.js/blob/5450ce54c8d9f46c3ff2d5bea2c48d9e1021da08/lib/utils/inputValidator.js
	defer func() {
		if panicErr := recover(); panicErr != nil {
			result = false
		}
	}()

	if address[:4] == "iota" || address[:4] == "atoi" { // Chrysalis address format
		return Bech32Validator(address)
	} else { // Legacy address format
		if len(address) == 90 { // Check if address with checksum
			return isTrytes(address, 90)
		} else if len(address) == 81 {
			return isTrytes(address, 81)
		}
	}

	return false
}

func isTrytes(trytes string, length int) bool {
	// Checks if input is correct trytes consisting of A-Z9 optionally validate length.
	// If no length specified, just validate the trytes.
	trytesRegex := `^[9A-Z]{` + strconv.Itoa(length) + `}$`
	matched, err := regexp.Match(trytesRegex, []byte(trytes))

	if err != nil {
		return false
	}

	return matched
}
