package db

import (
	"context"
	"reflect"
)

type Db[T any] interface {
	InsertOne(ctx context.Context, item T)
	ReadById(ctx context.Context, id int) T
	ReadByQuery(ctx context.Context, query string) []T
	Delete(ctx context.Context, id int)
	Update(ctx context.Context, id int, item T)
}

func getCollectionName[T any]() string {
	var v1 reflect.Type = reflect.TypeOf((*T)(nil)).Elem()
	return v1.Name()
}
