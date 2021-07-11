package service

import "github.com/chrismalcolm/recipe/pkg/model"

type RecipeAdmin interface {
	AddRecipe(recipe model.Recipe) (recipeID model.RecipeID, err error)
	EditRecipe(recipe model.Recipe) (err error)
	DeleteRecipe(recipeID model.RecipeID) (err error)
	SearchRecipes(userID model.UserID) (recipes []model.Recipe, err error)
}

type RecipeService struct {
	ua RecipeAdmin
}

func NewRecipeService(recipeAdmin RecipeAdmin) (us *RecipeService) {
	return &RecipeService{
		ua: recipeAdmin,
	}
}
