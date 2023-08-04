package main

import (
	"fmt"
	"time"
)

// Godogs available to eat
var Godogs int

type Item struct {
	name     string
	price    int
	quantity int
}

type Cart struct {
	items   []Item
	status  string
	created time.Time
}

var ItemsP []Item
var CartP Cart

func (c *Cart) AddItem(item Item) {
	c.items = append(c.items, item)
}

func (c *Cart) RemoveItemByName(name string) {
	index := -1
	for i, item := range c.items {
		if item.name == name {
			index = i
			break
		}
	}

	if index != -1 {
		c.items = append(c.items[:index], c.items[index+1:]...)
	}
}

func (c *Cart) Checkout() {
	totalPrice := 0
	for _, item := range c.items {
		totalPrice += item.price * item.quantity
	}
	fmt.Printf("Total Price: %d\n", totalPrice)

	c.status = "checked-out"
}

func CreateCart(items ...Item) Cart {
	currentTime := time.Now()
	cart := Cart{
		items:   items,
		status:  "open",
		created: currentTime,
	}
	return cart
}

func main() {
	cart := CreateCart()

	item1 := Item{name: "Product 1", price: 10, quantity: 2}
	item2 := Item{name: "Product 2", price: 20, quantity: 1}

	cart.AddItem(item1)
	cart.AddItem(item2)

	fmt.Printf("Cart Items before checkout: %+v\n", cart.items)

	cart.Checkout()

	fmt.Printf("Cart Status after checkout: %s\n", cart.status)
}
