package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
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

func insertInUserTableWithDataAppAndReturnID(userName string, firstName string, lastName string, password string, emailId string) (status string, id int64, err error) {
	db := dbConnect()
	defer db.Close()

	id = 0
	stmt, err := db.Prepare("INSERT USER_DETAIL SET user_name=?, first_name=?, last_name=?, password=?, last_update_date=?, email_id=?")
	if err != nil {
		status = "Error while insert user: " + userName
		return status, -1, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(userName, lastName, firstName, password, time.Now(), emailId)
	if err != nil {
		status = "Error while insert user: " + userName
		return status, -1, err
	}

	id, err = res.LastInsertId()
	status = "SUCCESS"

	return status, id, err
}

func validateUserNameAndPassword(userName string, password string) (status string) {
	db := dbConnect()
	defer db.Close()

	results, err := db.Query("SELECT id, user_name FROM USER_DETAIL where user_name = ? and password = ?", userName, password)
	if err != nil {
		panic(err.Error())
	}

	if results.Next() {
		status = "successful user login"
	} else {
		status = "username and/or password is invalid"
	}

	fmt.Println(status)
	return
}

type userDetail struct {
	ID         int    `json:"id"`
	UserName   string `json:"userName"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Password   string `json:"password"`
	EmailID    string `json:"emailId"`
	LastUpdate string `json:"lastUpdate"`
}

func getAllUserData() (userDetails []userDetail) {
	db := dbConnect()
	defer db.Close()

	userDetails = make([]userDetail, 0)

	results, err := db.Query("SELECT id, user_name, first_name, last_name, password, email_id, last_update_date FROM USER_DETAIL ")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var userDetail userDetail

		err = results.Scan(&userDetail.ID, &userDetail.UserName, &userDetail.FirstName, &userDetail.LastName, &userDetail.Password, &userDetail.EmailID, &userDetail.LastUpdate)
		if err != nil {
			panic(err.Error())
		}
		userDetails = append(userDetails, userDetail)
		//fmt.Println("In db method: ", userDetail.UserName)
	}

	results.Close()
	fmt.Println("Successfully select records from USER_DETAIL table")
	return
}

func getAllUserDataByUsername(searchUserName string) (userDetails []userDetail) {
	db := dbConnect()
	defer db.Close()

	userDetails = make([]userDetail, 0)

	results, err := db.Query("SELECT id, user_name, first_name, last_name, password, email_id, last_update_date FROM USER_DETAIL where user_name like ? ", searchUserName)
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var userDetail userDetail

		err = results.Scan(&userDetail.ID, &userDetail.UserName, &userDetail.FirstName, &userDetail.LastName, &userDetail.Password, &userDetail.EmailID, &userDetail.LastUpdate)
		if err != nil {
			panic(err.Error())
		}
		userDetails = append(userDetails, userDetail)
		fmt.Println("In db method: ", userDetail.UserName)
	}
	fmt.Println("size is: ", len(userDetails))
	results.Close()
	fmt.Println("Successfully select records from USER_DETAIL table")
	return
}

func deleteUserByID(id string) (status string) {
	db := dbConnect()
	defer db.Close()

	_, err := db.Query("DELETE FROM USER_DETAIL where id = ? ", id)
	if err != nil {
		status = "Error while deleting user with id: " + id
		panic(err.Error())
	}
	status = "SUCCESS"

	return
}

func getUserByID(id string) (userDetail userDetail, status string) {
	db := dbConnect()
	defer db.Close()

	results, err := db.Query("SELECT id, user_name, first_name, last_name, password, email_id, last_update_date FROM USER_DETAIL where id = ? ", id)
	if err != nil {
		status = "Error while selecting user with id: " + id
		fmt.Println(status)
		panic(err.Error())
	}

	if results.Next() {
		err = results.Scan(&userDetail.ID, &userDetail.UserName, &userDetail.FirstName, &userDetail.LastName, &userDetail.Password, &userDetail.EmailID, &userDetail.LastUpdate)
		if err != nil {
			status = "Error while selecting user with id: " + id
			panic(err.Error())
		}
		status = "SUCCESS"
		fmt.Println("Successfully selected record from USER_DETAIL table for user id: " + id)
	}

	results.Close()
	return
}

func getUserByUserName(userName string) (userDetail userDetail) {
	db := dbConnect()
	defer db.Close()

	results, err := db.Query("SELECT id, user_name, first_name, last_name, password, email_id, last_update_date FROM USER_DETAIL where user_name = ? ", userName)
	if err != nil {
		panic(err.Error())
	}

	if results.Next() {
		err = results.Scan(&userDetail.ID, &userDetail.UserName, &userDetail.FirstName, &userDetail.LastName, &userDetail.Password, &userDetail.EmailID, &userDetail.LastUpdate)
		if err != nil {
			panic(err.Error())
		}
	}

	results.Close()
	fmt.Println("Successfully selected record from USER_DETAIL table for user id: " + userName)
	return
}

func updateUserByID(u userDetail) (status string) {
	db := dbConnect()
	defer db.Close()

	results, err := db.Query("Update USER_DETAIL set first_name=?, last_name=?, password=?, email_id=?, last_update_date=? where id = ? ", u.FirstName, u.LastName, u.Password, u.EmailID, time.Now(), u.ID)
	if err != nil {
		status = "Error while updating user with id: " + strconv.Itoa(u.ID)
		panic(err.Error())
	}

	status = "SUCCESS"

	results.Close()
	fmt.Println("Successfully updated record from USER_DETAIL table for user id: ", u.ID)
	return
}
