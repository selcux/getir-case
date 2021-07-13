package persistent

import (
	"context"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
)

type Query interface {
	Connect() error
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
	err := db.Connect()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (d *db) Connect() error {
	uri := os.Getenv("MONGO_URI")
	log.Debugf("connecting to: %s", uri)

	ctx := context.TODO()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return errors.Wrap(err, "database connection error")
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return errors.Wrap(err, "unable to ping the primary")
	}

	dbName := os.Getenv("MONGO_DB")
	log.Debugf("DB Name: %s", dbName)
	database := client.Database(dbName)

	collectionName := os.Getenv("MONGO_COLLECTION")
	log.Debugf("collection: %s", collectionName)
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
	/*
		match1 := bson.D{{"$match", bson.D{
			{"createdAt", bson.D{
				{"$gte", bson.D{
					{"$dateFromString", bson.D{
						{"dateString", dq.StartDate},
					}}}},
				{"$lt", bson.D{
					{"$dateFromString", bson.D{
						{"dateString", dq.EndDate},
					}}}},
			}}}}}
	*/
	match1 := bson.D{{"$match", bson.D{
		{"createdAt", bson.D{
			{"$gte", dq.StartDate},
			{"$lt", dq.EndDate},
		}}}}}

	project := bson.D{{"$project", bson.D{
		{"_id", false},
		{"key", true},
		{"createdAt", true},
		{"totalCount", bson.D{
			{"$sum", "$counts"},
		}}}}}

	match2 := bson.D{{"$match", bson.D{
		{"totalCount", bson.D{
			{"$gte", dq.MinCount},
			{"$lt", dq.MaxCount},
		}},
	}}}

	/*
		opts := options.Aggregate()
		opts.SetAllowDiskUse(true)
		opts.SetBatchSize(5)
	*/
	showInfoCursor, err := collection.Aggregate(d.ctx, mongo.Pipeline{match1, project, match2})
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
