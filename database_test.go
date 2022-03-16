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
		t.Fatalf("Createpost is not working. %s", err.Error())
	}

	gp, err := GetPost(db, 1)
	if err != nil {
		t.Fatalf("GetPost is not working.")
	}
	if gp.id != 4 {
		t.Fatalf("Not able to fetch data from getpost function.")
	}

	post, err := FindPost(db, "dsads")
	if err != nil {
		t.Fatalf("FindPost is not working.")
	}
	if post.company != "test44" {
		t.Fatalf("Incorrect ID in FindPost.")
	}
}
