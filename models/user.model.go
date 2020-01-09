package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"golang-todo-mongo/DB"

	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/bson"
	"golang-todo-mongo/interfaces"
	"time"
)

 func CreateUser(data interfaces.User) (*mongo.InsertOneResult,error) {
 var User = DB.Database.Collection("users")
 var ctx,_ = context.WithTimeout(context.Background(),3*time.Second);
 	return User.InsertOne(ctx,data)
 }

 func FindUser(email string) (*mongo.SingleResult) {
 	var User = DB.Database.Collection("users")
 	var ctx,_ = context.WithTimeout(context.Background(),3*time.Second)
 	res := User.FindOne(ctx, bson.D{{"email" , email}})
	return res
 }
