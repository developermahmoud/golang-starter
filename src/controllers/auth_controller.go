package controllers

import (
	dto "bm-support/src/dto"
	"bm-support/src/repositories"
	"bm-support/src/utils/context"
	"bm-support/src/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userRepository  repositories.UserRepository
	tokenRepository repositories.TokenRepository
}

func NewAuthController() AuthController {
	return AuthController{
		userRepository:  repositories.NewUserRepository(),
		tokenRepository: repositories.NewTokenRepository(),
	}
}

func (controller AuthController) Login(c *gin.Context) {
	context := context.NewContext(c)
	var dto dto.LoginDTO

	// Validate body object
	if err := context.Ctx.ShouldBindJSON(&dto); err != nil {
		response.BindJsonError(context, err)
		return
	}

	// User login
	token, userID, err := controller.userRepository.Login(dto)
	if err != nil {
		response.WithError(context, err)
		return
	}

	// Delete last tokens
	if err := controller.tokenRepository.Delete(userID); err != nil {
		response.BindJsonError(context, err)
		return
	}

	// Store token
	if err := controller.tokenRepository.Store(token, userID); err != nil {
		response.BindJsonError(context, err)
		return
	}

	response.JSON(context, http.StatusOK, map[string]string{
		"token": token,
	}, "user loged in")
}
