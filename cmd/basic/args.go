package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Inout two args")
		os.Exit(1)
	}

	var x, y int
	x, _ = strconv.Atoi(os.Args[1])
	y, _ = strconv.Atoi(os.Args[2])
	
	fmt.Printf("x + y = %d + %d = %d", x, y, x+y)
}