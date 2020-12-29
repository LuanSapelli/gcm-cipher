package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"strings"
)

type GCM struct {
	cipher cipher.AEAD
	iv     []byte
	aead   []byte
}

func NewGCM(cryptKey, cryptIv string) (*GCM, error) {
	var err error
	key, _ := hex.DecodeString(cryptKey)
	iv, _ := hex.DecodeString(cryptIv)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCMWithNonceSize(block, 16)
	if err != nil {
		return nil, err
	}

	return &GCM{
		cipher: gcm,
		iv:     iv,
	}, nil
}

func (c *GCM) Encrypt(text string) (string, error) {
	cipherText := c.cipher.Seal(nil, c.iv, []byte(text), nil)
	return c.getEncryptedCipherTextWithTag(cipherText), nil
}

func (c *GCM) Decrypt(text string) (string, error) {
	cipherTextEncoded, _, err := c.getCipherTextWithTag(text)
	plaintext, err := c.cipher.Open(nil, c.iv, cipherTextEncoded, c.aead)
	return string(plaintext), err
}

func (c *GCM) getEncryptedCipherTextWithTag(ciphertext []byte) string {
	tagPosition := len(ciphertext) - c.cipher.Overhead()

	tag := ciphertext[tagPosition:]
	ciphertext = ciphertext[:tagPosition]

	tagEncoded := base64.StdEncoding.EncodeToString(tag)
	cipherTextEncoded := base64.StdEncoding.EncodeToString(ciphertext)

	return tagEncoded + "$" + cipherTextEncoded
}

func (c *GCM) getCipherTextWithTag(cipherText string) ([]byte, int, error) {
	if !strings.Contains(cipherText, "$") {
		return nil, -1, errors.New("invalid ciphertext")
	}
	split := strings.Split(cipherText, "$")
	tagEncoded, _ := base64.StdEncoding.DecodeString(split[0])
	cipherTextEncoded, _ := base64.StdEncoding.DecodeString(split[1])

	return append(cipherTextEncoded, tagEncoded...), len(tagEncoded), nil
}
