package main

import (
	"log"
	"net/http"

	bin "github.com/damascus-mx/photon-api/tenant/bin"
)

const port string = ":8080"

func main() {
	// Init HTTP REST server
	err := http.ListenAndServe(port, bin.InitHTTPServer())
	if err != nil {
		log.Fatal(err.Error())
	}

}
