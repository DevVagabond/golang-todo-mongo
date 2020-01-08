package interfaces

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	FirstName  	string `json:"first_name" bson:"first_name"`
	LatsName	string 	`json:"lats_name" bson:"last_name"`
	Password	string	`json:"password" bson:"password"`
}


type Todo struct {
	UserId      primitive.ObjectID `json:"user_id" bson:"user_id"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Completed   bool               `json:"completed" bson:"completed"`
}