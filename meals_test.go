package meals

import (
	"reflect"
	"testing"
	"time"
)

func makeI(name string, a ...IngredientAttr) *Ingredient {
	i, _ := NewIngredient(name, a...)
	return i
}

var dishes = []*Dish{
	D("Falukorv med Spagetti", 15, "Falukorv", "Spagetti"),
	D("Falukorv i ugn", 30, "Falukorv", "Ost", "Potatis", "Mjölk", "Smör"),
	D("Potatisbullar", 15, "Potatisbullar", "Lingonsylt", "Bacon"),
	D("Tacos", 30, "Köttfärs", "Tacoskal", "Tacokrydda", "Tomat", "Gräddfil", "Avocado", "Majskorn", "Tacosalsa", "Riven ost"),
	D("Lax i Ugn", 30, "Lax", "Romsås", "Potatis"),
	D("Blodpudding", 15, "Blodpudding", "Lingonsylt"),
	D("Ärtsoppa och Pannkakor", 45, "Ärtsoppa", "Mjöl", "Mjölk", "Smör", "Ägg"),
	D("Pizza", 60, "Mjöl", "Jäst", "Olivolja", "Tomatsås", "Riven ost", "Mozzarella", "Salami", "Champinjoner", "Paprika röd"),
	D("Korvstroganoff", 20, "Falukorv", "Ris", "Tomatpure", "Matgrädde"),
	D("Schnitzel med Klyftpotatis", 30, "Schnitzel", "Potatis", "Broccoli", "Bearnaise"),
	D("Fiskpinnar med Potatismos", 30, "Fiskpinnar", "Potatis", "Smör", "Mjölk", "Stuvad spenat", "Matgrädde", "Muskot"),
	D("Köttfärssås med Spagetti", 20, "Köttfärs", "Spagetti", "Krossade tomater", "Lök", "Tomatpure", "Paprika röd"),
	D("Köttbullar med Potatismos", 20, "Köttbullar", "Potatis", "Smör", "Mjölk", "Matgrädde", "Köttbuljong", "Soja", "Maizena"),
	D("Caesarsallad", 30, "Kycklingfilé", "Bacon", "Romansallad", "Krutonger", "Parmesanost", "Caesardressing"),
	D("Linssoppa med Chorizo", 30, "Linser", "Potatis", "Chorizo fresco", "Hönsbuljong"),
}

func TestCreateDish(t *testing.T) {
	i1, _ := NewIngredient("Falukorv", Sausage, Meat)
	i2, _ := NewIngredient("Makaroner", Pasta, Side)
	d := NewDish("Falukorv med Makaroner", time.Minute*15, i1, i2)
	assert(t, d != nil, "Could not create new dish")
	equals(t, d.Name, "Falukorv med Makaroner")
	equals(t, len(d.Ingrs), 2)
	equals(t, d.Ingrs[1].Attrs[0], Pasta)
}

func TestCreteIngredient(t *testing.T) {
	ingr = nil
	i1, err := NewIngredient("Falukorv", Sausage, Meat)
	ok(t, err)
	_, err = NewIngredient("Falukorv", Sausage, Meat)
	assert(t, err != nil, "Created duplicate ingredients")
	i2 := I("Falukorv")
	equals(t, i1, i2)
}

func TestCompareDishes(t *testing.T) {
	ingr = nil
	d1 := NewDish("Falukorv m. Spagetti", time.Minute,
		makeI("Falukorv", Sausage, Meat), makeI("Spagetti", Pasta, Side))
	d2 := NewDish("Falukorv i Ugn", time.Minute,
		I("Falukorv"), makeI("Potatismos", Potato, Side))

	c := BasicComparator{}
	d := c.Similarity(d1, d1)
	equals(t, d, 1.0)
	d = c.Similarity(d1, d2)
	assert(t, d < 1.0, "%s and %s are not equal", d1.Name, d2.Name)
}

func TestSelection1(t *testing.T) {
	base, _ := time.Parse("2006-01-02", "2014-01-21")
	var history []Meal
	for i := 0; i < 5; i++ {
		history = append(history, MakeMeal(dishes[i], base.AddDate(0, 0, i)))
	}
	sel := newSelection(dishes, 4, history)
	equals(t, len(sel), 4)
	sel2 := newSelection(dishes, 4, history)
	assert(t, !reflect.DeepEqual(sel, sel2), "Expecting two selection calls to return different results")
}

func TestSelection2(t *testing.T) {
	base, _ := time.Parse("2006-01-02", "2014-01-21")
	var history []Meal
	var prefs []Pref
	pBucket := []float64{Bad, Maybe, Good, Favorite}
	for i := 0; i < 5; i++ {
		history = append(history, MakeMeal(dishes[i], base.AddDate(0, 0, i)))
	}
	for i := 0; i < 15; i++ {
		prefs = append(prefs, MakePref(dishes[i], pBucket[i%4]))
	}
}
