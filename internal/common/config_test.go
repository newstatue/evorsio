package common

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	os.Setenv("NAME", "world")
	os.Setenv("VERSION", "0.0.0")
	os.Setenv("COUNT", "zxc")
	config, err := NewConfig()
	assert.Error(t, err, config)
}
