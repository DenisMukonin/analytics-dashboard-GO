package database

import (
	"fmt",
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB(mongoURI string) (*mongo.Client, *mongo.Database, error) {
	// Контекст с timeout

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Подключиться
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		fmt.Fatal("Failed to connect to MongoDB:", err)
		return nil, nil, err
	}

	// Проверить подключение
	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Fatal("Failed to ping MongoDB:", err)
		return nil, nil, err
	}

	fmt.Println("Connected to MongoDB")

	db := client.Database("analytics")

	// Создать индексы для оптимизации поиска
	createIndexes(db)

	return client, db, nil
}

func createIndexes(db *mongo.Database) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	indexModel := mongo.IndexModel{
		Keys: options.Index().SetKey("timestamp", "userId"),
	}
	db.Collection("events").Indexes().CreateOne(ctx, indexModel)

	indexModel = mongo.IndexModel{
		Keys: options.Index().SetKey("date", "metricName"),
	}
	db.Collection("metrics").Indexes().CreateOne(ctx, indexModel)

	fmt.Println("Indexes created")
}