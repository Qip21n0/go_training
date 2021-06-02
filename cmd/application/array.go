package main

import "fmt"

func main() {
	var a [5]int = [5]int{2, 3, 5, 7, 11}

	for i, v := range a {
		fmt.Printf("a[%d] = %d\n", i, v)
	}

	a[1] = 0
	a[3] = 0
	fmt.Println(a)

	var b [5]int

	b[0], b[2] = 99, 98
	fmt.Println(b)
}