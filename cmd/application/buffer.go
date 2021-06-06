package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool, 3)

	go func() {
		for i:= 0; i < 10; i++ {
			time.Sleep(10 * time.Millisecond)
			fmt.Printf("+")
		}
		done <- true
	}()

	go func() {
		for i:= 0; i < 10; i++ {
			time.Sleep(9 * time.Millisecond)
			fmt.Printf("-")
		}
		done <- true
	}()

	go func() {
		for i:= 0; i < 10; i++ {
			time.Sleep(8 * time.Millisecond)
			fmt.Printf("￥")
		}
		done <- true
	}()

	for i := 0; i < 10; i++ {
		time.Sleep(5 * time.Millisecond)
		fmt.Printf("/")
	}

	for i := 0; i < 3; i++ {
		<- done
	}
}