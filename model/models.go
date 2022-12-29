package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Course struct {
	CourseID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CourseName  string             `json:"course_name"`
	CoursePrice int                `json:"course_price"`
	Author      *Author            `json:"author"`
	CourseLive  bool               `json:"isCourseLiuve"`
}

type Author struct {
	Name    string `json:"name"`
	Website string `json:"website"`
}

func (c Course) IsEmpty() bool {
	return c.CourseName == ""
}
