package main

type Pizza interface {
	GetPrice() int
	GetDescription() string
}

type VeggieMania struct {
}

func (p *VeggieMania) GetPrice() int {
	return 15
}

func (p *VeggieMania) GetDescription() string {
	return "VeggieMania"
}

type ToppingDecorator struct {
	pizza        Pizza
	toppingPrice int
	toppingName  string
}

func (c *ToppingDecorator) GetPrice() int {
	return c.pizza.GetPrice() + c.toppingPrice
}

func (c *ToppingDecorator) GetDescription() string {
	return c.pizza.GetDescription() + ", " + c.toppingName
}
