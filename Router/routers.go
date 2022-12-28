package Router

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getCourseById).Methods("GET")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course", addCourse).Methods("POST")
	return router
}
