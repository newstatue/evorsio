package vault

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/newstatue/evorsio/internal/resource"
)

func TestVault_Encrypt(t *testing.T) {
	vault, _ := New("wdwad")
	encrypt, _ := vault.Encrypt(&Entry{
		Resource: resource.New("", resource.KindVault),
		Payload: Payload{
			Title:    "title",
			Username: "username",
			Password: "password",
			URL:      "url",
			Notes:    "note",
		},
	})
	fmt.Println(encrypt)
}

func TestVault_Decrypt(t *testing.T) {
	vault, _ := New("wdwad")
	encrypt, _ := vault.Encrypt(&Entry{
		Resource: resource.New("", resource.KindVault),
		Payload: Payload{
			Title:    "title",
			Username: "username",
			Password: "password",
			URL:      "url",
			Notes:    "note",
		},
	})
	decrypt, _ := vault.Decrypt(bytes.NewReader(encrypt))
	fmt.Printf("%+v\n", decrypt)
}
