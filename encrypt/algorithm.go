// encrypt package consists of all the logic to encrypt an input string.
package encrypt

func Plumbus(str string) string {
	encryptedStr := ""
	for _, c := range str {
		asciiCode := int(c)
		character := string(asciiCode + 3)
		encryptedStr += character
	}
	return encryptedStr
}
