package vault

import (
	"net/url"

	"github.com/newstatue/evorsio/internal/fsgen"
	"github.com/newstatue/evorsio/internal/resource"
)

type Entry struct {
	*resource.Resource
	Payload Payload
}

type Payload struct {
	Title    string `json:"title"`
	Username string `json:"username"`
	Password string `json:"password"`
	URL      string `json:"url"`
	Notes    string `json:"notes"`
}

func NewEntry(payload *Payload) (*Entry, error) {
	u, err := url.Parse(payload.URL)
	if err != nil {
		return nil, err
	}

	return &Entry{
		Resource: resource.New(u.Host, resource.KindVault),
		Payload:  *payload,
	}, nil
}

func NewEntryFromFS(dir string, entry *fsgen.Entry) *Entry {
	return &Entry{
		Resource: resource.NewFromFS(dir, entry),
		Payload: Payload{
			Title:    "",
			Username: "",
			Password: "",
			URL:      "",
			Notes:    "",
		},
	}
}
