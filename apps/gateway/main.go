package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	NewLogger("g")

	godotenv.Load("config/.env")
	port := os.Getenv("port")
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world")
	})

	if err := http.ListenAndServe("localhost:"+port, mux); err != nil {
		fmt.Println(err.Error())
	}

}
