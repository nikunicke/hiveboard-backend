package mongodb

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nikunicke/hiveboard"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoURI = os.Getenv("MONGODB")

type MongoDB struct {
	db *mongo.Database
}

func NewMongoDB() *MongoDB {
	return &MongoDB{}
}

func (db *MongoDB) Open(name string) error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db.db = client.Database(name)
	db.db.Client().Connect(ctx)
	return nil
}

func (db *MongoDB) PostTest(collection string) error {
	item := hiveboard.Event{}
	item.Name = "Thruster"
	col := db.db.Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := col.InsertOne(ctx, item)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

func (db *MongoDB) FindAll(collection string) {
	var results []bson.M

	cursor, err := db.db.Collection(collection).Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		fmt.Println(result)
	}
}

func (db *MongoDB) CheckConnection() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	return db.db.Client().Ping(ctx, readpref.Primary())
}
