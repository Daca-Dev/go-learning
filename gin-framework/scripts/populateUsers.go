package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("loading .env file: %v", err)
		return
	}

	users := map[string]string{
		"admin":      "fCRmh4Q2J7Rseqkz",
		"packt":      "RE4zfHB35VPtTkbT",
		"mlabouardy": "L3nSFRcZzNQ67bcc",
	}

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("users")

	for username, password := range users {
		h := sha256.New()
		h.Write([]byte(password))
		hashedPassword := hex.EncodeToString(h.Sum(nil))

		collection.InsertOne(ctx, bson.M{
			"username": username,
			"password": hashedPassword,
		})
	}

	log.Println("Users inserted successfully!")
}
