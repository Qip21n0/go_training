package main

import "fmt"

func main() {
	const PI = 3.14
	const (
		Sun int = 1 << iota
		Mon
		Tue
		Wed
		Thu
		Fri
		Sut
	)

	fmt.Println("π =", PI)
	fmt.Println(Sun, Mon, Tue, Wed, Thu, Fri, Sut)
}