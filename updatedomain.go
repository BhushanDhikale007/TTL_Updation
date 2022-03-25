package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	//"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://bhushan111:bhushan111@cluster0.hv2f2.mongodb.net/mongod?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	mongodDatabase := client.Database("mongod")
	TenantsCollection := mongodDatabase.Collection("Tenants")

	id, _ := primitive.ObjectIDFromHex("183d5251-56a6-460b-99b9-9686b6c507f1")

	result, err := TenantsCollection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{{"domains", bson.D{{"tgb1.com"}}}}},
		},
	)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated Docs %v \n", result.ModifiedCount)
}
