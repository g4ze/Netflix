package helper

import (
	"connections"
	"context"
	"fmt"
	"log"
	"model"
)

//MONGOdb helpers
func insertOneMovie(movie model.Movies) {
	inserted, err := connections.Collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Insterted 1 movie :", inserted.InsertedID)
}
