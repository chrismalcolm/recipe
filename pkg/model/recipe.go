package model

type RecipeID uint64

type Recipe struct {
	ID          RecipeID
	Creator     User
	Ingredients Ingredients
}

type Ingredients struct {
	Item     Item
	Quantity Quantity
	Unit     Unit
}

type Item struct {
	Name string
}

type Quantity struct {
	Magnitude Fraction
}

type Fraction struct {
	Numerator   int
	Denominator int
}

type Unit struct {
	Name      string
	Magnitude Fraction
	Category  Category
}

type Category struct {
	Name string
}

func ParseRecipesFromPlainText(plainText string) (recipes []Recipe) {
	return recipes
}

func Combine(recipes []Recipe) (recipe Recipe) {
	return recipe
}
