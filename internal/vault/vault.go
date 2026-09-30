package vault

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"

	"filippo.io/age"
)

const (
	vaultVersion = "evorsio_vault_v1"
)

var (
	ErrInvalidMasterPass = errors.New("主密码错误")
)

type Vault struct {
	recipient age.Recipient
	identity  age.Identity
}

func New() *Vault {
	return &Vault{}
}

func (v *Vault) Init(masterPass string) ([]byte, error) {
	recipient, err := age.NewScryptRecipient(masterPass)
	if err != nil {
		return nil, err
	}
	identity, err := age.NewScryptIdentity(masterPass)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	w, err := age.Encrypt(&buf, recipient)
	if err != nil {
		return nil, err
	}
	if _, err := w.Write([]byte(vaultVersion)); err != nil {
		_ = w.Close()
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}

	v.recipient = recipient
	v.identity = identity
	return buf.Bytes(), nil
}
func (v *Vault) IsLocked() bool {
	return v.recipient == nil || v.identity == nil
}

func (v *Vault) Unlock(masterPass string, data []byte) error {
	identity, err := age.NewScryptIdentity(masterPass)
	if err != nil {
		return err
	}
	r, err := age.Decrypt(bytes.NewReader(data), identity)
	if err != nil {
		return ErrInvalidMasterPass
	}
	plain, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	if (string(plain)) != vaultVersion {
		return ErrInvalidMasterPass
	}
	recipient, err := age.NewScryptRecipient(masterPass)
	if err != nil {
		return err
	}

	v.recipient = recipient
	v.identity = identity

	return nil
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
