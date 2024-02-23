package main

import (
	"context"
	"gin-framework/handlers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var err error

func initServer() (*context.Context, *mongo.Collection) {
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("loading .env file: %v", err)
	}

	ctx := context.Background()
	co := options.Client().ApplyURI(os.Getenv("MONGO_URI"))

	client, _ := mongo.Connect(ctx, co)
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatalf("ping to mongo: %v", err)
	}

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("recipes")

	log.Println("Connected to MongoDB!")

	return &ctx, collection
}

func main() {
	ctx, collection := initServer()
	rh := handlers.NewRecipeHandler(ctx, collection)

	router := gin.Default()

	router.GET("/recipes", rh.ListRecipes)
	router.GET("/recipes/:id", rh.GetRecipe)
	router.POST("/recipes", rh.CreateRecipe)
	router.PUT("/recipes/:id", rh.UpdateRecipe)
	router.DELETE("/recipes/:id", rh.DeleteRecipe)
	// router.GET("/recipes/search", SearchRecipesHandler)

	router.Run(":3000")
}
