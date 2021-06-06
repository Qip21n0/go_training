package main

import (
	"fmt"
	"time"
)

func printjob(s string) {
	for i:= 0; i < 10; i++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("%s", s)
	}
}

func main() {
	fmt.Println("A start")
	go printjob("A")

	fmt.Println("B start")
	go printjob("B")

	fmt.Println("C start")
	go printjob("C")

	time.Sleep(1 * time.Second)
	fmt.Printf("\nfin.\n")
}