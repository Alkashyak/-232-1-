package main
import "fmt"

type PizzaIngredients struct{
    dough    string
    sauces   []string
    toppings []string
    cheeses  []string
}

type Pizza struct{ // Product
    PizzaIngredients
}

func NewPizza(dough string, sauces []string, toppings []string, cheeses []string) Pizza{
    var p Pizza
    p.dough = dough
    p.sauces = sauces
    p.toppings = toppings
    p.cheeses = cheeses
    
    return p
}

func (p Pizza) String() string{
    str := p.dough
    for _, v := range p.sauces {
        str += " " + v
    }
    for _, v := range p.toppings {
        str += " " + v
    }
    for _, v := range p.cheeses {
        str += " " + v
    }
    return str
}

type PizzaRecipe interface{ // Builder
    newPizza() PizzaRecipe

    // Builder type 1
    makeBase() PizzaRecipe
    makeFilling() PizzaRecipe
    Bake() Pizza
    
    // Builder type 2
    setDough(dough string) PizzaRecipe
    addSauce(sauce string) PizzaRecipe
    addTopping(topping string) PizzaRecipe
    addCheese(cheese string) PizzaRecipe
}

type PizzaChef struct{ // Director
    recipe PizzaRecipe
}

func (c *PizzaChef) setRecipe(recipe PizzaRecipe) {
    c.recipe = recipe
}

func (c *PizzaChef) cook() Pizza{
    return c.recipe.
                newPizza().
                makeBase().
                makeFilling().
                Bake()
}

type MargheritaRecipe struct{
    PizzaIngredients
}

func (r *MargheritaRecipe) newPizza() PizzaRecipe{
    r.dough    = ""
    r.sauces   = nil
    r.toppings = nil
    r.cheeses  = nil
    return r
}

func (r *MargheritaRecipe) makeBase() PizzaRecipe{
    r.dough = "MargheritaDough"
    return r
}
func (r *MargheritaRecipe) makeFilling() PizzaRecipe{
    r.sauces = append(r.sauces, "Tomato paste")
    r.toppings = append(r.toppings, "Tomato", "Basil leaves")
    r.cheeses = append(r.cheeses, "Mozzarella", "Parmesan")
    return r
}
func (r *MargheritaRecipe) Bake() Pizza{
    return NewPizza(r.dough, r.sauces, r.toppings, r.cheeses)
}
func (r *MargheritaRecipe) setDough(dough string) PizzaRecipe{
    return r
}
func (r *MargheritaRecipe) addSauce(sauce string) PizzaRecipe{
    return r
}
func (r *MargheritaRecipe) addTopping(topping string) PizzaRecipe{
    return r
}
func (r *MargheritaRecipe) addCheese(cheese string) PizzaRecipe{
    return r
}

type CustomRecipe struct{
    PizzaIngredients
}

func (r *CustomRecipe) newPizza() PizzaRecipe{
    r.dough    = ""
    r.sauces   = nil
    r.toppings = nil
    r.cheeses  = nil
    return r
}

func (r *CustomRecipe) makeBase() PizzaRecipe{
    return r
}
func (r *CustomRecipe) makeFilling() PizzaRecipe{
    return r
}
func (r *CustomRecipe) Bake() Pizza{
    p := NewPizza(r.dough, r.sauces, r.toppings, r.cheeses)
    return p
}

func (r *CustomRecipe) setDough(dough string) PizzaRecipe{
    r.dough = dough
    return r
}
func (r *CustomRecipe) addSauce(sauce string) PizzaRecipe{
    r.sauces = append(r.sauces, sauce)
    return r
}
func (r *CustomRecipe) addTopping(topping string) PizzaRecipe{
    r.toppings = append(r.toppings, topping)
    return r
}
func (r *CustomRecipe) addCheese(cheese string) PizzaRecipe{
    r.cheeses = append(r.cheeses, cheese)
    return r
}

type PepperoniRecipe struct {
	PizzaIngredients
}

func (rec *PepperoniRecipe) newPizza() PizzaRecipe {
	rec.toppings = nil
	rec.cheeses = nil
	rec.sauces = nil
	rec.dough = ""
	return rec
}

func (rec *PepperoniRecipe) makeBase() PizzaRecipe {
	rec.dough = "PepperoniDough"
	return rec
}
func (rec *PepperoniRecipe) makeFilling() PizzaRecipe {
	rec.sauces = append(rec.sauces, "Tomato paste")
	rec.cheeses = append(rec.cheeses, "Mozzarella", "Fontina")
	rec.toppings = append(rec.toppings, "Pepperoni", "Garlic")
	return rec
}
func (rec *PepperoniRecipe) Bake() Pizza {
	return NewPizza(rec.dough, rec.sauces, rec.toppings, rec.cheeses)
}
func (rec *PepperoniRecipe) setDough(dough string) PizzaRecipe {
	return rec
}
func (rec *PepperoniRecipe) addSauce(sauce string) PizzaRecipe {
	return rec
}
func (rec *PepperoniRecipe) addTopping(topping string) PizzaRecipe {
	return rec
}
func (rec *PepperoniRecipe) addCheese(cheese string) PizzaRecipe {
	return rec
}

type HawaiianRecipe struct {
	PizzaIngredients
}

func (rec *HawaiianRecipe) newPizza() PizzaRecipe {
	rec.dough = ""
	rec.sauces = nil
	rec.toppings = nil
	rec.cheeses = nil
	return rec
}

func (rec *HawaiianRecipe) makeBase() PizzaRecipe {
	rec.dough = "HawaiianDough"
	return rec
}
func (rec *HawaiianRecipe) makeFilling() PizzaRecipe {
	rec.sauces = append(rec.sauces, "Tomato paste")
	rec.cheeses = append(rec.cheeses, "Mozzarella")
	rec.toppings = append(rec.toppings, "Pineapple", "Onion", "Bacon")
	return rec
}
func (rec *HawaiianRecipe) Bake() Pizza {
	return NewPizza(rec.dough, rec.sauces, rec.toppings, rec.cheeses)
}
func (rec *HawaiianRecipe) setDough(dough string) PizzaRecipe {
	return rec
}
func (rec *HawaiianRecipe) addSauce(sauce string) PizzaRecipe {
	return rec
}
func (rec *HawaiianRecipe) addTopping(topping string) PizzaRecipe {
	return rec
}
func (rec *HawaiianRecipe) addCheese(cheese string) PizzaRecipe {
	return rec
}


func main() {
    var chef PizzaChef

    var count int
    fmt.Scan(&count)
    for count > 0 {
        var pizzaType string
        fmt.Scan(&pizzaType)
        
        var p Pizza
        switch pizzaType {
            case "hawaiian": chef.setRecipe(&HawaiianRecipe{})
                             p = chef.cook()
            case "pepperoni": chef.setRecipe(&PepperoniRecipe{})
                              p = chef.cook()
            case "margherita": chef.setRecipe(&MargheritaRecipe{})
                               p = chef.cook()
            case "custom": p = (&CustomRecipe{}).newPizza().
                                              setDough("CustomDough").
                                              addSauce("CustomSauce").
                                              addTopping("CustomTopping1").
                                              addTopping("CustomTopping2").
                                              addCheese("CustomCheese").
                                              Bake()
        }
        fmt.Println(p)
        
        count--
    }
}
