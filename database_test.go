package main

import "testing"

func TestDatabase(t *testing.T) {
	db, err := CreateConnection()
	if err != nil {
		t.Fatalf("Database is not connecting")
	}

	_, err = db.Exec("create table if not exists JOBCOMPANY(id integer, company char, age integer, address char, salary integer)")
	if err != nil {
		t.Fatalf("tables are not created")
	}

	err = CreatePost(db, JobPost{1, "test44", 22, "sadkfdsfh45dsa", 43000})

	if err != nil {
		t.Fatalf("Createpost is not working.")
	}

	_, err = GetPost(db, 4)
	if err != nil {
		t.Fatalf("GetPost is not working.")
	}

	post, err := FindPost(db, "dsads")
	if err != nil {
		t.Fatalf("FindPost is not working.")
	}
	if post.id != 1 {
		t.Fatalf("Incorrect ID in FindPost")
	}
}
