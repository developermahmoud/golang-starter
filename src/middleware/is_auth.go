package middleware

import (
	"bm-support/src/repositories"
	"bm-support/src/utils/context"
	"bm-support/src/utils/response"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func IsAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		context := context.NewContext(c)

		// Get authorization token
		authorization := context.Ctx.Request.Header["Authorization"]
		if len(authorization) == 0 {
			err := errors.New("unauthorized")
			response.WithError(context, err)
			return
		}

		token := strings.Split(authorization[0], " ")

		if len(token) == 0 {
			err := errors.New("unauthorized")
			response.WithError(context, err)
			return
		}

		// Init repository
		tokenRepository := repositories.NewTokenRepository()

		// Get token from DB
		tokenObj, err := tokenRepository.GetByToken(token[1])
		if err != nil {
			response.WithError(context, err)
			return
		}

		// Check Token
		registerToken, err := jwt.ParseWithClaims(token[1], &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("APP_SECRET")), nil
		})

		if err != nil || !registerToken.Valid {
			// Delete token
			tokenRepository.Delete(tokenObj.UserID)

			err := errors.New("unauthorized")
			response.WithError(context, err)
			return
		}

		// Store UserID in header
		context.Ctx.Request.Header.Set("user", fmt.Sprint(tokenObj.UserID))

		// Close connection
		context.Destroy()

		context.Ctx.Next()
	}
}
