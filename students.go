package main

/*
Student represents the main persistent data structure.
It is of the form:
{
	"name": <value>, 	e.g. "Tom"
	"age": <value>		e.g. 21
	"studentid": <value>		e.c. "id0"
}
*/
type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	StudentID   string `json:"studentid"`
}

/*
StudentsDB is the handle to students in-memory storage.
*/
type StudentsDB struct {
	students map[string]Student
}

/*
Init initializes the in-memory storage.
*/
func (db *StudentsDB) Init() {
	db.students = make(map[string]Student)
}

/*
Add adds new students to the storage.
*/
func (db *StudentsDB) Add(s Student) {
	db.students[s.StudentID] = s
}

/*
Count returns the current count of the students in in-memory storage.
*/
func (db *StudentsDB) Count() int {
	return len(db.students)
}

/*
Get returns a student with a given ID or empty student struct.
*/
func (db *StudentsDB) Get(keyID string) (Student, bool) {
	s, ok := db.students[keyID]
	return s, ok
}
