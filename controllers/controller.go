package controllers

import (
	"awesomeProject/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

const connectionString = ""
const dbName = "courses"
const colName = "course"

var collection *mongo.Collection

// fake DB

var courses []model.Course

func init() {
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Mongo db connection created")
	collection = client.Database(dbName).Collection(colName)
}

// controllers
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to API by learncode online</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Response")
	w.Header().Set("Content-Type", "application/json")
	cursor, err = collection.Find(context.Background(), bson.M{})
	defer curs
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
	// getting request body
	var course model.Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("no data found in ")
		return
	}

	// dummy data appaproach (seeding manually)
	//// genrate uniqueId , and convert id to string
	//// add course to courses slice
	//
	//rand.Seed(time.Now().UnixNano())
	//course.CourseID = strconv.Itoa(rand.Intn(100))
	//courses = append(courses, course)

	// adding mongo collection
	inserted, err := collection.InsertOne(context.Background(), course)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("data inserted ", inserted)
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
