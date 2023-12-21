package cipher

func Encryption(str string) string {
	ciphertext := ""
	for _, c := range str {
		asciiCode := int(c)
		character := string(asciiCode + 3)
		ciphertext += character
	}
	return ciphertext
}