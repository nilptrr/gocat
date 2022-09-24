package test

import (
	"github.com/nilptrr/gocat/internal/gocat/validators"
	"testing"
)

func TestIOSTValidator(t *testing.T) {
	validAddresses := []string{
		"base.iost",
		"iostliebibp",
		"infstones",
		"neoply_iost",
	}

	for _, address := range validAddresses {
		if validationResult := validators.IOSTValidator(address); validationResult != true {
			t.Error("Expected true, got", validationResult, "for address", address)
		}
	}

	invalidAddresses := []string{
		"qwertyqwertyqwerty",
		"8mc.",
		"Lqqdouge",
		"",
		"31hr5x7HpgUTNJsdukGEUmjNNTiyVr9aT",
		"bc1qnpgqxy7nq7zt6snx0kn76lv3z6xz5dtceupg401",
		"мой биткоин адрес",
		"比特币地址",
		"Àåæ´ýú.ü.£ßòÒÀì£Ê·¸®±ô¿¯åÇÊ¯¯ï",
		"-[:%/~%`!,)+~;{.\\-.%.**_[(\\$",
	}

	for _, address := range invalidAddresses {
		if validationResult := validators.IOSTValidator(address); validationResult != false {
			t.Error("Expected false, got", validationResult, "for address", address)
		}
	}
}
