package main

import (
	"log"
	"net/http"

	app "github.com/damascus-mx/photon-api/users/bin"
)

func main() {
	err := http.ListenAndServe(":8080", app.InitApplication())
	// http.ListenAndServeTLS(":443", "./security/localhost.crt", "./security/localhost.key", app.InitApplication())
	if err != nil {
		log.Fatal(err.Error())
	}
}
