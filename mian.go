package main

import (
	"fmt"

	process_oriented "github.com/promacanthus/design-patterns/process-oriented"
)

func main() {
	sc := process_oriented.NewShoppingCart()
	sc.AddItem(process_oriented.ShoppingCartItem{
		Name:  "item1",
		Price: 10,
	})

	items := sc.GetItem()
	fmt.Println(items)
	items[0] = process_oriented.ShoppingCartItem{
		Name:  "new",
		Price: 0,
	}
	fmt.Printf("%+v", *sc)
}
