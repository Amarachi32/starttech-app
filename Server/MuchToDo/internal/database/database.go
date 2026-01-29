package database

import (
	"context"
	"time"
	"strings"
	"log/slog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// ConnectMongo establishes a connection to MongoDB using the provided URI.
func ConnectMongo(uri, dbName string) (*mongo.Client, error) {
	// Set a timeout for the connection attempt.
	uri = strings.TrimSpace(uri)
    uri = strings.Trim(uri, "\"") 
    uri = strings.Trim(uri, "'")

    slog.Debug("Connecting to MongoDB", "uri_preview", uri[:15]+"...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// Ping the primary node to verify that the connection is alive.
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return client, nil
}
