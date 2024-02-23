package handlers

import (
	"context"
	"gin-framework/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RecipesHandler struct {
	collection *mongo.Collection
	ctx        *context.Context
}

func (handler *RecipesHandler) ListRecipes(c *gin.Context) {
	// bson = binary json
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

	c.JSON(http.StatusOK, recipes)
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

func NewRecipeHandler(ctx *context.Context, collection *mongo.Collection) *RecipesHandler {
	return &RecipesHandler{collection, ctx}
}
