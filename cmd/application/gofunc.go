package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for _, c := range "ABC" {
		ch := c
		fmt.Printf("%c start\n", c)
		go func() {
			wg.Add(3)
			for i:= 0; i < 10; i++ {
				time.Sleep(10 * time.Millisecond)
				fmt.Printf("%c", ch)
			}
			defer wg.Done()
		}()
	}

	wg.Wait()

	fmt.Printf("\nfin.\n")
}