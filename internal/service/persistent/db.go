package persistent

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
)

type Query interface {
	init() error
	Get(dq *DataQuery) ([]DataQueryRecord, error)
	Disconnect() error
}

type db struct {
	ctx        context.Context
	client     *mongo.Client
	collection *mongo.Collection
}

func NewDb() (Query, error) {
	db := &db{}
	err := db.init()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (d *db) init() error {
	uri := os.Getenv("MONGO_URI")

	ctx := context.TODO()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return errors.Wrap(err, "database connection error")
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return errors.Wrap(err, "unable to ping the primary")
	}

	dbName := os.Getenv("MONGO_DB")
	database := client.Database(dbName)

	collectionName := os.Getenv("MONGO_COLLECTION")
	collection := database.Collection(collectionName)

	d.ctx = ctx
	d.client = client
	d.collection = collection

	return nil
}

func (d *db) Get(dq *DataQuery) ([]DataQueryRecord, error) {
	client := d.client
	collection := d.collection

	if err := client.Ping(d.ctx, readpref.Primary()); err != nil {
		return nil, errors.Wrap(err, "unable to ping the primary")
	}

	showInfoCursor, err := collection.Aggregate(d.ctx, aggregateRecordsWithCountSum(dq))
	if err != nil {
		return nil, errors.Wrap(err, "unexpected behavior in query")
	}

	var showsWithInfo []DataQueryRecord

	if err = showInfoCursor.All(d.ctx, &showsWithInfo); err != nil {
		return nil, errors.Wrap(err, "unable to iterate over the results")
	}

	return showsWithInfo, nil
}

func (d *db) Disconnect() error {
	err := d.client.Disconnect(d.ctx)
	if err != nil {
		return errors.Wrap(err, "database disconnecting error")
	}

	return nil
}
