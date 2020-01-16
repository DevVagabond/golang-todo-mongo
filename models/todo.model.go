package models

import (
	"context"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang-todo-mongo/DB"
	//"go.mongodb.org/mongo-driver/bson"
	"golang-todo-mongo/interfaces"
	"time"
)

func CreateTodo(data interfaces.Todo) (*mongo.InsertOneResult, error) {
	var Todo = DB.Database.Collection("todo")
	var ctx, _ = context.WithTimeout(context.Background(), 3*time.Second)
	return Todo.InsertOne(ctx, data)
}

func FindTodo(id string) []interfaces.Todo {
	list := make([]interfaces.Todo, 0)
	var Todo = DB.Database.Collection("todo")
	var ctx, _ = context.WithTimeout(context.Background(), 3*time.Second)
	userId, _ := primitive.ObjectIDFromHex(id)
	cur, _ := Todo.Find(ctx, bson.D{{"user_id", userId}})
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result interfaces.Todo
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		list = append(list, result)
	}
	return list
}

func UpdateTodo(id string, task interfaces.Todo) (*mongo.UpdateResult, error) {
	var Todo = DB.Database.Collection("todo")
	var ctx, _ = context.WithTimeout(context.Background(), 3*time.Second)
	taskId, _ := primitive.ObjectIDFromHex(id)
	updateTask := bson.M{
		"completed": task.Completed,
	}
	if task.Name != "" {
		updateTask["name"] = task.Name
	}
	if task.Description != "" {
		updateTask["description"] = task.Description
	}

	res, err := Todo.UpdateOne(ctx, bson.D{{"_id", taskId}}, bson.M{
		"$set": updateTask,
	})
	return res, err
}

func RemoveTodo(id string) (*mongo.DeleteResult, error) {
	var Todo = DB.Database.Collection("todo")
	var ctx, _ = context.WithTimeout(context.Background(), 3*time.Second)
	taskId, _ := primitive.ObjectIDFromHex(id)
	res, err := Todo.DeleteOne(ctx, bson.D{{"_id", taskId}})
	return res, err
}
