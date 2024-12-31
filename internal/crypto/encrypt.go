package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"os"

	"github.com/joho/godotenv"
)

func Encrypt(data string) (string, error) {
	plainText := []byte(data)

	// Recupera a key do arquivo .env
	if err := godotenv.Load(); err != nil {
		return "", err
	}
	key := []byte(os.Getenv("AES_KEY"))

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Gera um nonce unico para a critograficar
	// Para poder tornar mais segura
	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cryptedText := aesgcm.Seal(nil, nonce, plainText, nil)

	ciphertextWithNonce := append(nonce, cryptedText...)

	return base64.StdEncoding.EncodeToString(ciphertextWithNonce), nil
}
