//Practice using data structures to build object oriented programs
package main

import "fmt"

func main() {
	fmt.Printf("Price in Euro for 32595 : %s\n", Price.getPriceInEuro(32595))

	RegisterItem(Prices, "banana", 345)
	fmt.Printf("Prices: %v\n", Prices)
	// Re-register item
	RegisterItem(Prices, "banana", 341)
	fmt.Printf("Prices: %v\n", Prices)

	c := new(Cart)
	c.AddItem("eggs")
	c.AddItem("banana")
	fmt.Printf("c.HasItem(%v) = %v\n", "bread", c.HasItem("bread"))

	c.AddItem("chocolat")
	c.AddItem("milk")
	c.Checkout()
}

// Price is the cost of something.
type Price int64

// getPriceInEuro is the string representation of a Price in euro
// Example: 2595 centimes => €25.95
func (p Price) getPriceInEuro() string {
	var Ero string = fmt.Sprint(p/100)
	var Cent string = fmt.Sprint(p%100)

	return "€" + Ero + "." + Cent 
}

// Prices is a map from an item to its price.
var Prices = map[string]Price{
	"eggs":          519,
	"bread":         119,
	"apples":        595,
	"chips":         245,
	"milk":     150,
}

// RegisterItem adds the new item in the prices map.
// If the item is already in the prices map, a warning should be displayed to the user,
// but the value should be overwritten.
func RegisterItem(prices map[string]Price, item string, price Price) {

	
	if _, ok := prices[item]; ok {
		fmt.Println("error: overwrite " + item)
	}

	prices[item] = price

}

// Cart is a struct representing a shopping cart of items.
type Cart struct {
	Items      []string
	TotalPrice Price
}

// HasItem returns whether the shopping cart has the provided item name.
func (c *Cart) HasItem(item string) bool {
	for _, val := range c.Items {
		if val == item {
			return true
		}
	}
	return false
}

// AddItem adds the provided item to the cart and update the cart balance.
// If item is not found in the prices map, then do not add it and print an error.
func (c *Cart) AddItem(item string) {
	if _, ok := Prices[item]; ok {
		c.Items = append(c.Items, item)
		c.TotalPrice += Prices[item]
	} else {
		fmt.Println("error : Item has no price")
	}
}

// Checkout displays the final cart balance and clears the cart completely.
func (c *Cart) Checkout() {
	fmt.Println(c.TotalPrice.getPriceInEuro())
	c.Items = make([]string,0)
	c.TotalPrice = 0
}
