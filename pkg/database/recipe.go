package database

import "github.com/chrismalcolm/recipe/pkg/model"

func (db *Database) AddRecipe(recipe model.Recipe) (recipeID model.RecipeID, err error) {
	return recipeID, err
}

func (db *Database) EditRecipe(recipe model.Recipe) (err error) {
	return err
}

func (db *Database) DeleteRecipe(userID model.UserID) (err error) {
	return err
}

func (db *Database) SearchRecipes(userID model.UserID) (recipes []model.Recipe, err error) {
	return recipes, err
}
