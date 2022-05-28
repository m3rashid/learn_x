package main

import (
	"fmt"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJwt()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, validToken)
}

func HandleRoutes() {
	http.HandleFunc("/", Home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
