package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoURI = "mongodb+srv://npimenof:PWD@cluster0-z03gj.mongodb.net/test?retryWrites=true&w=majority"

type MongoDB struct {
	db *mongo.Database
}

func NewMongoDB() *MongoDB {
	return &MongoDB{}
}

func (db *MongoDB) Open(name string) error {
	// Getenv() and stuff here
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
	col := db.db.Client().Database("hiveboard").Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := col.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

func (db *MongoDB) CheckConnection() error {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	return db.db.Client().Ping(ctx, readpref.Primary())
}

// func Test() {
// 	client, err := mongo.NewClient(options.Client().ApplyURI(
// 		"mongodb+srv://npimenof:bllr54tGsFnbxPgJ@cluster0-z03gj.mongodb.net/test?retryWrites=true&w=majority",
// 	))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
// 	defer cancel()
// 	if err = client.Connect(ctx); err != nil {
// 		log.Fatal(err)
// 	}
// 	filter := bson.M{"hello": "world"}
// 	collection := client.Database("hiveboard").Collection("events")
// 	res, err := collection.DeleteOne(context.Background(), filter)
// 	// res, err := collection.InsertOne(context.Background(), bson.M{"hello": "world"})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("%d Item(s) Deleted\n", res.DeletedCount)
// 	// id := res.InsertedID
// 	// fmt.Printf("%v\n", id)
// }
