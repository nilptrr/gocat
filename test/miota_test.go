package test

import (
	"github.com/nilptrr/gocat/internal/gocat/validators"
	"testing"
)

func TestMIOTAValidator(t *testing.T) {
	validAddresses := []string{
		//"UDYXTZBE9GZGPM9SSQV9LTZNDLJIZMPUVVXYXFYVBLIEUHLSEWFTKZZLXYRHHWVQV9MNNX9KZC9D9UZWZRGJMIGPDW",
		"iota1qpjxtwjp4dz7lwqmtuffjcwlfzuwzuc0lx2dgdqf0wlcwp6t3r5ag8keyg7",
		"iota1qrxkunkcpmke27n9aju8h038lxqgwxm9huj4gpf63jhqkcjeuhhhjhp90zv",
		"iota1qqnhw4lx2fx8pl7m8lpm3zpehy37mckpwqeuyc3cc8aczwjxtzdlz0a2sw9",
		"atoi1qp484lf3t8pm6knjnscsr4jfv5gmrpq6rwcjf5mr36l2fzmpmlyv53mway6",
		"atoi1qz4ep7pqxuz8xhlscnzhl8meh27ymz00ctyrpxlp0ppdh6ghtrrky8079cd",
		"atoi1qr0w6u4rmy3wtw460wg8pga0nccn7el493d6387kp3ygtmq7zd3mzqfk93e",
	}

	for _, address := range validAddresses {
		if validationResult := validators.MIOTAValidator(address); validationResult != true {
			t.Error("Expected true, got", validationResult, "for address", address)
		}
	}

	invalidAddresses := []string{
		"JALLWDUOSTSJVL9EEHKW9YQFPBVBJAGLNKRVGSQZCGHQWEMIIILJMTHVAGVDXJVZMBAMOZTSBQNRVNLLSASD",
		"iota1qpjxtwjp4dz7lwqmtuffjcwlfzuwzuc0lx2dgdqf0wlcwp6t3r5ag8keygq",
		"atoi1qp484lf3t8pm6knjnscsr4jfv5gmrpq6rwcjf5mr36l2fzmpmlyv53mwaya",
		"atoi1qz4ep7pqxuz8xhlscnzhl8meh27ymz00ctyrpxlp0ppdh6ghtrrky8079c",
		"atoi1qr0w6u4rmy3wtw460wg8pga0nccn7el493d6387kp3ygtmq7zd3mzqfk93e1",
		"123adfdsafLWDUOSTSJVL9EEHKW9YQFPBVBJAGLNKRVGSQZCGHQWEMIIILJMTHVAGVDXJVZMBAMOZTSBQNRVNLLASD",
		"iota",
		"12345",
		"qwerty",
		"",
		"31hr5x7HpgUTNJsdukGEUmjNNTiyVr9aT",
		"bc1qnpgqxy7nq7zt6snx0kn76lv3z6xz5dtceupg401",
		"мой биткоин адрес",
		"比特币地址",
		"Àåæ´ýú.ü.£ßòÒÀì£Ê·¸®±ô¿¯åÇÊ¯¯ï",
		"-[:%/~%`!,)+~;{.\\-.%.**_[(\\$.",
	}

	for _, address := range invalidAddresses {
		if validationResult := validators.MIOTAValidator(address); validationResult != false {
			t.Error("Expected false, got", validationResult, "for address", address)
		}
	}
}
