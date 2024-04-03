package factory

import (
	"fmt"
	"os"

	"github.com/HeartLess10/coffee-shop/coffee-shop/projects/databaseManager/dbManager"
	"github.com/HeartLess10/coffee-shop/coffee-shop/projects/databaseManager/dbManager/mongoDbManager"
	"github.com/HeartLess10/coffee-shop/coffee-shop/utils/customLogger"
)

type DbManagerFactory interface {
	CreateDbManager(dbType string) dbManager.DbManager
	createMongoManager() dbManager.DbManager
}

type dbManagerFactory struct {
	l customLogger.CustomLogger
}

const ( // iota is reset to 0
	mongo = "mongo" // c0 == 0
)

func (dbmf *dbManagerFactory) CreateDbManager(dbType string) dbManager.DbManager {

	switch dbType {
	case mongo:
		return dbmf.createMongoManager()
	default:
		dbmf.l.Error(fmt.Errorf("database manager received an invalid database Type: %s", dbType))
		return nil
	}
}

func (dbmf *dbManagerFactory) createMongoManager() dbManager.DbManager {
	credentials := fmt.Sprintf("%s:%s", os.Getenv("MONGO_DB_USERNAME"), os.Getenv("MONGO_DB_PASSWORD"))
	hostname := fmt.Sprintf("%s:%s", os.Getenv("MONGO_DB_MANAGER_ADDRESS"), os.Getenv("MONGO_DB_PORT"))
	protocol := "mongodb"
	connectionOptions := ""
	dbAddr := fmt.Sprintf("%s://%s@%s/%s", protocol, credentials, hostname, connectionOptions)
	return mongoDbManager.NewMongoDbManager(dbAddr, dbmf.l)

}

func NewDbManagerFactory(l customLogger.CustomLogger) DbManagerFactory {
	l.Message("Creating Database manager factory.")
	return &dbManagerFactory{l: l}
}
