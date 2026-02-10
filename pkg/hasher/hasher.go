package hasher

import (
	"errors"
	"fmt"

	"github.com/m-bromo/atom-ly/config"
	"github.com/speps/go-hashids/v2"
)

const base62Alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const MinLength = 7

var (
	ErrInvalidCode = errors.New("the inserted code is invalid")
)

type Hasher interface {
	Encode(id int) (string, error)
	Decode(code string) (int, error)
}

type HashID struct {
	hash *hashids.HashID
}

func NewHashID(cfg *config.Config) *HashID {
	data := hashids.NewData()
	data.Salt = cfg.Env.Salt
	data.MinLength = MinLength
	data.Alphabet = base62Alphabet
	h, _ := hashids.NewWithData(data)

	return &HashID{
		hash: h,
	}
}

func (h *HashID) Encode(id int) (string, error) {
	code, err := h.hash.Encode([]int{id})
	if err != nil {
		return "", fmt.Errorf("failed to encode id: %w", err)
	}

	return code, nil
}

func (h *HashID) Decode(code string) (int, error) {
	if code == "" {
		return 0, ErrInvalidCode
	}

	id, err := h.hash.DecodeWithError(code)

	if err != nil {
		return 0, ErrInvalidCode
	}

	if len(id) == 0 {
		return 0, ErrInvalidCode
	}

	return id[0], nil
}
