package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB() *mongo.Client {
	var ctx context.Context
	var client *mongo.Client

	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb+srv://anzalasmuhammad:LFSyOMzgrMAtYWit@cluster0.tkvmkvc.mongodb.net/")
	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}
