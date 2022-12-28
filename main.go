package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Course struct {
	CourseID    string  `json:"course_id"`
	CourseName  string  `json:"course_name"`
	CoursePrice int     `json:"course_price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Name    string `json:"name"`
	Website string `json:"website"`
}

// fake DB

var courses []Course

func (c *Course) isEmpty() bool {
	return c.CourseName == ""
}

func main() {
	r := mux.NewRouter()
	//seed  data
	courses = append(courses, Course{CourseID: "2", CourseName: "React", CoursePrice: 299,
		Author: &Author{Name: "Paras", Website: "go.dev"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getCourseById).Methods("GET")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course", addCourse).Methods("POST")
	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to API by learncode online</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Response")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getCourseById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get course by id")
	w.Header().Set("Content-Type", "application/json")
	// get Id from request
	params := mux.Vars(r)
	//lop through courses
	for _, Course := range courses {
		if Course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(Course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course foudn with id")
	return
}

func addCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating new course")
	w.Header().Set("COntent-Type", "application/json")
	// check if body is empty

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}
	// check for {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {
		json.NewEncoder(w).Encode("no data found in ")
		return
	}

	// genrate uniqueId , and convert id to string
	// add course to courses slice

	rand.Seed(time.Now().UnixNano())
	course.CourseID = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// loop.id,remove, add with id
	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(course)
			course.CourseID = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deling course")
	params := mux.Vars(r)
	//loop / id / remove
	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("course deleted ")
			return
		}
	}

	json.NewEncoder(w).Encode("no id found")
	return
}
