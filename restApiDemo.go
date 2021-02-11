package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

/*
Simple install the package to your $GOPATH with the go tool from shell:

$ go get -u github.com/gorilla/mux
*/

var router = mux.NewRouter()

func loadAPIConfiguration() {
	// Rest and HTML
	router.HandleFunc("/", indexPageHandler)
	router.HandleFunc("/internal", internalPageHandler)
	router.HandleFunc("/internalUser", internalUserPageHandler)

	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/loginError", loginErrorHandler).Methods("GET")
	router.HandleFunc("/logout", logoutHandler).Methods("POST")

	router.HandleFunc("/signup", signupHandler).Methods("POST")
	router.HandleFunc("/signupSuccess", signupSuccessHandler).Methods("GET")

	router.HandleFunc("/delete/{id}", deleteUserHandler)

	router.HandleFunc("/update/{id}", updateUserFormHandler)
	router.HandleFunc("/updateUser/{id}", updateUserHandler)

	router.HandleFunc("/searchByUserName", searchUserHandler)

	// Rest API only and test case
	router.HandleFunc("/users", getAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	// Main api configuration
	http.Handle("/", router)
	http.ListenAndServe(":8000", nil)

}
