package main

import (
	"net/http"

	app "github.com/damascus-mx/photon-api/src/bin"
)

func main() {
	http.ListenAndServe(":3000", app.InitApplication())
}
