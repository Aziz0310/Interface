package main

import (
	"fmt"
)

type Product interface {
	describe() string
}

type Phone struct {
	model  string
	color  string
	memory int
}

type Car struct {
	model   string
	color   string
	price   int
	release int
}

type Book struct {
	name    string
	author  string
	genre   string
	release int
}

type Pc struct {
	model  string
	memory int
	ram    int
	price  int
}

func (p Phone) describe() string {
	return fmt.Sprintf("model: %v  color: %v  memory: %v", p.model, p.color, p.memory)
}

func (c Car) describe() string {
	return fmt.Sprintf("model: %v color: %v price: %v release: %v", c.model, c.color, c.price, c.release)
}

func (b Book) describe() string {
	return fmt.Sprintf("name: %v  author: %v  genre: %v  release: %v", b.name, b.author, b.genre, b.release)
}

func (l Pc) describe() string {
	return fmt.Sprintf("model: %v memory: %v  ram: %v price: %v", l.model, l.memory, l.ram, l.price)
}

func main() {
	obj1 := Phone{
		model:  "Iphone",
		color:  "blue",
		memory: 128,
	}

	Describe(obj1)

	obj2 := Car{
		model:   "Mustang",
		color:   "Yellow",
		price:   245000,
		release: 1957,
	}

	Describe(obj2)

	obj3 := Book{
		name:    "Mumu",
		author:  "Ivan Turgenev",
		genre:   "fiction",
		release: 1854,
	}

	Describe(obj3)

	obj4 := Pc{
		model:  "MacBook Air",
		memory: 1,
		ram:    8,
		price:  1500,
	}
	Describe(obj4)
}

func Describe(obj Product) {
	switch obj.(type) {
	case Phone:
		fmt.Println("Phone:")
	case Car:
		fmt.Println("Car:")
	case Book:
		fmt.Println("Book:")
	case Pc:
		fmt.Println("Pc:")
	default:
		fmt.Println("Unknown:")
	}
	fmt.Println(obj.describe())
}
