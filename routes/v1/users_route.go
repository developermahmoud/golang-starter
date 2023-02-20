package v1

import (
	"bm-support/src/controllers"

	"bm-support/src/middleware"

	"github.com/gin-gonic/gin"
)

var userController = controllers.NewUsersController()

func SetupUsersRoute(route *gin.RouterGroup) {
	users := route.Group("users")
	users.POST("/register", userController.Register)
	users.Use(middleware.IsAuth())
	{
		users.GET("/", userController.Index)
		users.GET("/:id", userController.GetByID)
	}
}
