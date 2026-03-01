package db

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func DBinstance() *mongo.Client{
	_ , err := gotdotenv.Load(".env")
	if err!= nil{
		log.Fatal("errro loading .env file"))
	}

	MongoDb:= os.Getenv("MONGODB_URL")

	client, err := mongo.NewClient(options.Client().applyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}	
	fmt.Println("Connected to MongoDB!")
	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionname string) *mongo.Collection{
	var collection *mongo.Collection = client.Database("golang").Collection(collectionname)
	return collection
} 