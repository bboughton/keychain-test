package main

import (
	"fmt"

	"github.com/lunixbochs/go-keychain"
)

func main() {
	service := "foo.bar"

	accessKey, _ := keychain.Find(service, "AWS_ACCESS_KEY")
	secretKey, _ := keychain.Find(service, "AWS_SECRET_KEY")

	fmt.Printf("Access Key: %s\n", accessKey)
	fmt.Printf("Secret Key: %s\n", secretKey)
}
