package repositories

type CategoriesRepository interface {
}

type categoriesRepository struct{}

func NewCategoriesRepository() CategoriesRepository {
	return &categoriesRepository{}
}

