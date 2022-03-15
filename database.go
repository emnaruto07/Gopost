package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type JobPost struct {
	id      int
	company string
	age     int
	address string
	salary  int
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func CreateConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "JobPost.db")

	if err != nil {
		panic(err)
	}
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	return db

}

func CreatePost(db *sql.DB, jp JobPost) {

	connection := CreateConnection()
	defer connection.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into JOBCREATE (id,company,age,address,salary) values (?,?,?,?,?)")
	_, err := stmt.Exec(&jp.id, &jp.company, &jp.age, &jp.address, &jp.salary)
	checkError(err)
	tx.Commit()
}

func GetPost(db *sql.DB, id int) JobPost {

	connection := CreateConnection()
	defer connection.Close()

	rows, err := db.Query("select * from JOBCREATE")
	checkError(err)
	for rows.Next() {
		var Post JobPost
		err = rows.Scan(&Post.id, &Post.company, &Post.age, &Post.address, &Post.salary)
		checkError(err)
		if Post.id == id {
			return Post
		}
	}
	return JobPost{}
}

func FindPost(db *sql.DB, company string) JobPost {

	connection := CreateConnection()
	defer connection.Close()
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("select * from JOBCREATE where company=?")
	_, err := stmt.Exec(company)

	checkError(err)
	for rows.Next() {
		var Post JobPost
		err = rows.Scan(&Post.id, &Post.company, &Post.age, &Post.address, &Post.salary)
		checkError(err)
		if Post.company == company {
			return Post
		}
	}
	return JobPost{}

}

func UpdatePost(db *sql.DB, id int, jp JobPost) {

	connection := CreateConnection()
	defer connection.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("update JOBCREATE set company=?,age=?,address=?,salary=? where id=?")
	_, err := stmt.Exec(&jp.id, &jp.company, &jp.age, &jp.address, id)
	checkError(err)
	tx.Commit()
}

func DeletePost(db *sql.DB, id int) {

	connection := CreateConnection()
	defer connection.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("delete from JOBCREATE where id=?")
	_, err := stmt.Exec(id)
	checkError(err)
	tx.Commit()
}
