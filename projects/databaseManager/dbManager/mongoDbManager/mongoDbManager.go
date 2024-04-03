package mongoDbManager

import (
	"context"

	"github.com/HeartLess10/coffee-shop/coffee-shop/projects/databaseManager/dbManager"
	"github.com/HeartLess10/coffee-shop/coffee-shop/utils/customLogger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDbManager struct {
	addr   string
	l      customLogger.CustomLogger
	client *mongo.Client
}

func (m *mongoDbManager) Connect() {
	m.l.Message("Connecting to mongodb server.")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(m.addr))
	if err != nil {
		m.l.Error(err)
	}
	m.client = client
}

func (m *mongoDbManager) Create() {
	panic("unimplemented")
}

func (m *mongoDbManager) Delete() {
	panic("unimplemented")
}

func (m *mongoDbManager) Read() {
	panic("unimplemented")
}

func (m *mongoDbManager) Update() {
	panic("unimplemented")
}
func NewMongoDbManager(addr string, l customLogger.CustomLogger) dbManager.DbManager {
	return &mongoDbManager{addr: addr, l: l, client: nil}
}
