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

func CreateConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "JobPost.db")

	if err != nil {
		return nil, err
	}
	err = db.Ping()

	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected!")

	return db, nil

}

func CreatePost(db *sql.DB, jp JobPost) error {

	connection, err := CreateConnection()
	if err != nil {
		return err
	}

	defer connection.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into JOBCREATE (id,company,age,address,salary) values (?,?,?,?,?)")
	_, err = stmt.Exec(&jp.id, &jp.company, &jp.age, &jp.address, &jp.salary)

	if err != nil {
		return err
	}

	tx.Commit()

	return err
}

func GetPost(db *sql.DB, id int) (JobPost, error) {

	connection, err := CreateConnection()

	if err != nil {
		return JobPost{}, err
	}

	defer connection.Close()

	rows, err := db.Query("select * from JOBCREATE")

	if err != nil {
		return JobPost{}, err
	}
	for rows.Next() {
		var Post JobPost
		err = rows.Scan(&Post.id, &Post.company, &Post.age, &Post.address, &Post.salary)

		if err != nil {
			return JobPost{}, err
		}

		if Post.id == id {
			return Post, nil
		}
	}
	return JobPost{}, nil
}

func FindPost(db *sql.DB, company string) (JobPost, error) {

	connection, _ := CreateConnection()
	defer connection.Close()
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("select * from JOBCREATE where company=?")
	rows, err := stmt.Query(company)

	if err != nil {
		return JobPost{}, err
	}
	for rows.Next() {
		var Post JobPost
		err = rows.Scan(&Post.id, &Post.company, &Post.age, &Post.address, &Post.salary)
		if err != nil {
			return JobPost{}, err
		}
		if Post.company == company {
			return Post, err
		}
	}
	return JobPost{}, err

}

func UpdatePost(db *sql.DB, id int, jp JobPost) (JobPost, error) {

	connection, err := CreateConnection()

	if err != nil {
		return JobPost{}, err
	}

	defer connection.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("update JOBCREATE set company=?,age=?,address=?,salary=? where id=?")
	_, err = stmt.Exec(&jp.id, &jp.company, &jp.age, &jp.address, id)
	if err != nil {
		return JobPost{}, err
	}
	tx.Commit()

	return JobPost{}, err
}

func DeletePost(db *sql.DB, id int) error {

	connection, err := CreateConnection()
	if err != nil {
		return err
	}
	defer connection.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("delete from JOBCREATE where id=?")
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	tx.Commit()

	return err
}
