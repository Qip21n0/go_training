package main

import "fmt"

func main() {
	var x [5]int = [5]int{1, 2, 3, 4, 5}
	var px *int

	px = &x[0]
	*px = 123
	fmt.Println("x =", x)

	var p *int = new(int)
	*p = 123
	fmt.Println("p =", *p)
}