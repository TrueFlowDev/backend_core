package adapter

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/TrueFlowDev/Backend/internal/module/auth/domain/port"
	"golang.org/x/crypto/argon2"
)

type argon2Params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

var defaultParams = argon2Params{
	memory:      64 * 1024,
	iterations:  3,
	parallelism: 2,
	saltLength:  16,
	keyLength:   32,
}

type PasswordHasher struct {
	params argon2Params
}

func NewPasswordHasher() *PasswordHasher { return &PasswordHasher{params: defaultParams} }

func (p *PasswordHasher) Hash(password string) (string, error) {
	salt := make([]byte, p.params.saltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("%w: %w", port.ErrInvalidHash, err)
	}

	hash := argon2.IDKey(
		[]byte(password),
		salt,
		p.params.iterations,
		p.params.memory,
		p.params.parallelism,
		p.params.keyLength,
	)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encoded := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		p.params.memory,
		p.params.iterations,
		p.params.parallelism,
		b64Salt,
		b64Hash,
	)

	return encoded, nil
}

func (p *PasswordHasher) Validate(password string, hashedPassword string) (bool, error) {
	params, salt, hash, err := decodeHash(hashedPassword)
	if err != nil {
		return false, fmt.Errorf("%w: %w", port.ErrInvalidHash, err)
	}

	computedHash := argon2.IDKey(
		[]byte(password), salt,
		params.iterations, params.memory, params.parallelism, params.keyLength,
	)

	if subtle.ConstantTimeCompare(hash, computedHash) != 1 {
		return false, port.ErrPasswordMismatch
	}

	return true, nil
}

func decodeHash(encoded string) (argon2Params, []byte, []byte, error) {
	parts := strings.Split(encoded, "$")
	if len(parts) != 6 {
		return argon2Params{}, nil, nil,
			fmt.Errorf("expected 6 parts, got %d", len(parts))
	}

	var version int
	if _, err := fmt.Sscanf(parts[2], "v=%d", &version); err != nil {
		return argon2Params{}, nil, nil,
			fmt.Errorf("failed to parse version: %w", err)
	}
	if version != argon2.Version {
		return argon2Params{}, nil, nil,
			fmt.Errorf("unsupported argon2 version: got %d, want %d", version, argon2.Version)
	}

	var params argon2Params
	if _, err := fmt.Sscanf(
		parts[3], "m=%d,t=%d,p=%d",
		&params.memory, &params.iterations, &params.parallelism,
	); err != nil {
		return argon2Params{}, nil, nil,
			fmt.Errorf("failed to parse params: %w", err)
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return argon2Params{}, nil, nil,
			fmt.Errorf("failed to decode salt: %w", err)
	}
	params.saltLength = uint32(len(salt))

	hash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return argon2Params{}, nil, nil,
			fmt.Errorf("failed to decode hash: %w", err)
	}
	params.keyLength = uint32(len(hash))

	return params, salt, hash, nil
}
