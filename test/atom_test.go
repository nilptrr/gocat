// package test
package test

import (
	"github.com/nilptrr/gocat/internal/gocat/validators"
	"testing"
)

func TestATOMValidator(t *testing.T) {
	validAddresses := []string{
		"cosmos16na5gpcj80tafv5gycm4gk7garj8jjsgydtkmj",
		"cosmos1la4t7jee73a4zqg9ccs2acjt462qjs5xkls4n7",
		"cosmos1wl34gmc5taa5h8dcr8rs3cmjdu7h0cy5cm2e39",
	}

	for _, address := range validAddresses {
		if validationResult := validators.Bech32Validator(address); validationResult != true {
			t.Error("Expected true, got", validationResult, "for address", address)
		}
	}

	invalidAddresses := []string{
		"cosmos16na5gpcj80tafv5gycm4gk7garj8jjsgydtkmq",
		"cosmos",
		"6026aeb850f355f63778fe86fefa9cf670e89e6d2a0a2be992c34c12009fa3ab",
		"12345",
		"qwerty",
		"",
		"31hr5x7HpgUTNJsdukGEUmjNNTiyVr9aT",
		"bc1qnpgqxy7nq7zt6snx0kn76lv3z6xz5dtceupg401",
		"мой биткоин адрес",
		"比特币地址",
		"Àåæ´ýú.ü.£ßòÒÀì£Ê·¸®±ô¿¯åÇÊ¯¯ï",
		`-[:%/~%!,)+~;{.\\-.%.**_[(\\$".`,
	}

	for _, address := range invalidAddresses {
		if validationResult := validators.Bech32Validator(address); validationResult != false {
			t.Error("Expected false, got", validationResult, "for address", address)
		}
	}
}
