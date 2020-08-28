package hackertime

import (
	"crypto/rand"
	"fmt"
	"log"
	"testing"
)

var (
	Message = "Be sure to drink you ovaltine"
)

func TestEncrypt(t *testing.T) {
	var (
		key   [32]byte
		nonce [24]byte
	)
	n, err := rand.Read(key[:])
	if err != nil {
		log.Fatalf("Could not read random bytes: %s", err)
	}
	if n != 32 {
		log.Fatalf("Generated %d bytes, the wrong number of bytes", n)
	}
	_, err = rand.Read(nonce[:])
	if err != nil {
		log.Fatalf("Could not read random bytes: %s", err)
	}
	ciphertext := Encrypt(Message, nonce, key)
	plaintext := Decrypt(ciphertext, nonce, key)
	fmt.Printf("Message: %s\n", Message)
	fmt.Printf("Ciphertext: %s\n", ciphertext)
	fmt.Printf("Plaintext: %s\n", plaintext)
	if Message != string(plaintext) {
		t.Fatalf("Unexpected message from decrypt: %s", plaintext)
	}

}
