package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Course struct {
	CourseID    primitive.ObjectID `json:"course_id"`
	CourseName  string             `json:"course_name"`
	CoursePrice int                `json:"course_price"`
	Author      *Author            `json:"author"`
}

type Author struct {
	Name    string `json:"name"`
	Website string `json:"website"`
}

func (c Course) IsEmpty() bool {
	return c.CourseName == ""
}
