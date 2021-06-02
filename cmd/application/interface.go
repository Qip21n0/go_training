package main

import "fmt"

type Animal interface {
	Cry()
}

type Dog struct {}

type Cat struct {}

type Book struct {}

func (d *Dog) Cry() {
	fmt.Println("wanwan")
}

func (c *Cat) Cry() {
	fmt.Println("nyanya")
}

func Cry(any interface{}) {
	a, ok := any.(Animal)
	if !ok {
		fmt.Println("Cannot cry.")
		return
	}
	a.Cry()
}

func main() {
	dog := new(Dog)
	cat := new(Cat)
	
	Cry(dog)
	Cry(cat)

	book := new(Book)
	Cry(book)
}