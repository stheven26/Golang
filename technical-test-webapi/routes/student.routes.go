package routes

import (
	"technical-test-webapi/controllers"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func Routes(route *gin.Engine) {
	router.POST("/student", controllers.RegisterStudent)
	router.GET("/student/:id", controllers.GetStudent)
	router.PUT("/student/:id", controllers.UpdateStudent)
	router.DELETE("/student/:id", controllers.DeleteStudent)
}
