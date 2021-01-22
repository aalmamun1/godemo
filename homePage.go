package main

import (
	"fmt"
	"net/http"
)

func indexPageHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, indexPage)
}
