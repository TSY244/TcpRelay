package encryption

func Encrypt(method string, data []byte, key string) []byte {
	switch method {
	case "6N_XOR":
		return SixN_Xor(data, key)
	default:
		return nil
	}
}

func Decrypt(method string, data []byte, key string) []byte {
	switch method {
	case "6N_XOR":
		return SixN_Xor(data, key)
	default:
		return nil
	}
}
