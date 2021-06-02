package main

import "fmt"

type Value int

func (v Value) Twice() Value {
	return 2 * v
}

func main() {
	var n, m Value
	n = 1
	m = n.Twice()

	fmt.Println(n, m)
	fmt.Println(m.Twice())
}