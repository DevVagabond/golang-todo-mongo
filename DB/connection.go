package DB

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var Database *mongo.Database

func InitDb () {
		ctx,_ := context.WithTimeout(context.Background(),3*time.Second);
		client,_ := mongo.Connect(ctx,  options.Client().ApplyURI("mongodb://localhost:27017"))
		Database = client.Database("testgo")
}

