package main

import "fmt"

// An interface is a listing of method signatures (name, parameters, return value).
// All types that have conforming method implementations automatically implement the
// interface (without an explicit implementation hint).
type Costing interface {
	Cost() int
}

// Structs are a simple way to bundle data that is logically related.
type Item struct {
	name     string
	price    int // No fractional part for simplicity reasons. Price for a single quanity.
	quantity int
}

// We implement the standard interface Stringer here for use in formatting functions.
// Its purpose is to return a string representation.
func (i Item) String() string {
	return fmt.Sprintf("%v %v x %v$", i.name, i.quantity, i.price)
}

// Here we implement our own interface Consting for the Item type. Note that there is
// no explicit statement about our implementation intend in Golang. The compiler sees
// that all signatures from the interface are implemented and that is enough.
func (i Item) Cost() int {
	return i.price * i.quantity
}

type Discount struct {
	percent int
}

// One classic option for calculating decimal numbers that are not integers is to use
// floating point types. Be careful though as they are not suited for use cases where exactness
// is required. You can look into alternative implementations for decimal calculation instead.
// In our case we just want to cut the part after the comma and therefore convert the result to int.
func (d Discount) Calc(costing Costing) int {
	return int((float32(d.percent) / float32(100)) * float32(costing.Cost()))
}

// Structs can also encapsulate data structures like slices.
type Cart struct {
	items []Item
}

// Methods are defined for a specific type which is noted between the func keyword and the method name.
// Note that we are using a value receiver here (no asterik, therefore no pointer) since reading is all
// we need. Getting a copy of Cart upon calling is totally fine here.
func (c Cart) Cost() int {
	sum := 0

	for _, item := range c.items {
		sum += item.Cost()
	}

	return sum
}

// Another implementation of Golangs standard Stringer interface, e.g. for custom printing.
func (c Cart) String() string {
	return fmt.Sprintf("%v", c.items)
}

// Like functions methods may also return multiple values.
func (c Cart) Find(name string) (*Item, int) {
	for index, current := range c.items {
		if current.name == name {
			// Remember to create the pointer from the original
			// element. current holds a copy due to range behavior.
			return &c.items[index], index
		}
	}

	return nil, -1
}

// Note that we are using a pointer receiver here since we need to
// modify the Cart value that this method is called upon. When using
// a pointer receiver the value is not copied, we can access the original
// Cart via its address in memory.
func (c *Cart) Add(item Item) {
	existing, _ := c.Find(item.name)

	if existing != nil {
		// Item already in cart, just increase the quantity.
		// We assume that prices are not different for simplicity.
		existing.quantity += item.quantity
	} else {
		c.items = append(c.items, item)
	}
}

func (c *Cart) Remove(name string) {
	_, index := c.Find(name)

	if index >= 0 {
		// Remove found element by using reslicing.
		c.items = append(c.items[:index], c.items[index+1:]...)
	}
}

func (c Cart) Size() int {
	return len(c.items)
}
