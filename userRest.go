package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func getAllUsers(response http.ResponseWriter, request *http.Request) {
	userDetails := getAllUserData()

	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(userDetails)

}

func getUser(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	userDetail, status := getUserByID(id)
	if status == "SUCCESS" {
		response.Header().Set("Content-Type", "application/json")
		json.NewEncoder(response).Encode(userDetail)
	} else {
		fmt.Fprintf(response, status)
	}

}

func createUser(response http.ResponseWriter, request *http.Request) {
	var userDetail userDetail
	_ = json.NewDecoder(request.Body).Decode(&userDetail)

	status, id, _ := insertInUserTableWithDataAppAndReturnID(userDetail.UserName, userDetail.FirstName, userDetail.LastName, userDetail.Password, userDetail.EmailID)
	if status == "SUCCESS" {
		//userDetail.ID = int(id)
		userDetail, _ = getUserByID(strconv.Itoa(int(id)))
		response.Header().Set("Content-Type", "application/json")
		json.NewEncoder(response).Encode(userDetail)
	} else {
		fmt.Fprintf(response, status)
	}
}

func updateUser(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	var userDetail userDetail

	if id != "" {
		_ = json.NewDecoder(request.Body).Decode(&userDetail)

		userDetail.ID, _ = strconv.Atoi(id)

		status := updateUserByID(userDetail)

		if status == "SUCCESS" {
			userDetail, _ = getUserByID(id)
			response.Header().Set("Content-Type", "application/json")
			json.NewEncoder(response).Encode(userDetail)
		} else {
			fmt.Fprintf(response, status)
		}
	}

}

func deleteUser(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	userDetail, _ := getUserByID(id)
	status := deleteUserByID(id)

	if status == "SUCCESS" {
		response.Header().Set("Content-Type", "application/json")
		json.NewEncoder(response).Encode(userDetail)
	} else {
		fmt.Fprintf(response, status)
	}
}
