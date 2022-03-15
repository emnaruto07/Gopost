package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type JobPost struct {
	Id      int
	Company string
	Age     int
	Address string
	Salary  int
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func CreatePost(db *sql.DB, jp JobPost) {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into JOBCREATE (jp.Id,jp.Company,jp.Age,jp.Address,jp.Salary) values (?,?,?,?,?)")
	_, err := stmt.Exec(&jp.Id, &jp.Company, &jp.Age, &jp.Address, &jp.Salary)
	checkError(err)
	tx.Commit()
}
func GetPost(db *sql.DB, id int) JobPost {
	// fmt.Println("Get not implemented")
	rows, err := db.Query("select * from JOBCREATE")
	checkError(err)
	for rows.Next() {
		var Post JobPost
		err = rows.Scan(&Post.Id, &Post.Company, &Post.Age, &Post.Address, &Post.Salary)
		checkError(err)
		if Post.Id == id {
			return Post
		}
	}
	return JobPost{}
}
func FindPost() {
	fmt.Println("Find not implemented")
}
func UpdatePost() {
	fmt.Println("Update not implemented")
}
func DeletePost() {
	fmt.Println("Delete not implemented")
}
