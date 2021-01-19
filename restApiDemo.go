package main

import (
	"net/http"
)

/*
func restTest() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/{id}/{name}", postApiRequest).Methods("POST")
	http.ListenAndServe(":8000", router)
	fmt.Println("rest worked")
}

func postApiRequest(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Get Api function called")

	vars := mux.Vars(request)
	id := vars["id"]
	name := vars["name"]
	fmt.Fprint(writer, "id: "+id+", name: "+name)

	var intId int64
	intId, _ = strconv.ParseInt(id, 10, 64)

	insertInUserTableWithData(intId, name)
}

*/

func loadAPIConfiguration() {

	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/loginError", loginErrorHandler).Methods("GET")
	router.HandleFunc("/logout", logoutHandler).Methods("POST")

	router.HandleFunc("/signup", signupHandler).Methods("POST")
	router.HandleFunc("/signupSuccess", signupSuccessHandler).Methods("GET")

	http.Handle("/", router)
	http.ListenAndServe(":8000", nil)

}
