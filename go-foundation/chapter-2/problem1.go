package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	fmt.Println("Hii crypto")

	str := "MySecretPassword"

	// Create new SHA-256 hasher
	hash := sha256.New()

	// Write the string to hasher (convert string to byte slice)
	hash.Write([]byte(str))

	// Get the final hash sum
	hashed := hash.Sum(nil) // returns []byte

	// Print the hash in hexadecimal format
	fmt.Printf("SHA-256 hash: %x\n", hashed)
}
