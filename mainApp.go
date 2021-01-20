package main

import (
	"github.com/gorilla/mux"
)

// server main method

var router = mux.NewRouter()

func main() {
	loadAPIConfiguration()
}
