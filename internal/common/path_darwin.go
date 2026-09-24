package common

import (
	"path/filepath"
)

func exePath(filename string, exeDir string) string {
	return filepath.Clean(filepath.Join(exeDir, "..", "Resources", filename))
}
