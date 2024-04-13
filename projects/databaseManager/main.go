package main

import (
	"context"
	"os"

	"github.com/HeartLess10/coffee-shop/coffee-shop/projects/databaseManager/dbManager"
	"github.com/HeartLess10/coffee-shop/coffee-shop/projects/databaseManager/server"
	"github.com/HeartLess10/coffee-shop/coffee-shop/utils/customLogger"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("./config/.env")
	l := customLogger.NewPrettyCustomLogger("coffee-shop")

	dbm := dbManager.NewDbManager(dbManager.Mongo, os.Getenv("DATABASE_NAME"), l)
	err := dbm.Connect(context.Background(), l)
	if err != nil {
		panic(err)
	}
	server := server.NewServer(os.Getenv("PORT"), os.Getenv("DB_MANAGER_ADDRESS"), dbm, l)

	defer dbm.Disconnect(context.Background())
	server.Serve()
}
