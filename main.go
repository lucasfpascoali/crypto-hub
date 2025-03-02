package main

import (
	"fmt"
	"github.com/lucasfpascoali/crypto-hub/classical"
)

func main() {
	shift := classical.NewShiftCipher(3, true)
	cipherText, err := shift.Encrypt([]byte("Hello World"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(cipherText))

	plainText, err := shift.Decrypt(cipherText)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(plainText))
}
