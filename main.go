package main 
import (
	"fmt"
	"github.com/adeelarifbhatti/golang-crypt/cipher"
	"github.com/adeelarifbhatti/golang-crypt/plaintext"
	)

func main(){
	fmt.Println("Enter data for Encryption ")
	var text string
	fmt.Scanln(&text)
	fmt.Println(cipher.Encryption(text))
	fmt.Println("Enter data for Decryption ")
	fmt.Scanln(&text)
	fmt.Println(plaintext.Decryption(text))
}