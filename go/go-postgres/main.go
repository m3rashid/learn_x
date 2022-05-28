package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/m3rashid/learn_x/go/go-postgres/routers"
)

// name of the database => gopostgres

func main() {
	r := routers.Router()
	fmt.Println("Server is running on port 5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}
