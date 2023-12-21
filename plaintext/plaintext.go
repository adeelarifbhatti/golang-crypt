package plaintext

func Decryption(str string) string {
	plaintext := ""
	for _, c := range str {
		asciiCode := int(c)
		character := string(asciiCode - 3)
		plaintext += character
	}
	return plaintext
}