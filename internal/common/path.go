package common

import (
	"os"
	"path/filepath"
)

func GetExeDir() string {
	exe, _ := os.Executable()
	return filepath.Dir(exe)
}

func GetExePath(filename string) string {
	return exePath(filename, GetExeDir())
}

func GetMountDir() string {
	return mountDir()
}
