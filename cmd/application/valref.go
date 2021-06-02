package main

import "fmt"

func byVal(a int) {
	a = a + 1
}

func byRef(a *int) {
	*a = *a + 1
}

func main() {
	n := 1

	byVal(n)
	fmt.Println(n)

	byRef(&n)
	fmt.Println(n)
}