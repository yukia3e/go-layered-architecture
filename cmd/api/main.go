package main

import (
	"log"

	"github.com/yukia3e/go-layered-architecture/internal/api"
)

func main() {
	err := api.Run()
	log.Fatal(err)
}
