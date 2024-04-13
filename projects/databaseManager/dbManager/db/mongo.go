package db

import (
	"context"
	"fmt"

	"github.com/HeartLess10/coffee-shop/coffee-shop/utils/customLogger"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoDb[T any] struct {
	l              customLogger.CustomLogger
	c              mongo.Client
	databaseName   string
	collectionName string
}

// Insert implements DbManager.
func (m *mongoDb[T]) InsertOne(ctx context.Context, item T) {
	m.l.Info(fmt.Sprintf("Inserting into database: %s, collection: %s, item:  %+v ", m.databaseName, m.collectionName, item))
	m.c.Database(m.databaseName).Collection(m.collectionName).InsertOne(ctx, item)
}

// Delete implements DbManager.
func (m *mongoDb[T]) Delete(ctx context.Context, id int) {
	panic("unimplemented")
}

// ReadById implements DbManager.
func (m *mongoDb[T]) ReadById(ctx context.Context, id int) T {
	panic("unimplemented")
}

// ReadByQuery implements DbManager.
func (m *mongoDb[T]) ReadByQuery(ctx context.Context, query string) []T {
	panic("unimplemented")
}

// Update implements DbManager.
func (m *mongoDb[T]) Update(ctx context.Context, id int, item T) {
	panic("unimplemented")
}

func NewMongoDb[T any](databaseName string, c mongo.Client, l customLogger.CustomLogger) Db[T] {
	return &mongoDb[T]{databaseName: databaseName, c: c, l: l, collectionName: getCollectionName[T]()}
}
