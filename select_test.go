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
	dd := sortDishes(&h)
	equals(t, dd[0], h.dishes[2])
}

func TestSimulate(t *testing.T) {
	h := Household{ratings: make(map[*Dish]float64)}
	// One simulation round:
	// - Generate new selection
	//	+ Select at most 5 known dishes
	//	+ Randomize among unknown
	// - Save to household history
	// - Randomize dish ratings
	// - Repeat
	//tot := 5
	own := 3
	k := sortDishes(&h)
	fmt.Println(k)
	if len(k) < own {
		own = len(k)
	}
}

func TestExp(t *testing.T) {
	n := 50
	k := 5
	p := .6
	s := make([]int, n)
	res := make(map[int]int, n)
	for i, _ := range s {
		s[i] = i
	}

	for i, _ := range s {

		if k == 0 {
			break
		}
		if p >= rand.Float64() {
			res[i]++
		}
	}

	for i, _ := range s {
		fmt.Printf("%d (%d)\n", i, res[i])
	}
}
