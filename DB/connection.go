package DB

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var Database *mongo.Database

func InitDb () {
		ctx,_ := context.WithTimeout(context.Background(),3*time.Second);
		client,err := mongo.Connect(ctx,  options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			log.Panic("DATABASE CONNECTION ERROR",err)
		}
		Database = client.Database("testgo")
}

