package main

import (
	"encoding/xml"

	"github.com/gin-gonic/gin"
)

type Person struct {
	XMLName   xml.Name `xml:"person"`
	FirstName string   `xml:"firstName,attr"`
	LastName  string   `xml:"lastName,attr"`
}

func IndexHandler(c *gin.Context) {
	name := c.Params.ByName("name")

	c.XML(200, Person{FirstName: name, LastName: "Doe"})
}

func main() {
	router := gin.Default()
	router.GET("/:name", IndexHandler)

	// router.Run()
	router.Run(":3000")
}
