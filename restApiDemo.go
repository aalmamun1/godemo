package main

import (
	"net/http"
)

func loadAPIConfiguration() {
	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/loginError", loginErrorHandler).Methods("GET")
	router.HandleFunc("/logout", logoutHandler).Methods("POST")

	router.HandleFunc("/signup", signupHandler).Methods("POST")
	router.HandleFunc("/signupSuccess", signupSuccessHandler).Methods("GET")

	http.Handle("/", router)
	http.ListenAndServe(":8000", nil)

}
