package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, m int
	var name string

	fmt.Printf("n = ")
	fmt.Scan(&n)

	fmt.Printf("m = ")
	fmt.Scanf("%d", &m)

	fmt.Printf("Your Name: ")
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	name = scanner.Text()

	fmt.Printf("n + m = %d\n", n+m)
	fmt.Printf("Hello! %s\n", name)
}