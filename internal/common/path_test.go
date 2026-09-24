package common

import (
	"fmt"
	"testing"
)

func TestGetExePath(t *testing.T) {
	fmt.Print(GetExePath("hello"))
}
