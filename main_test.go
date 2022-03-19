package main

import (
	"fmt"
	"testing"
)

// Often it is a good idea to test only one function or method per
// test function. This keeps tests clean and maintainable. There are
// exceptions though (see below).
func TestItemCost(t *testing.T) {
	item := Item{name: "banana", price: 1, quantity: 7}
	cost := item.Cost()

	checkEquals(t, cost, 7)
}

func TestItemString(t *testing.T) {
	item := Item{name: "banana", price: 1, quantity: 7}
	stringRepr := item.String()

	checkEquals(t, stringRepr, "banana 7 x 1$")
}

func TestDiscountCalc(t *testing.T) {
	item := Item{name: "ice cream", price: 2, quantity: 10}

	discount := Discount{10}
	saved := discount.Calc(item)
	checkEquals(t, saved, 2)

	discount = Discount{7}
	saved = discount.Calc(item)
	checkEquals(t, saved, 1)

	cart := Cart{}
	cart.Add(item)
	cart.Add(Item{name: "banana", price: 1, quantity: 7})
	saved = discount.Calc(cart)
	checkEquals(t, saved, 1)
}

// Sometimes it makes sense to test multiple operations in a single test function.
// We chose this option here since results of operations like add, remove, size
// are tightly related to each other.
func TestCart(t *testing.T) {
	cart := Cart{}

	banana := Item{name: "banana", price: 1, quantity: 7}
	cart.Add(banana)
	size := cart.Size()
	checkEquals(t, size, 1)

	item, _ := cart.Find("banana")
	checkEquals(t, *item, banana)

	eggs := Item{name: "eggs", price: 2, quantity: 3}
	cart.Add(eggs)
	fmt.Println(cart)
	size = cart.Size()
	checkEquals(t, size, 2)

	item, _ = cart.Find("eggs")
	checkEquals(t, *item, eggs)

	cart.Add(Item{name: "banana", price: 1, quantity: 2})
	size = cart.Size()
	checkEquals(t, size, 2)

	item, _ = cart.Find("banana")
	increasedBanana := Item{name: "banana", price: 1, quantity: 9}
	checkEquals(t, *item, increasedBanana)

	cart.Remove("eggs")
	size = cart.Size()
	checkEquals(t, size, 1)
	item, _ = cart.Find("eggs")
	var itemNotFound *Item = nil
	checkEquals(t, item, itemNotFound)

	cart.Remove("ketchup")
	size = cart.Size()
	checkEquals(t, size, 1)

	cart.Remove("banana")
	size = cart.Size()
	checkEquals(t, size, 0)
	item, _ = cart.Find("banana")
	checkEquals(t, item, itemNotFound)
}

// Extracting commonly used code can be helpful for tests, too. Here we are using
// the empty interface type in order to place no restrictions on the passed parameter
// types for comparison. Note that interfaces in Golang consist of two components:
// A type and a value. When you compare two nil parameters this might result in
// false if one of the parameters has nil type and the other a specific pointer type
// like *Item. Therefore, both parameters need to have the same underlying (specific)
// type and value in order to be considered equal.
func checkEquals(t *testing.T, actual, expected interface{}) {
	t.Helper()
	if actual != expected {
		t.Errorf("Expected %T %v but got %T %v", expected, expected, actual, actual)
	}
}
