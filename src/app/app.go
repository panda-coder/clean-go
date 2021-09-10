package app

import (
	"log"
	"net/http"
)

// função principal
func init() {
	log.Println("Server Starting")
	log.Fatal(http.ListenAndServe(":3000", setup_router()))
}
