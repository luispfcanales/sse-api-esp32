package main

import (
	"log"

	"github.com/luispfcanales/sse-api/api"
)

func main() {

	if err := api.Run(); err != nil {
		log.Fatal(err)
	}
}
