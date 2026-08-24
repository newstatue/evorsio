package main

import (
	"fmt"
	"os"

	"github.com/newstatue/evorsio/internal/common"
)

func main() {
	os.Setenv("TEST", "")
	config := common.NewConfig()
	fmt.Println(config)
}
