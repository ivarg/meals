package meals

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Vad köra
// - bortse från senaste x dagar.
// - gå igenom alla kända rätter och räkna poäng
// - sortera efter poäng

func TestSortDishes(t *testing.T) {
	r := map[*Dish]float64{
		dishes[0]: .5,
		dishes[1]: .4,
		dishes[2]: .6,
	}
	h := Household{ratings: r, dishes: dishes[:3]}
	dd := h.sortDishes()
	equals(t, dd[0], h.dishes[2])
}

func TestTruncStr(t *testing.T) {
	tvec := []struct {
		s  string
		ln uint
		ts string
	}{
		{"qwert", 5, "qw..."},
		{"qwert", 4, "q..."},
		{"qwert", 3, "q.."},
		{"qwert", 2, "q."},
		{"qwert", 1, "q"},
		{"qw", 5, "qw"},
	}
	for _, test := range tvec {
		if truncString(test.s, test.ln) != test.ts {
			t.Error()
		}
	}
}

func Generate(n int, start time.Time, h *Household) []Meal {
	own := 4

	// Sort household dishes
	k := h.sortDishes()

	// Select top household dishes and sample from all dishes
	if len(k) < own {
		own = len(k)
	}
	mysel := k[:own]
	newsel := newDishSel(n-own, h.dishes)
	sel := make([]*Dish, n)
	sel = mysel
	for _, d := range newsel {
		sel = append(sel, d)
	}

	// Create meals for the selected dishes
	var mls []Meal
	for i, d := range sel {
		mls = append(mls, MakeMeal(d, start.AddDate(0, 0, i)))
	}

	return mls
}

func simulate(h *Household, start time.Time) {
	// One simulation round:
	// - Generate new selection
	//	+ Select at most 5 known dishes
	//	+ Randomize among unknown
	// - Save to household history
	// - Randomize dish ratings
	// - Repeat
	//tot := 5

	mls := Generate(5, start, h)
	h.save(mls)

	// Rate the selection
	for _, m := range mls {
		if _, ok := h.ratings[m.dish]; !ok {
			h.ratings[m.dish] = .5
		} else {
			mod := 1
			if rand.NormFloat64() < 0 {
				mod = -1
			}
			h.ratings[m.dish] += (1 - h.ratings[m.dish]) * rand.Float64() * float64(mod)
			if h.ratings[m.dish] > 1 {
				h.ratings[m.dish] = 1
			}
			if h.ratings[m.dish] < 1 {
				h.ratings[m.dish] = 0
			}
		}
	}
}

func printTopDishes(h *Household, n int) {
	s := dishSorter{h.dishes, func()}
}

func TestSimulate(t *testing.T) {
	h := Household{ratings: make(map[*Dish]float64), dishes: dishes[:3]}
	now, _ := time.Parse("2006-01-02", "2014-01-01")

	n := 1
	for i := 0; i < n; i++ {
		simulate(&h, now)
	}
	printTopDishes(&h, 3)

	ds, cnt := h.mostSelected(10)
	for i, d := range ds {
		fmt.Println(truncString(d.Name, 10), cnt[i])
	}
}
