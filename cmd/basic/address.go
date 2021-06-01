package main

import "fmt"

func main() {
	var a = [5]int{1, 2, 3, 4, 5}

	var b = &a[0]
	var c = *b
	fmt.Println(a, b, c)
}