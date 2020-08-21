package hackertime

import (
	"crypto/rand"
	"fmt"
	"log"

	"golang.org/x/crypto/nacl/secretbox"
)

func Encrypt(s string) {
	var (
		out   []byte
		nonce [24]byte
		key   [32]byte
	)

	n, err := rand.Read(nonce[:])
	if err != nil {
		log.Fatalf("Could not read random bytes: %s", err)
	}
	n, err = rand.Read(key[:])
	if err != nil {
		log.Fatalf("Could not read random bytes: %s", err)
	}
	fmt.Printf("you read %d bytes\n", n)
	fmt.Printf("your bytes %s\n", nonce)
	fmt.Printf("your bytes %s\n", key)
	buf := secretbox.Seal(out, []byte(s), &nonce, &key)

	fmt.Println(buf)
}

func Decrypt() {

}
