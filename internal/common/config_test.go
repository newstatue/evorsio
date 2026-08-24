package common

import "testing"

func TestConfig_Init(t *testing.T) {
	config := NewConfig()
	t.Logf("%v\n", config)
}
