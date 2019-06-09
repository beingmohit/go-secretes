package database

import (
	"github.com/go-bongo/bongo"
)

type Client struct {
	connection *bongo.Connection
}

func CreateClient(uri string, database string) *Client {
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