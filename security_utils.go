package main

import (
	"fmt"
	"golang.org/x/crypto/argon2"
	"math/rand"
	"strconv"
)

// KeyDerivationParams holds parameters for key derivation
type KeyDerivationParams struct {
	Salt        []byte
	Iterations  uint32
	Memory      uint32
	Parallelism uint8
	KeyLength   uint32
}

// DeriveKey generates a derived key using Argon2
func DeriveKey(password string, params KeyDerivationParams) ([]byte, error) {
	hashedKey := argon2.IDKey([]byte(password), params.Salt, params.Iterations, params.Memory, params.Parallelism, params.KeyLength)
	return hashedKey, nil
}

func getRandom() uint64 {
	r := rand.New(rand.NewSource(112233))
	return r.Uint64()
}

func getRandomDerivedKey() []byte {
	// Example usage
	password := "yourSecretPassword"
	// Set your key derivation parameters
	random := getRandom()
	params := KeyDerivationParams{
		Salt:        []byte(strconv.FormatUint(random, 10)), // Generate a random salt and store it securely
		Iterations:  3,
		Memory:      1 << 14, // 16MB of memory
		Parallelism: 2,
		KeyLength:   32,
	}

	// Derive the key
	derivedKey, err := DeriveKey(password, params)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	fmt.Printf("Derived Key (hex): %x\n", derivedKey)
	return derivedKey
}
