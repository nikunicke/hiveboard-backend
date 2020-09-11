package mongodb

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/nikunicke/hiveboard"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoURI = os.Getenv("MONGODB")

// MongoDB ...
type MongoDB struct {
	db *mongo.Database
}

// NewMongoDB ...
func NewMongoDB() *MongoDB {
	return &MongoDB{}
}

// Open ...
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

// PostTest ...
func (db *MongoDB) PostTest(collection string) error {
	item := hiveboard.Event{}
	item.Name = "HIVEBOARD TEST 4"
	item.Hiveboard = true
	item.BeginAt = time.Now().AddDate(0, 0, 1)
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

// CheckConnection ...
func (db *MongoDB) CheckConnection() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	return db.db.Client().Ping(ctx, readpref.Primary())
}
