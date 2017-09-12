package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func handlerHello(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 4 {
		status := 400
		http.Error(w, http.StatusText(status), status)
		return
	}
	name := parts[2]
	fmt.Fprintln(w, parts)
	fmt.Fprintf(w, "Hello %s %s!\n", name, parts[3])
}

func replyWithAllStudents(w http.ResponseWriter, db *StudentsDB) {
	if db.students == nil {
		json.NewEncoder(w).Encode([]Student{})
	} else {
		json.NewEncoder(w).Encode(db.students)
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

	if r.Method == "POST" {
		http.Error(w, "not implemented yet", http.StatusNotImplemented)

		return
	}
	// if r.Method == "GET"
	http.Header.Add(w.Header(), "content-type", "application/json")
	// alternative way:
	// w.Header().Add("content-type", "application/json")
	parts := strings.Split(r.URL.Path, "/")
	// error handling
	if len(parts) != 3 {
		// handle error
		return
	}
	// handle the /student/
	if parts[2] == "" {
		replyWithAllStudents(w, &db)
	} else {
		replyWithStudent(w, &db, parts[2])
	}
}

// -------------
var db StudentsDB
// -------------

func main() {
	db = StudentsDB{}
	db.Init()

	http.HandleFunc("/hello/", handlerHello)
	http.HandleFunc("/student/", handlerStudent)
	http.ListenAndServe(":8080", nil)
}
