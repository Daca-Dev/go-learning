package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Recipe struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Tags         []string `json:"tags"`
	Ingredients  []string `json:"ingredients"`
	Instructions []string `json:"instructions"`
	PusblisedAt  string   `json:"publishedAt"`
}

var ctx context.Context
var err error
var client *mongo.Client
var recipes []Recipe

func loadRecipes() {
	recipes = make([]Recipe, 0)
	file, _ := os.ReadFile("recipes.json")
	json.Unmarshal(file, &recipes)
}

func main() {
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("loading .env file: %v", err)
		return
	}

	ctx = context.Background()
	co := options.Client().ApplyURI(os.Getenv("MONGO_URI"))

	client, err = mongo.Connect(ctx, co)
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatalf("ping to mongo: %v", err)
	}
	defer client.Disconnect(ctx)

	log.Println("Connected to MongoDB!")

	var listOfRecipes []interface{}
	loadRecipes()

	for _, recipe := range recipes {
		listOfRecipes = append(listOfRecipes, recipe)
	}

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("recipes")
	insertManyResult, err := collection.InsertMany(
		ctx, listOfRecipes)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Inserted recipes: ", len(insertManyResult.InsertedIDs))
}
