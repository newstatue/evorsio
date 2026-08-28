package main

import (
	"log"

	"github.com/newstatue/evorsio/internal/common"
	_ "modernc.org/sqlite"
)

func main() {
	cfg, err := common.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(cfg)
}
