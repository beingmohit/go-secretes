package database

import (
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
)

type Client struct {
	connection *bongo.Connection
}

func NewClient(uri string, database string) *Client {
	config := &bongo.Config{
		ConnectionString: uri,
		Database: database,
	}

	connection, error := bongo.Connect(config)

	if error != nil {
		panic("Couldn't connect to mongo")
	}

	client := Client{connection: connection}

	return &client
}

func (client *Client) Save(collection string, model bongo.Document) error {
	error := client.connection.Collection(collection).Save(model)

	return error
}

func (client *Client) Find(collection string, model bongo.Document, conditions bson.M) error {
	error := client.connection.Collection(collection).FindOne(conditions, model)

	return error
}
