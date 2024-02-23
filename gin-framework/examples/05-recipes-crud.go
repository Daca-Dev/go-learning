package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

type Recipe struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Tags         []string `json:"tags"`
	Ingredients  []string `json:"ingredients"`
	Instructions []string `json:"instructions"`
	PusblisedAt  string   `json:"publishedAt"`
}

var recipes []Recipe

func initServer() {
	recipes = make([]Recipe, 0)
	file, _ := os.ReadFile("recipes.json")
	json.Unmarshal(file, &recipes)
}

func NewRecipeHandler(c *gin.Context) {
	var recipe Recipe

	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	recipe.ID = xid.New().String()
	recipe.PusblisedAt = time.Now().Format(time.RFC3339)

	recipes = append(recipes, recipe)

	c.JSON(http.StatusCreated, recipe)
}

func ListRecipesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, recipes)
}

func GetRecipesHandler(c *gin.Context) {
	id := c.Param("id")

	for _, recipe := range recipes {
		if recipe.ID == id {
			c.JSON(http.StatusOK, recipe)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "recipe not found"})
}

func UpdateRecipeHandler(c *gin.Context) {
	id := c.Param("id")

	index := -1

	for i, recipe := range recipes {
		if recipe.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "recipe not found"})
		return
	}

	var recipe Recipe
	c.ShouldBindJSON(&recipe)

	recipes[index] = recipe

	c.JSON(http.StatusOK, recipe)
}

func DeleteRecipeHandler(c *gin.Context) {
	id := c.Param("id")

	index := -1

	for i, recipe := range recipes {
		if recipe.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "recipe not found"})
		return
	}

	recipes = append(recipes[:index], recipes[index+1:]...)

	c.JSON(http.StatusNoContent, nil)
}

func SearchRecipesHandler(c *gin.Context) {
	tag := c.Query("tag")
	listOfRecipes := make([]Recipe, 0)

	for i := 0; i < len(recipes); i++ {
		found := false
		for _, t := range recipes[i].Tags {
			if strings.EqualFold(t, tag) {
				found = true
				break
			}
		}

		if found {
			listOfRecipes = append(listOfRecipes, recipes[i])
		}
	}

	c.JSON(http.StatusOK, listOfRecipes)
}

func main() {
	initServer()

	router := gin.Default()

	router.POST("/recipes", NewRecipeHandler)
	router.GET("/recipes", ListRecipesHandler)
	router.GET("/recipes/:id", GetRecipesHandler)
	router.PUT("/recipes/:id", UpdateRecipeHandler)
	router.DELETE("/recipes/:id", DeleteRecipeHandler)
	router.GET("/recipes/search", SearchRecipesHandler)

	router.Run(":3000")
}
