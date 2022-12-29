package Router

import (
	"awesomeProject/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	// routing
	r.HandleFunc("/", controllers.ServeHome).Methods("GET")
	r.HandleFunc("/courses", controllers.GetAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", controllers.GetCourseById).Methods("GET")
	r.HandleFunc("/course/{id}", controllers.DeleteCourse).Methods("DELETE")
	r.HandleFunc("/course/{id}", controllers.UpdateCourse).Methods("PUT")
	r.HandleFunc("/course", controllers.AddCourse).Methods("POST")
	return r
}
