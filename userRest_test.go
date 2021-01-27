package main

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

var testRouter *mux.Router

func TestGetUserEndpoint(t *testing.T) {
	// ARRANGE
	testRouter = mux.NewRouter()
	testRouter.HandleFunc("/users/{id}", getUser).Methods("GET")

	response := httptest.NewRecorder()

	// ACT
	request, _ := http.NewRequest("GET", "/users/{id}", nil)
	testRouter.ServeHTTP(response, request)

	// ASSERT
	assert.Equal(t, 200, response.Code, "OK response is expected for user endpoint")
}

func TestGetAllUsersEndpoint(t *testing.T) {
	// ARRANGE
	testRouter = mux.NewRouter()
	testRouter.HandleFunc("/users", getAllUsers).Methods("GET")

	response := httptest.NewRecorder()

	// ACT
	request, _ := http.NewRequest("GET", "/users", nil)
	testRouter.ServeHTTP(response, request)

	// ASSERT
	assert.Equal(t, 200, response.Code, "OK response is expected for all user endpoint")
}

func TestCreateUsersEndpoint(t *testing.T) {
	// ARRANGE
	userDetail := &userDetail{
		UserName:  "abdullah01",
		FirstName: "Abdullah",
		LastName:  "Almamun",
		Password:  "abd@123",
		EmailID:   "abdullah01@gmail.com",
	}
	jsonUser, _ := json.Marshal(userDetail)

	testRouter = mux.NewRouter()
	testRouter.HandleFunc("/users", createUser).Methods("POST")

	response := httptest.NewRecorder()

	// ACT
	request, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonUser))
	//request, _ := http.NewRequest("POST", "/users", nil)
	testRouter.ServeHTTP(response, request)

	// ASSERT
	assert.Equal(t, 200, response.Code, "OK response is expected for create user endpoint")
}

func TestUpdateUsersEndpoint(t *testing.T) {
	// ARRANGE

	userDetail := &userDetail{
		UserName:  "abdullah01",
		FirstName: "Abdullah_UP",
		LastName:  "Almamun_UP",
		Password:  "abd@123_UP",
		EmailID:   "abdullah01_UP@gmail.com",
	}
	jsonUser, _ := json.Marshal(userDetail)

	userDetailTemp := getUserByUserName(userDetail.UserName)
	id := strconv.Itoa(userDetailTemp.ID)

	testRouter = mux.NewRouter()
	testRouter.HandleFunc("/users/{id}", updateUser).Methods("PUT")

	response := httptest.NewRecorder()

	// ACT
	request, _ := http.NewRequest("PUT", "/users/"+id, bytes.NewBuffer(jsonUser))
	testRouter.ServeHTTP(response, request)

	// ASSERT
	assert.Equal(t, 200, response.Code, "OK response is expected for update user endpoint")
}

func TestDeleteUserEndpoint(t *testing.T) {
	// ARRANGE
	testRouter = mux.NewRouter()
	testRouter.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	userDetailTemp := getUserByUserName("dev01")
	id := strconv.Itoa(userDetailTemp.ID)

	response := httptest.NewRecorder()

	// ACT
	request, _ := http.NewRequest("DELETE", "/users/"+id, nil)
	testRouter.ServeHTTP(response, request)

	// ASSERT
	assert.Equal(t, 200, response.Code, "OK response is expected for delete user endpoint")
}
