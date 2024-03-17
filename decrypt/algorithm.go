// decrypt package consists of all the logic to decrypt an encrypted string.
package decrypt

// decrypts by reducing the ascii code of each character by 3
func Plumbus(str string) string {
	decryptedStr := ""
	for _, c := range str {
		asciiCode := int(c)
		character := string(asciiCode - 3)
		decryptedStr += character
	}
	return decryptedStr
}
