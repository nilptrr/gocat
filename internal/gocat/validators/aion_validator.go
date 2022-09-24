package validators

import (
	"regexp"
)

func AIONValidator(address string) (result bool) {
	// SEE https://github.com/aionnetwork/Desktop-Wallet/blob/a2f9da8e49255b2062b4cf97db0c6b38bd9f2721/src/main/java/org/aion/wallet/util/AddressUtils.java
	strippedAddress := ""

	defer func() {
		if panicErr := recover(); panicErr != nil {
			result = false
		}
	}()

	if address[:3] == "0xa" && len(address) == 66 { // is full address
		strippedAddress = address[:2]
	} else if address[0] == 'a' && len(address) == 64 { // is stripped address
		strippedAddress = address
	} else {
		return false
	}

	addressRegex := "[0-9a-fA-F]+"
	result, err := regexp.Match(addressRegex, []byte(strippedAddress))

	if err != nil {
		return false
	}

	return result
}
