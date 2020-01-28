package main

import (
	"log"
	"net/http"
	"os"

	"github.com/apex/gateway"
	v1 "github.com/feud72/Wetube-go/v1"
)

func main() {

	addr := ":3000"
	mode := os.Getenv("GIN_MODE")

	app := v1.Engine()
	if mode == "release" {
		log.Fatal(gateway.ListenAndServe(addr, app))
	} else {
		log.Fatal(http.ListenAndServe(addr, app))
	}
}
