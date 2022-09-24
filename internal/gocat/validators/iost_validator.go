package validators

import "regexp"

func IOSTValidator(address string) (result bool) {
	// SEE https://developers.iost.io/docs/en/2-intro-of-iost/Account.html
	addressRegex := `^[a-z0-9_.]{5,11}(?<!\.)$`
	result, err := regexp.Match(addressRegex, []byte(address)) // TODO invalid or unsupported Perl syntax

	if err != nil {
		return false
	}

	return
}
