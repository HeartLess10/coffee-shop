package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("./config/.env")
	port := os.Getenv("PORT")
	mydir, _ := os.Getwd()
	fmt.Print("hhhh:" + mydir)
	fmt.Print("hhhh:" + port)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world")
	})

	if err := http.ListenAndServe("localhost:"+port, mux); err != nil {
		fmt.Println(err.Error())
	}

}
