package main

import (
	"fmt"

	"github.com/TankEngine-ish/Crypt_Me/decrypt"
	"github.com/TankEngine-ish/Crypt_Me/encrypt"
)

func main() {
	encryptedStr := encrypt.Plumbus("hello")
	fmt.Println(encryptedStr)
	fmt.Println(decrypt.Plumbus(encryptedStr))
}
