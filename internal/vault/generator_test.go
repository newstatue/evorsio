package vault

import (
	"fmt"
	"testing"
)

func TestGenerator_Generate(t *testing.T) {
	generator := NewGenerator()
	opt := Options{
		Length:    10,
		Uppercase: true,
		Lowercase: true,
		Numbers:   true,
		Symbols:   true,
	}
	for range 100 {
		password, err := generator.Generate(opt)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println(password)
	}
}
