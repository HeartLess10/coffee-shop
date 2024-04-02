package main

import (
	"os"

	"github.com/HeartLess10/coffee-shop/coffee-shop/projects/gateway/server"
	"github.com/HeartLess10/coffee-shop/coffee-shop/utils/customLogger"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("./config/.env")
	port := os.Getenv("PORT")
	l := customLogger.NewPrettyCustomLogger("coffee-shop")
	server := server.NewServer(port, "localhost", l)
	server.Serve()

}
