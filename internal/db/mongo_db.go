package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDBClient struct {
	databaseURL  string
	databaseName string
}

type IMongoDBClient interface {
	InitConnection() (*mongo.Client, error)
	PingConnection(client *mongo.Client) error
	Disconnection(client *mongo.Client) error
	SetDatabase(client *mongo.Client, databaseName string) *mongo.Database
}

func NewMongoDB(databaseURL string, databaseName string) IMongoDBClient {
	return &mongoDBClient{
		databaseURL:  databaseURL,
		databaseName: databaseName,
	}
}

func (m *mongoDBClient) InitConnection() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf(`mongodb://%s`, m.databaseURL))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	return client, err
}

func (m *mongoDBClient) SetDatabase(client *mongo.Client, databaseName string) *mongo.Database {
	db := client.Database(databaseName)
	return db
}

func (m *mongoDBClient) PingConnection(client *mongo.Client) error {
	err := client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}
	fmt.Println("Success Ping to MongoDB!")

	return nil
}

func (m *mongoDBClient) Disconnection(client *mongo.Client) error {
	err := client.Disconnect(context.TODO())
	if err != nil {
		return err
	}
	fmt.Println("Success Disconnection to mongodb")
	return nil
}
