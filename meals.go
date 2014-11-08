package meals

import "time"

var (
	ingr []*Ingredient
)

const (
	Sausage IngredientAttr = iota
	Meat
	Pasta
	Side
	Spice
	Sauce

	Soup DishAttr = iota
	Desert
	Main
)

type IngredientAttr int

type DishAttr int

type Ingredient struct {
	Name  string
	Attrs []IngredientAttr
}

type Dish struct {
	T     DishAttr
	Name  string
	Time  time.Duration
	Desc  string
	Ingrs []*Ingredient
}

func NewIngredient(name string, attr ...[]IngredientAttr) *Ingredient {
	return nil
}

func NewDish(name string, t time.Duration, i ...*Ingredient) *Dish {
	return &Dish{Name: name, Time: t, Ingrs: i}
}
