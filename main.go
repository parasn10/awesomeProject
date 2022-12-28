package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Course struct {
	CourseID    string  `json:"course_id"`
	CourseName  string  `json:"course_name"`
	CoursePrice string  `json:"course_price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Name    string `json:"name"`
	Website string `json:"website"`
}

// fake DB

var course []Course

func (c *Course) isEmpty() bool {
	return c.CourseID == "" && c.CourseName == ""
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to API by learncode online</h1>"))
}

// serve home toutr

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Response")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(course)
}
