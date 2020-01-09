package core

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

var (
	errInvalidHash         = errors.New("The encoded hash is not in the correct format")
	errIncompatibleVersion = errors.New("Incompatible version of argon2")
)

type argonParams struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

// HashString Hashes string with argon2 algorithm
func HashString(toHash string) (encodedHash string, err error) {
	params := &argonParams{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		saltLength:  6,
		keyLength:   32,
	}

	// Generate a cryptographically secure random salt.
	salt, err := generateRandomBytes(params.saltLength)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(toHash), salt, params.iterations, params.memory, params.parallelism, params.keyLength)

	// Base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// Return a string using the standard encoded hash representation.
	encodedHash = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, params.memory, params.iterations, params.parallelism, b64Salt, b64Hash)

	return encodedHash, nil
}

// CompareString Compares given string with the given argon2 hash
func CompareString(nonHashed, encodedHash string) (match bool, err error) {
	// Extract the parameters, salt and derived key from the encoded password
	// hash.
	p, salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		return false, err
	}

	// Derive the key from the other string using the same parameters.
	otherHash := argon2.IDKey([]byte(nonHashed), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	// Check that the contents of the hashed strings are identical. Note
	// that we are using the subtle.ConstantTimeCompare() function for this
	// to help prevent timing attacks.
	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}
	return false, nil
}

// Generate random bytes for salt
func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Decodes argon2 hash
func decodeHash(encodedHash string) (p *argonParams, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, errInvalidHash
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, errIncompatibleVersion
	}

	p = &argonParams{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.memory, &p.iterations, &p.parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.saltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.keyLength = uint32(len(hash))

	return p, salt, hash, nil
}
