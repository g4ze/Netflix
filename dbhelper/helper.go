package helper

import (
	"connections"
	"context"
	"fmt"
	"log"
	"model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func init() {
	collection = (connections.Collection)
}

//MONGOdb helpers
func InsertOneMovie(movie model.Movies) {

	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Insterted 1 movie :", inserted.InsertedID)
}

//update one record
func UpdateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, _ := collection.UpdateOne(context.Background(), filter, update)
	fmt.Println("Value updated of movie:", collection.FindOne(context.Background(), bson.M{"_id": id}))
	fmt.Println("modified counr:", result.ModifiedCount)
}

//delete one record
func DeleteOneMovie(moviesId string) {
	id, _ := primitive.ObjectIDFromHex((moviesId))
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		fmt.Println("Error in helper block")
	}
	fmt.Println("Movie gor delete with delete count: ", deleteCount)
}

//get all movies
func GetAllMovies() []primitive.M {
	cur, _ := collection.Find(context.Background(), bson.D{{}})
	var movies []primitive.M
	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cur.Close(context.Background())
	return movies
}
