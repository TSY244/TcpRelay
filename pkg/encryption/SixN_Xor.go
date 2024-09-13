package encryption

func checkKey(key string) bool {
	if len(key) != 6 {
		return false
	}
	// check if key is six number
	for _, v := range key {
		if v < '0' || v > '9' {
			return false
		}
	}
	return true
}

func SixN_Xor(data []byte, key string) []byte {
	if !checkKey(key) {
		return nil
	}
	keyBytes := []byte(key)
	for i := 0; i < len(data); i++ {
		data[i] ^= keyBytes[i%len(key)]
	}
	return data
}
