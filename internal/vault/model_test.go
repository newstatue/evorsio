package vault

import (
	"fmt"
	"testing"
)

func TestNewEntry(t *testing.T) {
	entry, _ := NewEntry(&Payload{
		Title:    "title",
		Username: "user",
		Password: "pass",
		URL:      "https://github.com/login",
		Notes:    "note",
	})
	fmt.Printf("Resource: %+v\n", entry.Resource)
	fmt.Printf("Payload: %+v\n", entry.Payload)
}
