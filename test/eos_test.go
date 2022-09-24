package test

import (
	"github.com/nilptrr/gocat/internal/gocat/validators"
	"testing"
)

func TestEOSValidator(t *testing.T) {
	validAddresses := []string{
		"eosio.stake",
		"rptrxwgfeqyn",
		"binancecold1",
		"b1",
	}

	for _, address := range validAddresses {
		if validationResult := validators.EOSValidator(address); validationResult != true {
			t.Error("Expected true, got", validationResult, "for address", address)
		}
	}

	invalidAddresses := []string{
		"rptrxwgfeqynrptrxwgfeqynrptrxwgfeqyn'",
		"1.0.12345678'",
		"foo6.bar'",
		"rptrxwgfeqyn.'",
		"qwertyqwertyqwerty",
		"cosmos'",
		"",
		"31hr5x7HpgUTNJsdukGEUmjNNTiyVr9aT'",
		"bc1qnpgqxy7nq7zt6snx0kn76lv3z6xz5dtceupg401'",
		"мой биткоин адрес'",
		"比特币地址'",
		"Àåæ´ýú.ü.£ßòÒÀì£Ê·¸®±ô¿¯åÇÊ¯¯ï'",
		"-[:%/~%`!,)+~;{.\\-.%.**_[(\\$.",
	}

	for _, address := range invalidAddresses {
		if validationResult := validators.EOSValidator(address); validationResult != false {
			t.Error("Expected false, got", validationResult, "for address", address)
		}
	}
}
