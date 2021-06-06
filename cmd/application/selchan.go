package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan rune)
	ch2 := make(chan int)
	done := make(chan bool)

	go func() {
		s := "ABCDEFG"
		for _, c := range s {
			time.Sleep(10 * time.Millisecond)
			fmt.Printf("transmission from ch1: %c\n", c)
			ch1 <- c
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 8; i++ {
			time.Sleep(8 * time.Millisecond)
			fmt.Printf("transmission from ch2: %d\n", i+1)
			ch2 <- i+1
		}
		done <- true
	}()

	defer fmt.Println("fin.")
	count := 0
	for {
		select {
		case r := <- ch1:
			fmt.Printf("receive from ch1: %c\n", r)
		case n := <- ch2:
			fmt.Printf("receive from ch2: %d\n", n)
		case <- done:
			count++
			if count > 1 {
				return
			}
		}
	}
}