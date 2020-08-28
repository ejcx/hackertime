package hackertime

import (
	"log"

	"golang.org/x/crypto/nacl/secretbox"
)

func Encrypt(s string, nonce [24]byte, key [32]byte) []byte {
	var (
		out []byte
	)
	ciphertext := secretbox.Seal(out, []byte(s), &nonce, &key)
	return ciphertext
}

func Decrypt(ciphertext []byte, nonce [24]byte, key [32]byte) []byte {
	var (
		out []byte
	)
	plaintext, ok := secretbox.Open(out, ciphertext, &nonce, &key)
	if !ok {
		log.Fatal("Open operation not okay")
	}
	return plaintext
}
