package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strings"
)

func main() {
	// password := "pass@word1"
	password := "123456"
	fmt.Println("Hashing:", string(password))

	//NODEJS: password-hash generated hased password
	// nodejsPassword := "sha1$7beac50a$1$83d522fea44799b322c639c22b2f0afeff71cf9a"
	p2 := "sha1$0472254a$1$001885ad69468553ab18e6049a0a16ceafcf4a27"

	matched := verifyNodePassword(password, p2)
	fmt.Printf("Password compared: %v\n", matched)
}

func verifyNodePassword(rawPassword, nodeHashedPassword string) bool {
	if len(rawPassword) == 0 || len(nodeHashedPassword) == 0 {
		return false
	}

	parts := strings.Split(nodeHashedPassword, "$")
	if len(parts) != 4 {
		return false
	}
	salt := parts[1]
	//use salt as key to generate hash
	mac := hmac.New(sha1.New, []byte(salt))
	mac.Write([]byte(rawPassword))
	hashedPassword := mac.Sum(nil)
	hashedPasswordStr := hex.EncodeToString(hashedPassword)
	//compare hashed string is equal
	return hashedPasswordStr == parts[3]
}