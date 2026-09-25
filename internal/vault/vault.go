package vault

import (
	"bytes"
	"encoding/json"
	"io"

	"filippo.io/age"
)

type Vault struct {
	recipient age.Recipient
	identity  age.Identity
}

func New(masterPass string) (*Vault, error) {
	recipient, err := age.NewScryptRecipient(masterPass)
	if err != nil {
		return nil, err
	}
	identity, err := age.NewScryptIdentity(masterPass)
	if err != nil {
		return nil, err
	}
	return &Vault{
		recipient: recipient,
		identity:  identity,
	}, nil
}

func (v *Vault) Encrypt(payload *Payload) ([]byte, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	w, err := age.Encrypt(&buf, v.recipient)
	if err != nil {
		return nil, err
	}
	defer func(w io.WriteCloser) {
		_ = w.Close()
	}(w)

	if _, err := w.Write(data); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (v *Vault) Decrypt(r io.Reader) (*Payload, error) {
	r, err := age.Decrypt(r, v.identity)
	if err != nil {
		return nil, err
	}

	var payload Payload
	if err := json.NewDecoder(r).Decode(&payload); err != nil {
		return nil, err
	}

	return &payload, nil
}
