package main

import (
	"log"
	"net/http"

	"github.com/felixphew/pokepaste"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", pokepaste.Handler))
}
