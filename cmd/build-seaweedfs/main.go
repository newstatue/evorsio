package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/newstatue/evorsio/internal/common"
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

	data, err := os.ReadFile(src)
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(dst, data, 0755); err != nil {
		panic(err)
	}

	fmt.Printf("SeaweedFS: %s -> %s\n", src, dst)
}
