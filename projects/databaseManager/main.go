package main

import (
	"fmt"
	"os"

	"github.com/HeartLess10/coffee-shop/coffee-shop/projects/databaseManager/dbManager/factory"
	"github.com/HeartLess10/coffee-shop/coffee-shop/projects/databaseManager/server"
	"github.com/HeartLess10/coffee-shop/coffee-shop/utils/customLogger"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("./config/.env")
	l := customLogger.NewPrettyCustomLogger("coffee-shop")
	server := server.NewServer(os.Getenv("PORT"), os.Getenv("DB_MANAGER_ADDRESS"), l)
	factory := factory.NewDbManagerFactory(l)
	if factory == nil {
		l.Error(fmt.Errorf("couldn't create database manager factory."))
	}
	dbManager := factory.CreateDbManager("mongo")
	if dbManager == nil {
		l.Error(fmt.Errorf("couldn't create database manager"))
	}
	dbManager.Connect()
	server.Serve()
}
