package main

import (
	"fmt"
	"log"
	"router"
)

func main() {
	fmt.Print("Netflix Movies MongoDB API\n")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.listenAndServe(":4000", r))
}
