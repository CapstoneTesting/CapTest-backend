package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type YourStruct struct {
	Field1 string `json:"field1"`
	Field2 int    `json:"field2"`
}

func main() {
	//Initiate an engine instance.
	router := gin.Default()

	//Logger to track request method, status and request duration. Use() attaches global middleware to router.
	router.Use(LoggerMiddleware())

	router.GET("/", func(c *gin.Context) {
		fmt.Println(c, "its here")
		c.String(200, "Hello world") //Writes given string into response body
	})

	router.POST("/test", func(c *gin.Context) {
		// Parse the JSON request body into a struct
		var data YourStruct // Replace YourStruct with your data structure

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("HERE")
		// You can now access the data from the request body in the 'data' variable
		// Do something with the data, e.g., store it in a database
		// Respond with a success message
		c.JSON(http.StatusOK, gin.H{"message": "Data received successfully", "data": data})
	})

	router.Run(":8081") //Serve and listen for localhost port 8080, attaches router to http.Server.
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Std time. Returns current local time.
		start := time.Now()
		//Context.Next; only used in middleware, executes pending handlers in the chain inside the calling handler.
		c.Next()
		//Std time. Returns time elapsed since start.
		duration := time.Since(start)
		//%s, %d are Context types that have exxcess to request and writer types
		log.Printf("\nRequest - Method: %s | Status: %d | Duration: %v \n", c.Request.Method, c.Writer.Status(), duration)
	}
}
