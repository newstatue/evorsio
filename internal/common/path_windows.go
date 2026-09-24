package common

import (
	"path/filepath"
)

func exePath(filename string, exeDir string) string {
	return filepath.Join(exeDir, filename+".exe")
}
