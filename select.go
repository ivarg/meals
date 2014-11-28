package meals

import (
	"fmt"
	"math/rand"
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

func (h *Household) mostSelected(n int) ([]*Dish, []int) {
	return nil, nil
}

func (h *Household) save(mls []Meal) {
	h.hist = append(h.hist, mls...)
}

// Simple score by multiplying rating with age.
func score(d *Dish, h *Household) float64 {
	r := h.ratings[d]
	date := h.latestMeal(d).day
	res := r * float64(time.Since(date))
	return res
}

func truncString(s string, ln uint) string {
	if ln == 0 {
		return ""
	}
	if uint(len(s)+3) <= ln {
		return s
	}
	pad := "..."
	var ts string
	if ln < 4 {
		pad = pad[:ln-1]
		ts = s[0:1]
	} else {
		ts = s[:(ln - 3)]
	}
	return fmt.Sprintf("%s%s", ts, pad)
}

type dishSorter struct {
	d    []*Dish
	less func(i, j int) bool
	swap func(i, j int)
}

func (ds dishSorter) Len() int {
	return len(ds.d)
}

func (ds dishSorter) Less(i, j int) bool {
	return ds.less(i, j)
}

func (ds dishSorter) Swap(i, j int) {
	td := ds.d[i]
	ds.d[i] = ds.d[j]
	ds.d[j] = td
	ds.swap(i, j)
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

func (h *Household) sortDishes() []*Dish {
	r := make([]float64, len(h.dishes))
	for i, d := range h.dishes {
		r[i] = score(d, h)
	}
	d := make([]*Dish, len(h.dishes))
	copy(d, h.dishes)
	ds := dishSorter{
		d,
		func(i, j int) bool { return r[i] < r[j] },
		func(i, j int) { t := r[i]; r[i] = r[j]; r[j] = t },
	}
	//s := sorter{d, r}
	sort.Sort(sort.Reverse(ds))
	return ds.d
}

func newDishSel(size int, not []*Dish) []*Dish {
	smpl := make([]*Dish, size)
	for i := range smpl {
		smpl[i] = dishes[rand.Intn(len(dishes))]
	}
	return smpl
}
