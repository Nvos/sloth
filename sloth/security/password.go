package security

import (
	"github.com/alexedwards/argon2id"
)

type Hasher interface {
	Hash(string) (string, error)
	Compare(password, hash string) (bool, error)
}

type HasherConfig struct {
	time      uint32
	memory    uint32
	threads   uint8
	keyLength uint32
}

type hasher struct {
	config *argon2id.Params
}

func NewHasher() Hasher {
	config := &argon2id.Params{
		Memory:      64 * 1024,
		KeyLength:   32,
		SaltLength:  16,
		Parallelism: 4,
		Iterations:  1,
	}

	return &hasher{
		config: config,
	}
}

func (h *hasher) Hash(password string) (string, error) {
	return argon2id.CreateHash(password, h.config)
}

func (h *hasher) Compare(password, hash string) (bool, error) {
	return argon2id.ComparePasswordAndHash(password, hash)
}
