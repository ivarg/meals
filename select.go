package meals

import (
	"sort"
	"time"
)

// What is given is a set of rated dishes and a history of past meals.
// Each rated dish gets its score calculated by combining its rate with
// its history value, i.e a function of the number of days since its last
// meal.

type Household struct {
	ratings map[*Dish]float64
	hist    []Meal
	dishes  []*Dish
}

func (h *Household) getRating(d *Dish) float64 {
	return 0
}

func (h *Household) latestMeal(d *Dish) Meal {
	return Meal{dish: d}
}

// Simple score by multiplying rating with age.
func score(d *Dish, h *Household) float64 {
	r := h.ratings[d]
	date := h.latestMeal(d).day
	res := r * float64(time.Since(date))
	return res
}

type sorter struct {
	d []*Dish
	r []float64
}

func (s sorter) Len() int {
	return len(s.d)
}

func (s sorter) Less(i, j int) bool {
	return s.r[i] < s.r[j]
}

func (s sorter) Swap(i, j int) {
	td, tr := s.d[i], s.r[i]
	s.d[i], s.r[i] = s.d[j], s.r[j]
	s.d[j], s.r[j] = td, tr
}

func sortDishes(h *Household) []*Dish {
	r := make([]float64, len(h.dishes))
	for i, d := range h.dishes {
		r[i] = score(d, h)
	}
	d := make([]*Dish, len(h.dishes))
	copy(d, h.dishes)
	s := sorter{d, r}
	sort.Sort(sort.Reverse(s))
	return s.d
}
