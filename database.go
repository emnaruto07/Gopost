package main

import (
	"database/sql"

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

func CreatePost(db *sql.DB, jp JobPost) {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into JOBCREATE (jp.id,jp.company,jp.age,jp.address,jp.salary) values (?,?,?,?,?)")
	_, err := stmt.Exec(&jp.id, &jp.company, &jp.age, &jp.address, &jp.salary)
	checkError(err)
	tx.Commit()
}
func GetPost(db *sql.DB, id int) JobPost {
	// fmt.Println("Get not implemented")
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
	// fmt.Println("Find not implemented")
	rows, err := db.Query("select * from JOBCREATE where company=?")
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
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("update JOBCREATE set jp.company=?,jp.age=?,jp.address=?,jp.salary=? where jp.id=?")
	_, err := stmt.Exec(&jp.id, &jp.company, &jp.age, &jp.address, id)
	checkError(err)
	tx.Commit()
}
func DeletePost(db *sql.DB, id int) {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("delete from JOBCREATE where id=?")
	_, err := stmt.Exec(id)
	checkError(err)
	tx.Commit()
}
