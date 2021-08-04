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
