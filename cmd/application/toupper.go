package main

import (
	"fmt"
	"strings"
)

func main() {
	ch := make(chan string)

	go func() {
		for {
			str := <- ch
			fmt.Printf("%s\n", strings.ToUpper(str))
		}
	}()

	var s string

	fmt.Println("Input a word.(End with \"quit\")")
	for {
		fmt.Scan(&s)
		if s == "quit" {
			break
		}
		ch <- s
	}
}