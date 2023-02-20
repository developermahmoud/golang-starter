package response

import (
	"bm-support/src/utils/context"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ValidationErrorResponse struct {
	FailedField string `json:"field"`
	Rule        string `json:"rule"`
	Value       string `json:"value"`
}

func JSON(ctx context.Context, statusCode int, payload interface{}, message string) {
	ctx.Destroy()

	ctx.Ctx.JSON(statusCode, map[string]interface{}{
		"payload": payload,
		"message": message,
		"status":  true,
	})
}

func BindJsonError(ctx context.Context, err error) {
	ctx.Destroy()

	var errors []*ValidationErrorResponse

	if err.Error() == "EOF" {
		// Return response
		ctx.Ctx.JSON(http.StatusBadRequest, context.Payload{
			Payload: errors,
			Message: "invalid json object",
			Status:  false,
		})
		return
	}

	for _, value := range err.(validator.ValidationErrors) {
		var element ValidationErrorResponse
		element.FailedField = strings.ToLower(value.Field())
		element.Rule = value.Tag()
		element.Value = value.Param()
		errors = append(errors, &element)
	}

	// Return response
	ctx.Ctx.JSON(http.StatusBadRequest, context.Payload{
		Payload: errors,
		Message: "validation error",
		Status:  false,
	})
}

func WithError(ctx context.Context, err error) {
	ctx.Destroy()

	var statusCode int
	if err.Error() == "bad request" {
		statusCode = http.StatusBadRequest
	} else if err.Error() == "unauthorized" {
		statusCode = http.StatusUnauthorized
	} else if err.Error() == "not found" {
		statusCode = http.StatusNotFound
	}

	ctx.Ctx.JSON(statusCode, context.Payload{
		Payload: "",
		Message: err.Error(),
		Status:  false,
	})
	ctx.Ctx.Abort()
}
