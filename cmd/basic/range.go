package main

import "fmt"

func main() {
	var x, y, z int
	x = 3
	z = 100
	{
		var x = 12
		fmt.Println(x, z)
	}
	y = x * x
	fmt.Println(y)
}