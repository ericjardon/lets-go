package main

import (
	"net/http"

	"lets-go/controllers"
)

func main() {
	// Register a route controller
	controllers.RegisterControllers()

	// In Go we use the front-controller / back-controller pattern
	http.ListenAndServe(":3011", nil) // second argument is the so called ServeMux, handles high-level routing
}

// run with `go run main.go` || `go run lets-go` || `go run .` (from root)

// Whole application is run with `go build .`
// and then `./executable-name`
