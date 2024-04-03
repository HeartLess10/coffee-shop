package databaseHandlers

import (
	"fmt"
	"net/http"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}
