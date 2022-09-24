package validators

import (
	"crypto/sha256"
	"github.com/btcsuite/btcutil/base58"
	"github.com/btcsuite/btcutil/bech32"
	"github.com/tnakagawa/goref/bech32m"
	"reflect"
)

func DefaultValidator(address string) bool {
	if validationResult := Base58Validator(address); validationResult {
		return validationResult
	} else if validationResult := Bech32Validator(address); validationResult {
		return validationResult
	} else {
		return false
	}
}

func Base58Validator(address string) bool {
	// SEE https://en.bitcoin.it/wiki/Base58Check_encoding
	decoded := base58.Decode(address)
	embedded := getEmbeddedChecksum(decoded)
	computed := getComputedChecksum(decoded)

	if len(embedded) > 0 && len(computed) > 0 && reflect.DeepEqual(embedded, computed) {
		return true
	}

	return false
}

func Bech32Validator(address string) bool {
	// supported address prefixes
	var bech32Prefixes = [...]string{
		"bc",
		"cosmos",
		"bnb",
		"tbnb",
		"iota",
		"atoi",
	}

	// bech32 decoding
	// SEE https://github.com/bitcoin/bips/blob/1f0b563738199ca60d32b4ba779797fc97d040fe/bip-0350.mediawiki
	hrp, decoded, err := bech32.Decode(address)

	for _, prefix := range bech32Prefixes {
		if hrp == prefix && len(decoded) > 0 && err == nil {
			return true
		}
	}

	// bech32m decoding
	// SEE https://github.com/bitcoin/bips/blob/1f0b563738199ca60d32b4ba779797fc97d040fe/bip-0350.mediawiki
	hrp, decoded, spec, err := bech32m.Decode(address)

	for _, prefix := range bech32Prefixes {
		if hrp == prefix && len(decoded) > 0 && spec != -1 && err == nil {
			return true
		}
	}

	return false
}

func getEmbeddedChecksum(decoded []byte) (result []byte) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			result = []byte{}
		}
	}()

	return decoded[21:]
}

func getComputedChecksum(decoded []byte) (hashed []byte) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			hashed = []byte{}
		}
	}()

	hashed = doubleSHA256(decoded[:21])
	return hashed[:4]
}

func doubleSHA256(data []byte) []byte {
	first := sha256.Sum256(data)
	second := sha256.Sum256(first[:])

	return second[:]
}
