package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tomarbhavuk/mongoapi/router"
)

func main() {
	fmt.Println("Welcome to mongo tutorial")
	fmt.Println("Server is getting started...")

	r := router.Router()
	http.ListenAndServe(":8080", r)

	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("We are listning at 8080")
}
