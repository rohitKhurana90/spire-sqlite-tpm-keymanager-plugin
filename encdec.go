package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
)

// EncryptData encrypts a string using AES-GCM with the provided key
func EncryptDataWithRand(plaintext []byte, key []byte) (string, error) {
	// Create a new AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Generate a random nonce
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// Create a new AES-GCM cipher
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Encrypt the plaintext
	ciphertext := aesGCM.Seal(nil, nonce, plaintext, nil)

	// Combine the nonce and ciphertext for storage
	encrypted := append(nonce, ciphertext...)

	// Return the encrypted data as a hexadecimal string
	return hex.EncodeToString(encrypted), nil
}

// DecryptData decrypts a string using AES-GCM with the provided key
func DecryptData(encryptedHex string, key []byte) ([]byte, error) {
	// Decode the hexadecimal string to get the encrypted data
	encrypted, err := hex.DecodeString(encryptedHex)
	if err != nil {
		return nil, err
	}

	// Create a new AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Extract the nonce from the encrypted data
	nonceSize := 12
	nonce := encrypted[:nonceSize]

	// Create a new AES-GCM cipher
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Decrypt the ciphertext
	decryptedData, err := aesGCM.Open(nil, nonce, encrypted[nonceSize:], nil)
	if err != nil {
		return nil, err
	}

	// Return the decrypted plaintext as a string
	return decryptedData, nil
}
