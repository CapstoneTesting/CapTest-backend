package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	//Initiate an engine instance.
	router := gin.Default()

	//Logger to track request method, status and request duration. Use() attaches global middleware to router.
	router.Use(LoggerMiddleware())

	router.GET("/", func(c *gin.Context) {
		fmt.Println(c, "its here")
		c.String(200, "Hello world") //Writes given string into response body
	})

	router.Run(":8080") //Serve and listen for localhost port 8080, attaches router to http.Server.
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
