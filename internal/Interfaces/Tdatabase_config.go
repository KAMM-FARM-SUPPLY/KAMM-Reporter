package interfaces

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type DatabaseConfig interface {
	Close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc)
	Connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error)
	Ping(client *mongo.Client, ctx context.Context) error
}
