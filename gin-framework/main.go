package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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

func initServer() (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("init: %w", err)

			return
		}
	}()

	data, err := os.ReadFile("./recipes.json")
	if err != nil {
		return err
	}

	if err = json.Unmarshal(data, &recipes); err != nil {
		return err
	}

	return err
}

func NewRecipeHandler(c *gin.Context) {
	recipe := Recipe{}

	if err := c.BindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	recipe.ID = xid.New().String()
	recipe.PusblisedAt = time.Now().Format(time.RFC3339)

	fmt.Printf("New recipe: %v\n", recipe)
}

func main() {
	err := initServer()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	router.POST("/recipes", NewRecipeHandler)

	router.Run(":3000")
}
