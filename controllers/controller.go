package controllers

import (
	"awesomeProject/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

const connectionString = "mongodb+srv://parasn10:Paras%4010@gocluster.hgvqz8z.mongodb.net/?retryWrites=true&w=majority"
const dbName = "courses"
const colName = "course"

var collection *mongo.Collection

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
func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to API by learncode online</h1>"))
}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Response")
	w.Header().Set("Content-Type", "application/json")
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var courses []model.Course
	for cur.Next(context.Background()) {
		var course model.Course
		err := cur.Decode(&course)
		if err != nil {
			log.Fatal(err)
		}
		courses = append(courses, course)
	}
	defer cur.Close(context.Background())
	json.NewEncoder(w).Encode(courses)
	return
}

func GetCourseById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get course by id")
	w.Header().Set("Content-Type", "application/json")
	// get Id from request
	params := mux.Vars(r)

	//// getting data from seed slice and returning back course
	////lop through courses
	////for _, Course := range courses {
	////	if Course.CourseID == params["id"] {
	////		json.NewEncoder(w).Encode(Course)
	////		return
	////	}
	////}
	// getting course by id from mongo collection
	objectId, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	log.Println(objectId)
	filter := bson.M{"_id": objectId}
	result := collection.FindOne(context.Background(), filter)
	log.Println(result)
	var course model.Course
	json.NewEncoder(w).Encode(result.Decode(course))
	return
}

func AddCourse(w http.ResponseWriter, r *http.Request) {
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
		log.Fatal(err) // more controlled version of panic
	}
	fmt.Println("data inserted ", inserted.InsertedID)
	json.NewEncoder(w).Encode(course)
	return
}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// loop.id,remove, add with id
	//for index, course := range courses {
	//	if course.CourseID == params["id"] {
	//		courses = append(courses[:index], courses[index+1:]...)
	//		var course Course
	//		_ = json.NewDecoder(r.Body).Decode(course)
	//		course.CourseID = params["id"]
	//		courses = append(courses, course)
	//		json.NewEncoder(w).Encode(course)
	//		return
	//	}
	//}
	//
	id, _ := primitive.ObjectIDFromHex(params["id"])
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"isCourse=Live": true}}
	// bson.M= shorter and clearer result and bson.D = when we want order values
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result updated", result)
}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deling course")
	params := mux.Vars(r)
	////loop / id / remove
	////for index, course := range courses {
	////	if course.CourseID == params["id"] {
	////		courses = append(courses[:index], courses[index+1:]...)
	////		json.NewEncoder(w).Encode("course deleted ")
	////		return
	////	}
	////}
	//
	objectId, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	collection.DeleteOne(context.Background(), bson.M{"_id": objectId})
	json.NewEncoder(w).Encode("record deleted")
	return
}
