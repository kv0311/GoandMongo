package mongoconfig

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongo(db, collect string) *mongo.Collection {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	//set client options
	client, _ := mongo.Connect(context.TODO(), clientOptions)
	collection := client.Database(db).Collection(collect)

	return collection
}
