package models

import (
	"errors"
	"fmt"
)

/* This package is an API for User structs stored in our application */
type User struct {
	ID      int
	Name    string
	Surname string
}

// In Go we can use variable blocks, act as package-level State.
var (
	users []*User // a slice holding pointers to Users

	nextID = 1 // at var block, := initialization is not required
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {

	if u.ID != 0 {
		return User{}, errors.New("new user ID should be zero or undefined")
	}

	u.ID = nextID
	nextID++

	users = append(users, &u) // append new User ptr

	return u, nil
}

func GetUserByID(id int) (User, error) {
	for _, userPointer := range users {
		// In GO, no need to use dereference to access a pointer's struct field
		if userPointer.ID == id {
			return *userPointer, nil
		}
	}
	// Not Found

	return User{}, fmt.Errorf("User with ID '%v' not found", id) // Error f allows us to create errors as formatted strings
}

func UpdateUser(u User) (User, error) {
	for i, userPointer := range users {
		if userPointer.ID == u.ID {

			users[i] = &u
			return u, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", u.ID)
}

func RemoveUserByID(id int) error {

	for i, u := range users {
		if u.ID == id {
			// Append receives 1 slice and then as many params as necessary
			// Use ... to unpack an array as its individual elements
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("User with ID '%v' not found", id) // Error f allows us to create errors as formatted strings
}
