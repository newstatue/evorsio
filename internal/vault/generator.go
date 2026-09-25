package vault

import (
	"crypto/rand"
	"math/big"
)

const (
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	numbers   = "0123456789"
	symbols   = "!@#$%^&*()-_=+"
)

type Generator struct{}

func NewGenerator() *Generator {
	return &Generator{}
}

type Options struct {
	Length    int
	Uppercase bool
	Lowercase bool
	Numbers   bool
	Symbols   bool
}

func DefaultOptions() Options {
	return Options{
		Length:    12,
		Uppercase: true,
		Lowercase: true,
		Numbers:   true,
		Symbols:   false,
	}
}

func (g *Generator) Generate(opt Options) (string, error) {
	var chars string
	if opt.Uppercase {
		chars += uppercase
	}
	if opt.Lowercase {
		chars += lowercase
	}
	if opt.Numbers {
		chars += numbers
	}
	if opt.Symbols {
		chars += symbols
	}
	if chars == "" {
		return "", nil
	}

	result := make([]byte, opt.Length)
	for i := range result {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", err
		}

		result[i] = chars[n.Int64()]
	}

	return string(result), nil
}
