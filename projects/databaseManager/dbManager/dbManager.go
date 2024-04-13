package dbManager

import (
	"context"
	"fmt"
	"os"

	"github.com/HeartLess10/coffee-shop/coffee-shop/projects/databaseManager/dbManager/db"
	"github.com/HeartLess10/coffee-shop/coffee-shop/utils/customLogger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	Mongo = "mongo"
)

// BUG:(lidor) save in dbmanager get base key so i wont need to type the key everywhere
type DbManager struct {
	connectionStatus bool
	dbType           string
	c                *mongo.Client
	databaseName     string
	l                customLogger.CustomLogger
}

func NewDbManager(dbType string, databaseName string, l customLogger.CustomLogger) *DbManager {
	if len(databaseName) == 0 || l == nil {
		panic("invalid parameters given databaseName cant be empty and l cant be nil")
	}

	//TODO:(lidor) Think about checking here db type is valid

	return &DbManager{dbType: dbType, connectionStatus: false, databaseName: databaseName, l: l}
}

func (dbm *DbManager) Connect(ctx context.Context, l customLogger.CustomLogger) error {
	var err error = nil
	switch dbm.dbType {
	case Mongo:
		dbUri := createMongoUri()
		dbm.c, err = mongo.Connect(ctx, options.Client().ApplyURI(dbUri))
		dbm.connectionStatus = true
	default:
		err = fmt.Errorf("trying to connect to an invalid database type %s", dbm.dbType)
	}
	return err
}

func (dbm *DbManager) Disconnect(ctx context.Context) error {
	switch dbm.dbType {
	case Mongo:
		return dbm.c.Disconnect(ctx)

	default:
		return fmt.Errorf("database type is not supported: %s", dbm.dbType)
	}
}

func GetDbCollection[T any](dbm *DbManager) db.Db[T] {
	if dbm == nil {
		panic("db manager cant be nil when trying to receive a db collection")
	}
	if !dbm.connectionStatus {
		panic("db manager is not connected to the db")
	}
	if dbm.l == nil {
		panic("db manager has to have a custom logger")
	}
	switch dbm.dbType {
	case Mongo:
		//TODO:(lidor) Connect to db if not connected + make sure its not locking the factory
		if dbm.c != nil {
			return db.NewMongoDb[T](dbm.databaseName, *dbm.c, dbm.l)
		}
		return nil
	default:
		return nil
	}
}

func createMongoUri() string {
	credentials := fmt.Sprintf("%s:%s", os.Getenv("MONGO_DB_USERNAME"), os.Getenv("MONGO_DB_PASSWORD"))
	hostname := fmt.Sprintf("%s:%s", os.Getenv("MONGO_DB_MANAGER_ADDRESS"), os.Getenv("MONGO_DB_PORT"))
	protocol := "mongodb"
	connectionOptions := ""
	dbUri := fmt.Sprintf("%s://%s@%s/%s", protocol, credentials, hostname, connectionOptions)
	return dbUri
}
