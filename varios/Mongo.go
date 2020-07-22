package varios

import (
	"aplicacaoExemplo/domain"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func saveMongo(person domain.Person) *mongo.InsertOneResult {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("person").Collection("persons")

	result, err := collection.InsertOne(context.TODO(), &person)

	return result

}
