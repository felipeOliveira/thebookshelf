package data

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database interface {
	List(coll string, filter interface{}, opts *options.FindOptions) (*mongo.Cursor, error)
	Get(coll string, filter interface{}, opts *options.FindOneOptions) *mongo.SingleResult
	Close() error
}

type appDb struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewDatabase(connectionString string) (Database, error) {

	cmdMonitor := &event.CommandMonitor{
		Started: func(_ context.Context, evt *event.CommandStartedEvent) {
			log.Print(evt.Command)
		},
	}

	mongoClient, err := mongo.Connect(context.TODO(), options.Client().SetMonitor(cmdMonitor).ApplyURI(connectionString))

	if err != nil {
		return nil, err
	}

	if err := mongoClient.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
	}

	return &appDb{
		client: mongoClient,
		db:     mongoClient.Database("catalog"),
	}, nil
}

func (database *appDb) List(coll string, filter interface{}, opts *options.FindOptions) (*mongo.Cursor, error) {

	cur, err := database.db.Collection(coll).Find(context.Background(), filter, opts)

	if err != nil {
		return nil, err
	}

	return cur, nil

}

func (database *appDb) Get(coll string, filter interface{}, opts *options.FindOneOptions) *mongo.SingleResult {

	return database.db.Collection(coll).FindOne(context.Background(), filter, opts)
}

func (database *appDb) Close() error {
	return database.client.Disconnect(context.Background())
}
