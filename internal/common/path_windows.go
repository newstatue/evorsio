package common

import (
	"fmt"
	"os"
	"path/filepath"
)

func exePath(filename string, exeDir string) string {
	return filepath.Join(exeDir, "bin", filename+".exe")
}

func mountDir() string {
	for c := 'Z'; c >= 'D'; c-- {
		drive := fmt.Sprintf("%c:", c)

		if _, err := os.Stat(drive + `\`); err != nil {
			return drive
		}
	}
	return ""
}
