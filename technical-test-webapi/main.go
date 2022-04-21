package main

import (
	"technical-test-webapi/routes"

	"technical-test-webapi/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.Routes(router)

	router.POST("/student", controllers.RegisterStudent)
	router.GET("/student/:id", controllers.GetStudent)
	router.PUT("/student/:id", controllers.UpdateStudent)
	router.DELETE("/student/:id", controllers.DeleteStudent)

	router.Run(":3000")
}
