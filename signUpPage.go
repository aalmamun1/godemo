package main

import (
	"fmt"
	"net/http"
)

// sgnup handler
func signupHandler(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("signupName")
	firstName := request.FormValue("firstName")
	lastName := request.FormValue("lastName")
	pass := request.FormValue("signupPassword")
	emailId := request.FormValue("email")
	redirectTarget := "/"
	if name != "" && pass != "" {
		// .. check credentials ..
		setSession(name, response)

		insertInUserTableWithDataApp(name, firstName, lastName, pass, emailId)
		//status := insertInUserTableWithDataApp(name, firstName, lastName, pass)
		//response.Header().Set("status", status)
		redirectTarget = "/signupSuccess"
	}
	http.Redirect(response, request, redirectTarget, 302)
}

func signupSuccessHandler(response http.ResponseWriter, request *http.Request) {
	//status := response.Header().Get("status")
	//fmt.Fprint(response, status)
	fmt.Fprint(response, "User Inserted successfully")
}
