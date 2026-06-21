package main

import (
	"log"

	"ipfix_observability/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	log.Printf("Loaded configuration: %+v", cfg)
}
