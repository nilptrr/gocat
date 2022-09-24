package validators

import "regexp"

func EOSValidator(address string) (result bool) {
	// SEE https://developers.eos.io/welcome/v2.1/protocol-guides/accounts_and_permissions/#2-accounts
	addressRegex := `^[a-z1-5.]{1,12}(?<!\.)$`
	result, err := regexp.Match(addressRegex, []byte(address)) // TODO invalid or unsupported Perl syntax

	if err != nil {
		return false
	}

	return
}
