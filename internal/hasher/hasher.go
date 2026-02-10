package hasher

import (
	"errors"
	"log/slog"

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
	cdg  *config.Config
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
		slog.Error("failed to encode id", "error", err.Error())
		return "", err
	}

	return code, nil
}

func (h *HashID) Decode(code string) (int, error) {
	if code == "" {
		slog.Error("malformed short code")
		return 0, ErrInvalidCode
	}

	id, err := h.hash.DecodeWithError(code)

	if err != nil {
		slog.Error("failed to decode hash", "error", err.Error())
		return 0, ErrInvalidCode
	}

	if len(id) == 0 {
		slog.Warn("invalid code")
		return 0, ErrInvalidCode
	}

	return id[0], nil
}
