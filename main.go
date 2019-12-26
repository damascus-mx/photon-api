package main

import (
	"log"
	"net/http"

	app "github.com/damascus-mx/photon-api/src/bin"
)

func main() {
	err := http.ListenAndServe(":3000", app.InitApplication())
	// http.ListenAndServeTLS(":443", "./security/localhost.crt", "./security/localhost.key", app.InitApplication())
	if err != nil {
		log.Fatal(err.Error())
	}
}
