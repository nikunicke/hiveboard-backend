package mongodb

import (
	"context"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func TestOpen(t *testing.T) {
	db := NewMongoDB()
	got := db.Open("hiveboard")
	if got != nil {
		t.Errorf("Open(name string) = %s; want nil", got.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	ping := db.db.Client().Ping(ctx, readpref.Primary())
	if ping != nil {
		t.Errorf("Ping test failed: %s.\nMake sure your IP address has access to MongoDB", ping.Error())
	}
}
