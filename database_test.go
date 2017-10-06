package main

import (
	"gopkg.in/mgo.v2"
	"testing"
)

func setupDB(t *testing.T) *StudentsMongoDB {
	db := StudentsMongoDB{
		"mongodb://localhost",
		"testStudentsDB",
		"students",
	}

	session, err := mgo.Dial(db.DatabaseURL)
	defer session.Close()

	if err != nil {
		t.Error(err)
	}
	return &db
}

func tearDownDB(t *testing.T, db *StudentsMongoDB) {
	session, err := mgo.Dial(db.DatabaseURL)
	defer session.Close()
	if err != nil {
		t.Error(err)
	}

	err = session.DB(db.DatabaseName).DropDatabase()
	if err != nil {
		t.Error(err)
	}
}

func TestStudentsMongoDB_Add(t *testing.T) {
	db := setupDB(t)
	defer tearDownDB(t, db)

	db.Init()
	if db.Count() != 0 {
		t.Error("database not properly initialized. student count() should be 0.")
	}

	student := Student{"Tom", 21, "id1"}
	db.Add(student)

	if db.Count() != 1 {
		t.Error("adding new student failed.")
	}
}

func TestStudentsMongoDB_Get(t *testing.T) {
	db := setupDB(t)
	defer tearDownDB(t, db)

	db.Init()
	if db.Count() != 0 {
		t.Error("database not properly initialized. student count() should be 0.")
	}

	student := Student{"Tom", 21, "id1"}
	db.Add(student)

	if db.Count() != 1 {
		t.Error("adding new student failed.")
	}

	newStudent, ok := db.Get(student.StudentID)
	if !ok {
		t.Error("couldn't find Tom")
	}

	if newStudent.Name != student.Name ||
		newStudent.Age != student.Age ||
		newStudent.StudentID != student.StudentID {
		t.Error("students do not match")
	}
}

func TestStudentsMongoDB_Duplicates(t *testing.T) {
	db := setupDB(t)
	defer tearDownDB(t, db)

	db.Init()
	if db.Count() != 0 {
		t.Error("database not properly initialized. student count() should be 0.")
	}

	student := Student{"Tom", 21, "id1"}
	db.Add(student)

	if db.Count() != 1 {
		t.Error("adding new student failed.")
	}

	// TODO no error handling
	db.Add(student)

	if db.Count() != 1 {
		t.Error("adding new student failed.")
	}
}
