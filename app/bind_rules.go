package app

import (
	"bm-support/config/database"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var unique validator.Func = func(fl validator.FieldLevel) bool {
	// Get Value Of Field
	value := fl.Field()
	// Get Field Name
	fieldName := strings.ToLower(fl.FieldName())
	// Get Table Name
	tableName := fl.Param()
	// Assign Result To
	result := map[string]interface{}{}
	// Check In DB
	record := database.DB.Table(tableName).Where(fieldName+" = ?", value).Take(&result)
	if record.Error != nil {
		return true
	} else {
		return false
	}
}

func bindCustomRules() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("unique", unique)
	}
}
