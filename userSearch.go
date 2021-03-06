package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func searchUserHandler(response http.ResponseWriter, request *http.Request) {
	userName := getUserName(request)

	if userName != "" {
		var table = ""
		searchUserName := "%" + request.FormValue("search") + "%"
		userDetails := getAllUserDataByUsername(searchUserName)

		for _, user := range userDetails {
			deleteHref := "/delete/" + strconv.Itoa(user.ID)
			updateHref := "/update/" + strconv.Itoa(user.ID)
			table = table + "<tr>" +
				"<td>" + strconv.Itoa(user.ID) + "</td>" +
				"<td>" + user.UserName + "</td>" +
				"<td>" + user.FirstName + "</td>" +
				"<td>" + user.LastName + "</td>" +
				"<td>" + user.EmailID + "</td>" +
				"<td>" + user.LastUpdate + "</td>" +
				"<td> <a href=" + deleteHref + ">Delete</a></td>" +
				"<td> <a href=" + updateHref + ">Update</a></td>" +
				"</tr>"
		}

		finalPage := internalPageTop + table + internalPageDown
		fmt.Fprintf(response, finalPage)

	} else {
		http.Redirect(response, request, "/", 302)
	}
}
