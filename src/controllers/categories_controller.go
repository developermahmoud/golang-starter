package controllers

import (
	"bm-support/src/repositories"
)

type CategoriesController struct {
	categoriesRepository repositories.CategoriesRepository
}

func NewCategoriesController() CategoriesController {
	return CategoriesController{
		categoriesRepository: repositories.NewCategoriesRepository(),
	}
}