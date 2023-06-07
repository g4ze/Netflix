package main

import (
	"fmt"
	"log"
	"net/http"
	"router"
)

func main() {
	fmt.Print("Netflix Movies MongoDB API\n")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Running at 4000")

}
