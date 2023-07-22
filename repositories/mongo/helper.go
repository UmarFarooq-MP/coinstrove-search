package mongo

import (
	"coinstrove-search/internal/core/domain/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

// updateTheLates will make sure the previous value is set to false so the latest will be the one which we will insert
// note :: this function will be called whenever we have the new value for a particular exchange
func updateTheLates(collectionName string, client *mongo.Client) {
	coll := client.Database("coinstrove").Collection(collectionName)
	// filter by latest
	filter := bson.D{{"latest", true}}
	// set the latest to false
	update := bson.D{{"$set", bson.D{{"latest", false}}}}

	_, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println("Failed To Update Latest")
	}
}

func updateDBHelper(collectionName string, client *mongo.Client, message model.Exchange) {
	message.Latest = true
	coll := client.Database("coinstrove").Collection(collectionName)
	_, err := coll.InsertOne(context.TODO(), message)
	if err != nil {
		log.Printf("UpdateDB Error: %v", err)
	}
}
