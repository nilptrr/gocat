package test

import (
	"github.com/nilptrr/gocat/internal/gocat/validators"
	"testing"
)

func TestBTCValidator(t *testing.T) {
	validAddresses := []string{
		"15e15hWo6CShMgbAfo8c2Ykj4C6BLq6Not",
		"35PBEaofpUeH8VnnNSorM1QZsadrZoQp4N",
		"bc1q42lja79elem0anu8q8s3h2n687re9jax556pcc",
		"31hr5x7HpgUTNJsdukGEUmjNNTiyVr9aT1",
		"1GRKR41WzKpcrD5RpUeMQ1NK6JG46sAtFh",
		"bc1qnpgqxy7nq7zt6snx0kn76lv3z6xz5dtceupg40",
		"bc1q7w9p2unf3hngtzmq3sq47cjdd0xd9sw7us5h6p",
		"bc1p7qxqlecjcqej0564745xl42lmm7kddgqp5llhjd3w32n4vxmnduqj9n6hp",
		"bc1plyq2dmxqtmwnvmfpk47pygp6426vjj2sj5e9n54x96ge967rrpgq27jc2v",
		"bc1p3nhp6eawl9a9lxaydlkl9ewhumatgt9v9trcnuutexda86gz3lwqsw3yxf",
		"bc1psx9ukcl7ft7aadjsma6k6xld8yhgkhjv72qs246hstc4ayzhxwwq3ju86w",
		"bc1plyq2dmxqtmwnvmfpk47pygp6426vjj2sj5e9n54x96ge967rrpgq27jc2v",
	}

	for _, address := range validAddresses {
		if validationResult := validators.DefaultValidator(address); validationResult != true {
			t.Error("Expected true, got", validationResult, "for address", address)
		}
	}

	invalidAddresses := []string{
		"31hr5x7HpgUTNJsdukGEUmjNNTiyVr9aT",
		"31hr5x7HpgUTNJsdukGEUmjNNTiyVr9aT11",
		"21hr5x7HpgUTNJsdukGEUmjNNTiyVr9aT11",
		"31hr5x7HpgUTNJsdukGEUmjNNTiyVr9aT2",
		"bc1qnpgqxy7nq7zt6snx0kn76lv3z6xz5dtceupg4",
		"bc1qnpgqxy7nq7zt6snx0kn76lv3z6xz5dtceupg401",
		"bc1qnpgqxy7nq7zt6snx0kn76lv3z6xz5dtceupg41",
		"bc1q7w9p2unf3hngtzmq3sq47cjdd0xd9sw7us5h6p1",
		"bc1psx9ukcl7ft7aadjsma6k6xld8yhgkhjv72qs246hstc4ayzhxwwq3ju86",
		"bc1plyq2dmxqtmwnvmfpk47pygp6426vjj2sj5e9n54x96ge967rrpgq27jc2vc",
		"bc1psx9ukcl7ft7aadjsma6k6xld8yhgkhjv72qs246hstc4ayzhxwwq3ju86q",
		"bc1plyq2dmxqtmwnvmfpk47pygp6426vjj2sj5e9n54x96ge967rrpgq27jc2z",
		"cosmos16na5gpcj80tafv5gycm4gk7garj8jjsgydtkmq",
		"12345",
		"qwerty",
		"",
		"31hr5x7HpgUTNJsdukGEUmjNNTiyVr9aT",
		"bc1qnpgqxy7nq7zt6snx0kn76lv3z6xz5dtceupg401",
		"мой биткоин адрес",
		"比特币地址",
		"Àåæ´ýú.ü.£ßòÒÀì£Ê·¸®±ô¿¯åÇÊ¯¯ï",
		"-[:%/~%!,)+~;{.\\-.%.**_[(\\$.",
	}

	for _, address := range invalidAddresses {
		if validationResult := validators.DefaultValidator(address); validationResult != false {
			t.Error("Expected false, got", validationResult, "for address", address)
		}
	}
}
