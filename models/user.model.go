package models

import (
	"context"
	"golang-todo-mongo/DB"
	"golang-todo-mongo/interfaces"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

 func CreateUser(data interfaces.User) (*mongo.InsertOneResult,error) {
 var User = DB.Database.Collection("users");
 var ctx,_ = context.WithTimeout(context.Background(),3*time.Second);
 	return User.InsertOne(ctx,data)
 }
