package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/newstatue/evorsio/internal/common"
)

func main() {
	cfg := common.NewConfig()
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{name}", func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		_, _ = fmt.Fprintf(w, "Hello, %s", name)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.HTTP.Port), mux))
}
