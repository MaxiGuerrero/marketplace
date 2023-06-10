package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DbConnector struct{
	client *mongo.Client
	database string
}

type Context context.Context

func CreateDbConnector(ctxParent Context,url string,database string) *DbConnector {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatalln(err.Error())
	}
	ctx, cancel := context.WithTimeout(ctxParent,10* time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return &DbConnector{client,database}
}

func (connector DbConnector) GetCollection(collection string)*mongo.Collection{
	return connector.client.Database(connector.database).Collection(collection);
}