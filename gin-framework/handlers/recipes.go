package handlers

import (
	"context"
	"encoding/json"
	"gin-framework/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RecipesHandler struct {
	collection  *mongo.Collection
	ctx         *context.Context
	redisClient *redis.Client
}

func (handler *RecipesHandler) ListRecipes(c *gin.Context) {
	val, err := handler.redisClient.Get(*handler.ctx, "recipes").Result()
	if err == redis.Nil {
		log.Println("Requesting to MongoDB")

		cur, err := handler.collection.Find(*handler.ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer cur.Close(*handler.ctx)

		recipes := make([]models.Recipe, 0)
		for cur.Next(*handler.ctx) {
			var recipe models.Recipe
			cur.Decode(&recipe)

			recipes = append(recipes, recipe)
		}

		data, _ := json.Marshal(recipes)
		handler.redisClient.Set(*handler.ctx, "recipes", string(data), time.Minute*10)

		c.JSON(http.StatusOK, recipes)
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		log.Println("Requesting to Redis")

		recipes := make([]models.Recipe, 0)

		json.Unmarshal([]byte(val), &recipes)
		c.JSON(http.StatusOK, recipes)
	}

}

func (handler *RecipesHandler) GetRecipe(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var recipe models.Recipe
	err = handler.collection.FindOne(*handler.ctx, bson.M{"_id": objID}).Decode(&recipe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, recipe)
}

func (handler *RecipesHandler) CreateRecipe(c *gin.Context) {
	var recipe models.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipe.ID = primitive.NewObjectID()
	recipe.PusblisedAt = time.Now()

	_, err := handler.collection.InsertOne(*handler.ctx, recipe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, recipe)
}

func (handler *RecipesHandler) UpdateRecipe(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var recipe models.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = handler.collection.UpdateOne(*handler.ctx, bson.M{"_id": objID}, bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{Key: "name", Value: recipe.Name},
				{Key: "instructions", Value: recipe.Instructions},
				{Key: "ingredients", Value: recipe.Ingredients},
				{Key: "tags", Value: recipe.Tags},
				{Key: "publishedAt", Value: time.Now()},
			},
		}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, recipe)
}

func (handler *RecipesHandler) DeleteRecipe(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = handler.collection.DeleteOne(*handler.ctx, bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func NewRecipeHandler(ctx *context.Context, collection *mongo.Collection, redisClient *redis.Client) *RecipesHandler {
	return &RecipesHandler{collection, ctx, redisClient}
}
