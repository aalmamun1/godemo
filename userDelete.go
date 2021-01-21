package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func deleteUserHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	status := deleteUserByID(id)

	if status == "SUCCESS" {
		fmt.Println("Deleted user with id: ", id, " successfully")
		http.Redirect(response, request, "/internal", 302)
	} else {
		fmt.Fprintf(response, status)
	}
}
