package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/newstatue/evorsio/internal/common"
	"github.com/newstatue/evorsio/internal/constant"
)

func main() {
	cfg, err := common.NewConfig()
	if err != nil {
		panic(err)
	}

	out, err := exec.Command("mise", "where", "github:seaweedfs/seaweedfs").Output()
	if err != nil {
		panic(err)
	}

	dir := strings.TrimSpace(string(out))

	srcName := "weed"

	if runtime.GOOS == "windows" {
		srcName = "weed.exe"
	}

	src := filepath.Join(dir, srcName)

	dst := filepath.Join(cfg.FS.Path)

	if err := os.MkdirAll(constant.BuildDir, 0755); err != nil {
		panic(err)
	}

	data, err := os.ReadFile(src)
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(dst, data, 0755); err != nil {
		panic(err)
	}

	fmt.Printf("SeaweedFS: %s -> %s\n", src, dst)
}
