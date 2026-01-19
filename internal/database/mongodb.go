package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.uber.org/zap"
)

type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func NewMongo(ctx context.Context, uri, dbName string, logger *zap.Logger) (*MongoDB, error) {

	clientOptions := options.Client().ApplyURI(uri).SetServerSelectionTimeout(10 * time.Second)
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		logger.Error("Failed to connect to MongoDB", zap.Error(err))
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Ping the database to verify connection
	if err := client.Ping(ctx, nil); err != nil {
		logger.Error("Failed to ping MongoDB", zap.Error(err))
		return nil, err
	}

	logger.Info("connected to mongodb", zap.String("uri", uri), zap.String("db", dbName))

	db := client.Database(dbName)

	return &MongoDB{Client: client, Database: db}, nil
}
