package main

import "fmt"

func Print(a interface{}) {
	v, ok := a.(int)
	if ok {
		fmt.Printf("%d(10)\t%x(16)\n", v, v)
	} else {
		fmt.Println(a)
	}
}

func main() {
	Print(23)
	Print("hello")
	Print(123.45)
	Print(1+2i)
}