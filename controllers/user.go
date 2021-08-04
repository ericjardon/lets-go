package controllers

import (
	"encoding/json"
	"lets-go/models"
	"net/http"
	"regexp"
	"strconv"
)

/* We can add behavior (methods) to struct datatypes */
type userController struct {
	userIDPatt *regexp.Regexp
}

// Constructor
func newUserController() *userController {
	// Initializes a userController and returns the address pointer
	return &userController{
		userIDPatt: regexp.MustCompile(`^/users/(\d+)/?`),
	}

}

// Receives and processes a Network Http Request
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/users" {

		// Decide what function to send the request to
		switch r.Method {
		case http.MethodGet:
			// Get list of users
			uc.getAll(w, r)
		case http.MethodPost:
			// New User
			uc.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented) // 501
		}

	} else {
		// Check if url path matches /users/:idNumber
		matches := uc.userIDPatt.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}

		parsed_id, err := strconv.Atoi(matches[1])

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}

		// All ok, decide method handler
		switch r.Method {
		case http.MethodGet:
			// Get a user by id
			uc.get(parsed_id, w)
		case http.MethodPut:
			// Update a user by id
			uc.put(parsed_id, w, r)
		case http.MethodDelete:
			// Delete a user by id
			uc.delete(parsed_id, w)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}

	}
}

func (uc *userController) getAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(models.GetUsers(), w)
}

func (uc *userController) get(id int, w http.ResponseWriter) {
	u, err := models.GetUserByID(id)

	if err != nil {
		w.Write([]byte("Requested User does not exist"))
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	encodeResponseAsJSON(u, w)

}

func (uc *userController) post(w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}

	u, err = models.AddUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := []byte(err.Error())
		w.Write(errorMessage)
		return
	}
	encodeResponseAsJSON(u, w)
}

func (uc *userController) put(id int, w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		w.Write([]byte("Could not parse User Object"))
		return
	}

	if id != u.ID {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID of submitted user should match ID in Request"))
		return
	}

	u, err = models.UpdateUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := []byte(err.Error())
		w.Write(errorMessage)
		return
	}

	encodeResponseAsJSON(u, w)
}

func (uc *userController) delete(id int, w http.ResponseWriter) {
	err := models.RemoveUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := []byte(err.Error())
		w.Write(errorMessage)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (uc *userController) parseRequest(r *http.Request) (models.User, error) {
	decoder := json.NewDecoder(r.Body)
	var u models.User

	decoding_err := decoder.Decode(&u) // read the request body into a User variable

	if decoding_err != nil {
		return models.User{}, decoding_err
	}

	return u, nil
}
