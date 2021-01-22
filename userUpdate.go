package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func updateUserFormHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	userDetail := getUserByID(id)

	fmt.Println("To update user: ", userDetail)
	fmt.Fprintf(response, updatePage, userDetail.ID, userDetail.UserName, userDetail.Password, userDetail.FirstName, userDetail.LastName, userDetail.EmailID)
}

func updateUserHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	var userDetail userDetail

	if id != "" {
		userDetail.ID, _ = strconv.Atoi(id)
		userDetail.UserName = request.FormValue("signupName")
		userDetail.FirstName = request.FormValue("firstName")
		userDetail.LastName = request.FormValue("lastName")
		userDetail.Password = request.FormValue("signupPassword")
		userDetail.EmailID = request.FormValue("email")

		status := updateUserByID(userDetail)

		if status == "SUCCESS" {
			fmt.Println("Deleted user with id: ", id, " successfully")
			http.Redirect(response, request, "/internalUser", 302)
		} else {
			fmt.Fprintf(response, status)
		}
	}
}
