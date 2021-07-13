package mongo_client

import (
	"context"
	"face-recognition-service/internals/models"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"os/signal"
	"time"
)

type mongoDB struct {
	config models.Config
	collection *mongo.Collection
}

func Connect(config models.Config) *mongoDB {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.MongoUri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("face-recognition-service")
	personsCollection := db.Collection("persons")

	go func() {
		sigChan := make(chan os.Signal)
		signal.Notify(sigChan, os.Interrupt)
		signal.Notify(sigChan, os.Kill)

		sig := <- sigChan
		log.Println("Received terminate, graceful mongo disconnection", sig)

		client.Disconnect(ctx)
	}()

	log.Println("connected to mongo")

	return &mongoDB{config, personsCollection}
}

func (this *mongoDB) InsertOne(ctx context.Context, document interface{}) error {
	_, err := this.collection.InsertOne(ctx, document)
	return err
}

func (this *mongoDB) Find(ctx context.Context, filter interface{}) (*[]models.Person, error) {
	cursor, _ := this.collection.Find(ctx, filter)
	var results []models.Person
	if err := cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	fmt.Println(results)

	return &results, nil
}

func (this *mongoDB) Count(ctx context.Context, filter interface{}) (int64, error) {
	return this.collection.CountDocuments(ctx, filter)
}
