package main

import (
	"net/http"

	"github.com/damascus-mx/photon-api/authentication/bin"
)

const (
	port string = ":8080"
)

func main() {
	// Get Bootstrapper
	bootstrap := new(bin.Bootstrapper)
	bootstrap.StartServices()
	// Init HTTP Server
	err := http.ListenAndServe(port, bootstrap.StartHTTP())
	if err != nil {
		panic("Cannot run Photon's Authentication Service HTTP")
	}
}
