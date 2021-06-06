package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func printjob(s string) {
	defer wg.Done()
	for i:= 0; i < 10; i++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("%s", s)
	}
}

func main() {
	wg.Add(3)

	fmt.Println("A start")
	go printjob("A")

	fmt.Println("B start")
	go printjob("B")

	fmt.Println("C start")
	go printjob("C")

	wg.Wait()
	
	fmt.Printf("\nfin.\n")
}