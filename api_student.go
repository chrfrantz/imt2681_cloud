package main

import (
	"net/http"
	"encoding/json"
	"fmt"
	"strings"
)

func replyWithAllStudents(w http.ResponseWriter, db *StudentsDB) {
	if db.students == nil {
		json.NewEncoder(w).Encode([]Student{})
	} else {
		a := make([]Student, 0, len(db.students))
		for _, s := range db.students {
			a = append(a, s)
		}
		json.NewEncoder(w).Encode(a)
	}
}


func replyWithStudent(w http.ResponseWriter, db *StudentsDB, id string) {
	// make sure that i is valid
	s, ok := db.Get(id)
	if !ok {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	// handle /student/<id>
	json.NewEncoder(w).Encode(s)
}


func handlerStudent(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if r.Body == nil {
			http.Error(w, "Student POST request must have JSON body", http.StatusBadRequest)
			return
		}
		var s Student
		err := json.NewDecoder(r.Body).Decode(&s)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// check if the student is new
		_, ok := db.Get(s.ID)
		if ok {
			// TODO find a better Error Code (HTTP Status)
			http.Error(w, "Student already exists. Use PUT to modify.", http.StatusBadRequest)
			return
		}
		// new student
		db.Add(s)
		fmt.Fprint(w, "ok") // 200 by default
		return
	case "GET": 
		http.Header.Add(w.Header(), "content-type", "application/json")
		// alternative way:
		// w.Header().Add("content-type", "application/json")
		parts := strings.Split(r.URL.Path, "/")
		// error handling
		if len(parts) != 3 {
			http.Error(w, "Malformed URL", http.StatusBadRequest)
			return
		}
		// handle the request /student/  which will return ALL students as array of JSON objects
		if parts[2] == "" {
			replyWithAllStudents(w, &db)
		} else {
			replyWithStudent(w, &db, parts[2])
		}

	default:
		http.Error(w, "not implemented yet", http.StatusNotImplemented)
		return
	}
}


