package main

import "fmt"

func main() {
	const PI = 3.14
	fmt.Println(PI)
	
	var n float32 = 2.345e02
	var n1 int = 012
	var n2 int = 0x23
	fmt.Println(n, n1, n2)
	
	var r rune = 'あ'
	var r1 = rune(0x3044)
	fmt.Println(r, r1)

	l := 'a'
	l1 := '\n'
	l2 := '\u3042'
	fmt.Println(l, l1, l2)

	var c complex128 = 1 + 2i
	var c1 complex128 = complex(5, 9)
	var c2 complex128 = -1.25i
	fmt.Println(c, c1, c2)

	var s = "abc\tdef"
	var s1 = `abc\tdef`
	fmt.Println(s, s1)
}