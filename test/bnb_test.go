// package test
package test

import (
	"github.com/nilptrr/gocat/internal/gocat/validators"
	"testing"
)

func TestBNBValidator(t *testing.T) {
	validAddresses := []string{
		"bnb1hrku3lxq2mpcl6e7jdhxlpkc0xglq3el3jjns4",
		"bnb17fk9lwcve6vufcc8tflwvref878myw9mgy3ewn",
		"bnb13eekke9gt5msqpcu3hxek3e9q5v6xk448464wz",
		"tbnb1hexqyu3m8uuudqdnnpnsnlwe6xg0n3078lx68l",
		"tbnb1r6l0c0fxu458hlq6m7amkcltj8nufyl9mr2wm5",
		"tbnb135mqtf9gef879nmjlpwz6u2fzqcw4qlzrqwgvw",
	}

	for _, address := range validAddresses {
		if validationResult := validators.Bech32Validator(address); validationResult != true {
			t.Error("Expected true, got", validationResult, "for address", address)
		}
	}

	invalidAddresses := []string{
		"bnb1hrku3lxq2mpcl6e7jdhxlpkc0xglq3el3jjns",
		"bnb1hrku3lxq2mpcl6e7jdhxlpkc0xglq3el3jjns4q",
		"bnb1hrku3lxq2mpcl6e7jdhxlpkc0xglq3el3jjns1",
		"17fk9lwcve6vufcc8tflwvref878myw9mgy3ewn",
		"tbnb1hexqyu3m8uuudqdnnpnsnlwe6xg0n3078lx68l1",
		"tbnb1r6l0c0fxu458hlq6m7amkcltj8nufyl9mr2wm",
		"tbnb135mqtf9gef879nmjlpwz6u2fzqcw4qlzrqwgv",
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
