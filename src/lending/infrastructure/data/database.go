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
	List(coll string, filter interface{}, opts *options.FindOptions, data interface{}) error
	Get(coll string, filter interface{}, opts *options.FindOneOptions, data interface{}) error
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

func (database *appDb) List(coll string, filter interface{}, opts *options.FindOptions, data interface{}) error {

	cur, err := database.db.Collection(coll).Find(context.Background(), filter, opts)

	if err != nil {
		return err
	}

	err = cur.All(context.Background(), data)

	if err != nil && err != mongo.ErrNoDocuments {
		return err
	}

	return nil

}

func (database *appDb) Get(coll string, filter interface{}, opts *options.FindOneOptions, data interface{}) error {

	if err := database.db.Collection(coll).FindOne(context.Background(), filter, opts).
		Decode(data); err != nil && err != mongo.ErrNoDocuments {
		return err
	}

	return nil

}

func (database *appDb) Close() error {
	return database.client.Disconnect(context.Background())
}
