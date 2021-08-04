/*
Acts as the front end router,
roughly matching route paths to the correct controller.
*/
package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers() {
	uc := newUserController() // is a pointer to userController

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

// Static method of a controller to send responses in JSON
func encodeResponseAsJSON(data interface{}, w io.Writer) {
	// Receives an object and returns in JSON representation
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
