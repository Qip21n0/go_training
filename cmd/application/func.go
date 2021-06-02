package main

import "fmt"

type Dog struct {
	name string
	age int
}

func Prints(args ... interface{}) {
	fmt.Println(args...)
}

func main() {
	Prints("vegjil")
	Prints(-1, "bdjsbe", "高冨富士山")

	dog := Dog{"Pochi", 12}
	Prints(dog, "nani")
}