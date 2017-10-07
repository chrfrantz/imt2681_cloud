package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
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

// -----------------
var db StudentsStorage

// -----------------

func main() {
	// Using in-memory storage
	// db = &StudentsDB{}

	// Using MongoDB based storage
	db = &StudentsMongoDB{}

	db.Init()

	http.HandleFunc("/hello/", handlerHello)
	http.HandleFunc("/student/", handlerStudent)
	http.ListenAndServe(":8080", nil)
}
