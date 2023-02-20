package controllers

import (
	"bm-support/src/dto"
	"bm-support/src/repositories"
	"bm-support/src/utils/context"
	"bm-support/src/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	userRepository repositories.UserRepository
}

func NewUsersController() UsersController {
	return UsersController{
		userRepository: repositories.NewUserRepository(),
	}
}

func (controller UsersController) Register(c *gin.Context) {
	context := context.NewContext(c)
	var dto dto.UserRegisterDTO

	if err := context.Ctx.ShouldBindJSON(&dto); err != nil {
		response.BindJsonError(context, err)
		return
	}

	user, _ := controller.userRepository.Store(dto)

	response.JSON(context, http.StatusOK, user, "create new user")
}

func (controller UsersController) GetByID(c *gin.Context) {
	context := context.NewContext(c)
	id, _ := strconv.ParseUint(context.Ctx.Param("id"), 10, 64)

	user, err := controller.userRepository.GetByID(id)
	if err != nil {
		response.BindJsonError(context, err)
		return
	}

	response.JSON(context, http.StatusOK, user, "get user by id")
}

func (controller UsersController) Index(c *gin.Context) {
	context := context.NewContext(c)

	user, _ := controller.userRepository.Index()

	response.JSON(context, http.StatusOK, user, "show user")
}
