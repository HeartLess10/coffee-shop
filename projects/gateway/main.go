package main

import (
	"os"

	"github.com/HeartLess10/coffee-shop/coffee-shop/projects/gateway/server"
	"github.com/HeartLess10/coffee-shop/coffee-shop/utils/customLogger"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("./config/.env")
	l := customLogger.NewPrettyCustomLogger("coffee-shop")
	server := server.NewServer(os.Getenv("PORT"), os.Getenv("GATEWAY_ADDRESS"), l)
	server.Serve()
}
