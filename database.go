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

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("insert into JOBCOMPANY (id,company,age,address,salary) values (?,?,?,?,?)")
	if err != nil {
		return err
	}
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

	tx, err := db.Begin()
	if err != nil {
		return JobPost{}, err
	}
	stmt, err := tx.Prepare("select * from JOBCOMPANY where id=?")
	if err != nil {
		return JobPost{}, err
	}

	rows, err := stmt.Query(id)
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

	tx, err := db.Begin()
	if err != nil {
		return JobPost{}, err
	}
	stmt, err := tx.Prepare("select * from JOBCOMPANY where company=?")
	if err != nil {
		return JobPost{}, err
	}

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

	tx, err := db.Begin()
	if err != nil {
		return JobPost{}, err
	}

	stmt, err := tx.Prepare("update JOBCOMPANY set company=?,age=?,address=?,salary=? where id=?")
	if err != nil {
		return JobPost{}, err
	}
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

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("delete from JOBCOMPANY where id=?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	tx.Commit()

	return err
}
