package meals

import (
	"testing"
	"time"
)

func TestOne(t *testing.T) {
	i1 := NewIngredient("Falukorv", []IngredientAttr{Sausage, Meat})
	i2 := NewIngredient("Makaroner", []IngredientAttr{Pasta, Side})
	d := NewDish("Falukorv med Makaroner", time.Minute*15, i1, i2)
	assert(t, d != nil, "Could not create new dish")
	assert(t, d.Name == "Falukorv med Makaroner", "Dish name incorrect")
	assert(t, len(d.Ingrs) == 2, "Wrong number of ingredients")
}
