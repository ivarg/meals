package meals

import (
	"fmt"
	"math/rand"
	"time"
)

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
	Potato

	Soup DishAttr = iota
	Desert
	Main

	Bad = iota * .25
	Maybe
	Good
	Favorite
	Unrated = -1
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

func (d *Dish) String() string {
	return d.Name
}

type Meal struct {
	dish *Dish
	day  time.Time
}

func (m Meal) String() string {
	return fmt.Sprintf("%s@%s", m.dish, m.day.Format("2006-01-02"))
}

type Pref struct {
	dish *Dish
	p    float64
}

func MakeMeal(d *Dish, t time.Time) Meal {
	return Meal{dish: d, day: t}
}

func MakePref(d *Dish, p float64) Pref {
	return Pref{dish: d, p: p}
}

func inHistory(d *Dish, h []Meal) bool {
	for _, m := range h {
		if m.dish == d {
			return true
		}
	}
	return false
}

type PrefDishes []Pref

func (d PrefDishes) Len() int {
	return len(d)
}

func (d PrefDishes) Less(i, j int) bool {
	return d[i].p < d[j].p
}

func (d PrefDishes) Swap(i, j int) {
	tmp := d[i]
	d[i] = d[j]
	d[j] = tmp
}

func unrated(d *Dish, ps []Pref) bool {
	for _, p := range ps {
		if p.dish == d {
			return true
		}
	}
	return false
}

// For each request, create a sorted list of dishes as follows:
// - Last periods meals last (sorted according to Pref)
// - Dishes with prefs are sorted according to pref combined with time-since-last
// - The list is then locally randomized with overlap, so that there is a non-zero
//   probability that a very low rated dish might end up in the selection.
func newSelection(dishes []*Dish, n int, hist []Meal) []*Dish {
	var src []*Dish
	for _, d := range dishes {
		if !inHistory(d, hist) {
			src = append(src, d)
		}
	}

	if n > len(src) {
		n = len(src)
	}

	var psrc []*Dish
	for _, i := range rand.Perm(len(src)) {
		psrc = append(psrc, src[i])
	}

	return psrc[:n]
}

func NewIngredient(name string, attr ...IngredientAttr) (*Ingredient, error) {
	for _, i := range ingr {
		if i.Name == name {
			return nil, fmt.Errorf("Ingredient already exists")
		}
	}
	i := &Ingredient{Name: name, Attrs: attr}
	ingr = append(ingr, i)
	return i, nil
}

func NewDish(name string, t time.Duration, i ...*Ingredient) *Dish {
	return &Dish{Name: name, Time: t, Ingrs: i}
}

// Lookup and return the ingredient with the given name, and nil otherwise.
func I(name string) *Ingredient {
	for _, i := range ingr {
		if i.Name == name {
			return i
		}
	}
	i, _ := NewIngredient(name)
	return i
}

func FindDish(ds []*Dish, name string) *Dish {
	for _, d := range ds {
		if d.Name == name {
			return d
		}
	}
	return nil
}

func D(name string, t int, ingr ...string) *Dish {
	var ii []*Ingredient
	for _, i := range ingr {
		ii = append(ii, I(i))
	}
	return NewDish(name, time.Minute*time.Duration(t), ii...)
}

type BasicComparator struct{}

// Simply compare the ingredients. The more ingredients in common, the higher
// the score.
func (c *BasicComparator) Similarity(d1 *Dish, d2 *Dish) float64 {
	var score = 0.0
	for _, i := range d1.Ingrs {
	Outer:
		for _, j := range d2.Ingrs {
			if i == j {
				score = score + 1
				break Outer
			}
		}
	}
	n := float64(len(d1.Ingrs)+len(d2.Ingrs)) / 2
	return (score / n)
}
