package util

//changes a byte array based upon a given key
fun encodeDecode(input []byte, key string) []byte {
	var bArr = make([]byte, len(input))
	for i :=0; i < len(input); i++ {
		bArr[i] += input[i] ^ key[i%len(key)]
	}

	return bArr
}

//returns the encoded message
func XorEncode(decode []byte, key string) []byte {
	return encodeDecode(decode, key)
}

//returns the decoded message
func XorDecode(encode []byte, key string) []byte {
	return encodeDecode(encode, key)
}