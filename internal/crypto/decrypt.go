package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Decrypt(data string) (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", err
	}
	key := []byte(os.Getenv("AES_KEY"))

	cryptedText, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesgcm.NonceSize()
	if len(cryptedText) < nonceSize {
		return "", fmt.Errorf("texto criptografado inválido: o tamanho do texto é menor que o nonce")
	}

	nonce, ciphertext := cryptedText[:nonceSize], cryptedText[nonceSize:]

	plainText, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}
