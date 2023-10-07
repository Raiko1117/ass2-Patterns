package main

import "fmt"

func DecoratorExample() {
	pizza := &VeggieMania{}

	pizzaWithCheese := &ToppingDecorator{pizza: pizza, toppingPrice: 10, toppingName: "Cheese"}

	pizzaWithCheeseAndTomato := &ToppingDecorator{pizza: pizzaWithCheese, toppingPrice: 7, toppingName: "Tomato"}

	fmt.Printf("Price pizza with cheese and tomato: $%d\n", pizzaWithCheeseAndTomato.GetPrice())
	fmt.Printf("Pizza Description: %s\n", pizzaWithCheeseAndTomato.GetDescription())
}

func FactoryExample() {
	ak47, _ := createGun("ak47")
	P90, _ := createGun("P90")

	printDetails(ak47)
	printDetails(P90)
}

func main() {
	FactoryExample()
	//DecoratorExample()
}
