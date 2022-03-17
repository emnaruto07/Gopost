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

	err = CreatePost(db, JobPost{2, "test44", 22, "sadkfdsfh45dsa", 43000})

	if err != nil {
		t.Fatalf("Createpost is not working. %s", err.Error())
	}

	gp, err := GetPost(db, 2)
	if err != nil {
		t.Fatalf("GetPost is not working.")
	}
	if gp.id != 2 {
		t.Fatalf("Not able to fetch data from getpost function.")
	}

	// post, err := FindPost(db, "dsads")
	// if err != nil {
	// 	t.Fatalf("FindPost is not working.")
	// }
	// if post.company != "test44" {
	// 	t.Fatalf("Incorrect ID in FindPost.")
	// }

	Upost, err := UpdatePost(db, 2, JobPost{2, "test56", 23, "dassasad", 34300})
	if err != nil {
		t.Fatalf("UpdatePost is not working. %s", err)
	}
	if Upost.company != "test56" {
		t.Fatalf("Data has been not updated")
	}

	err = DeletePost(db, 2)
	if err != nil {
		t.Fatalf("Post Not deleted. %s", err)
	}
}
