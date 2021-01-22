package main

import (
	"net/http"
)

func loadAPIConfiguration() {

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

	http.Handle("/", router)
	http.ListenAndServe(":8000", nil)

}
