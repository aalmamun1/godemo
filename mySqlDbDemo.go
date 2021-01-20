package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func dbConnect() (db *sql.DB) {
	fmt.Println("my sql database connect test")

	db, err := sql.Open("mysql", "connect:connect@tcp(127.0.0.1:3306)/GoDemo")
	if err != nil {
		panic(err.Error())
	}
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
