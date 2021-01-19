package main

import (
	"github.com/gorilla/mux"
)

// server main method

var router = mux.NewRouter()

func main() {
	loadAPIConfiguration()
}

/*
func main() {
	fmt.Println("Hello World")

	a , b := calculateSumAndSubstraction1(9 , 4)
	fmt.Println("Sum is: ", a, ", Substraction is: ", b)
	stringTest("Abdullah Rehman")

	checkDbConnection()
	restTest()
}

func calculateSumAndSubstraction(first int, second int) (sum int, substraction int) {
	sum = first + second
	substraction = first - second
	return sum, substraction
}

func calculateSumAndSubstraction1(first, second int) (sum, substraction int) {
	sum = first + second
	substraction = first - second

	return
}

func stringTest(name string) {
	lengthOfName := len(name)
	fmt.Println("Length of string ", name, " is: ", lengthOfName)

}

*/
