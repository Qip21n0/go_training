package main

import "fmt"

func main() {
	var a = [...]int{1, 3, 5, 7, 9}
	fmt.Printf("type(a) = %T\n", a)

	var s = a[0:4]
	fmt.Printf("type(s) = %T\n", s)

	for i, v := range s {
		fmt.Printf("s[%d] = %d\n", i, v)
	}

	s[1], s[3] = 0, 0
	fmt.Println(s)

	var b = make([]int, 5, 20)

	b[1], b[3] = 18, 38
	fmt.Println(b)

	b = append(b, 88)
	fmt.Println(b)

	var d = make([]int, 5, 20)
	copy(d, s)
	fmt.Println(d)
}