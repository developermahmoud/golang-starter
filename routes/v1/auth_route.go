package v1

import (
	"bm-support/src/controllers"
	"github.com/gin-gonic/gin"
)

var authController = controllers.NewAuthController()

func SetupAuthRoute(route *gin.RouterGroup) {
	users := route.Group("auth")
	users.POST("/login", authController.Login)
}
