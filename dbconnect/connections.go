package connections

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://nilaygupta3003:<>@cluster0.xox98ed.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchList"

// making a connection now
var Collection *mongo.Collection

func init() {
	//client oprions
	clientOptions := options.Client().ApplyURI(connectionString)
	//connecting
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connectoion established")

	Collection = client.Database(dbName).Collection(colName)
	fmt.Print("collection instance is ready")
}
