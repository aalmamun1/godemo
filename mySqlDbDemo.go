package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//type user struct {
//	Id int `json:"id"`
//	FirstName string `json:"firstName"`
//	MiddleName string `json:"middleName"`
//	LastName string `json:"lastName"`
//}

/*
func checkDbConnection() {
	fmt.Println("my sql database connect test")

	db, err := sql.Open("mysql", "connect:connect@tcp(127.0.0.1:3306)/GoDemo")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//insertInUserTable(db)
	//insertInAddressTable(db)

}

func insertInUserTable(db *sql.DB) {
	insert, err := db.Query("insert into USER(id, first_name, middle_name, last_name) values (4, 'Kishan', 'middle1', 'last1')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	fmt.Println("Successfully inserted into USER table")

}

func insertInUserTableWithData(id int64, firstName string) {
	fmt.Println("my sql database connect test")

	db, err := sql.Open("mysql", "connect:connect@tcp(127.0.0.1:3306)/GoDemo")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	insert, err := db.Query("insert into USER(id, first_name, middle_name, last_name) values (?, ?, 'middle1', 'last1')", id, firstName)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	fmt.Println("Successfully inserted into USER table")
}

func insertInAddressTable(db *sql.DB) {
	insert, err := db.Query("insert into ADDRESS(id, address, user_id) values (1, 'California', 1)")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	fmt.Println("Successfully inserted into ADDRESS table")

}
*/

// ********************************* My Sql related code *********************************
func dbConnect() (db *sql.DB) {
	fmt.Println("my sql database connect test")

	db, err := sql.Open("mysql", "connect:connect@tcp(127.0.0.1:3306)/GoDemo")
	// db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/demo")
	if err != nil {
		panic(err.Error())
	}
	// defer db.Close()
	return db
}

func insertInUserTableWithDataApp(userName string, firstName string, lastName string, password string, emailId string) (status string) {
	db := dbConnect()
	defer db.Close()

	insert, err := db.Query("insert into USER_DETAIL (user_name, first_name, last_name, password, last_update_date, email_id) values (?, ?, ?, ?, ?, ?)", userName, lastName, firstName, password, time.Now(), emailId)
	if err != nil {
		status = "Error while insert user: " + userName
		panic(err.Error())
	}
	defer insert.Close()

	status = "User: " + userName + " inserted successfully"
	fmt.Println("Successfully inserted into USER table")
	return
}
