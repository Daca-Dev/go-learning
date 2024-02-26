package main

import (
	"context"
	"gin-framework/handlers"
	"gin-framework/middlewares"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var recipeHandler *handlers.RecipesHandler
var authHandler *handlers.AuthHandler

func initServer() {
	err := godotenv.Load(".env")
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

	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "",
		DB:       0,
	})
	status := redisClient.Ping(ctx)

	log.Println("Connected to Redis!", status)

	recipeHandler = handlers.NewRecipeHandler(
		&ctx,
		collection,
		redisClient,
	)

	authHandler = &handlers.AuthHandler{}
}

func main() {
	initServer()

	router := gin.Default()
	privateRecipesRoutes := router.Group("/recipes")

	// privateRecipesRoutes.Use(middlewares.APIKeyAuth())
	privateRecipesRoutes.Use(middlewares.JWTAUth())
	privateRecipesRoutes.POST("/", recipeHandler.CreateRecipe)
	privateRecipesRoutes.PUT("/:id", recipeHandler.UpdateRecipe)
	privateRecipesRoutes.DELETE("/:id", recipeHandler.DeleteRecipe)

	router.POST("/signin", authHandler.SignHandler)
	router.GET("/recipes", recipeHandler.ListRecipes)
	router.GET("/recipes/:id", recipeHandler.GetRecipe)
	// router.GET("/recipes/search", SearchRecipesHandler)

	router.Run(":3000")
}
