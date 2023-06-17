package main

import (
	"authServer/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main(){
	// setting the port using env
	port := os.Getenv("PORT")

	// error catch if no port is defined so default fall back to
	if port==""{
		port = "8000"
	}
	// creating a router 
	router := gin.New()
	router.Use(gin.Logger())

	// pasing the router to both route
	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	// creating auth user route
	router.GET("/api-1", func(c *gin.Context)  {
		// Setting headers and response messages
		c.JSON(200, gin.H{"success":"Acess Granted for the auth-user"})
	})
	router.GET("/api-2", func(c *gin.Context)  {
		// Setting headers and response messages again
		c.JSON(200, gin.H{"success":"Acess Granted for the regular-user"})
	})

	// Starting the router
	router.Run(":" + port)
}
