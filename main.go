package main

import (
	"log"

	"github.com/luispfcanales/sse-api-esp32/api"
)

func main() {

	if err := api.Run(); err != nil {
		log.Fatal(err)
	}
}
