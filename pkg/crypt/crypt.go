package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"io"
	"os"

	files "github.com/cazier/resume/pkg/files"
	shared "github.com/cazier/resume/pkg/shared"
)

func getGCM(key string) (output cipher.AEAD) {
	h := sha256.New()
	_, err := h.Write([]byte(key))

	shared.HandleError(err)

	ci, err := aes.NewCipher(h.Sum(nil))

	shared.HandleError(err)

	output, err = cipher.NewGCM(ci)

	shared.HandleError(err)

	return output
}

func Decrypt(encrypted []byte, key string) (plaintext []byte) {
	gcm := getGCM(key)

	size := gcm.NonceSize()

	nonce, encrypted_text := encrypted[:size], encrypted[size:]

	plaintext, err := gcm.Open(nil, nonce, encrypted_text, nil)

	shared.HandleError(err)

	return plaintext
}

func Encrypt(plaintext []byte, key string) (encrypted []byte) {
	gcm := getGCM(key)

	n := make([]byte, gcm.NonceSize())

	_, err := io.ReadFull(rand.Reader, n)

	shared.HandleError(err)

	encrypted = gcm.Seal(n, n, plaintext, nil)

	return encrypted
}

func Conversion(key, in, out string, overwrite bool, conv func(in []byte, key string) []byte) {
	if !files.Exists(in) {
		shared.Exit(1, "Could not find the input file: %s", in)
	} else if files.Exists(out) && !overwrite {
		shared.Exit(1, "An output file exists and overwrite was not enabled: %s", out)
	}

	data, err := os.ReadFile(in)
	shared.HandleError(err)

	output := conv(data, key)

	err = os.WriteFile(out, []byte(output[:]), 0644)
	shared.HandleError(err)
	shared.Exit(0, "Success! The file was saved to %s", out)

}
