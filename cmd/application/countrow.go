package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fname := "data/test.txt"

	file, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	count := 0
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		if err := scan.Err(); err != nil {
			break
		}
		count++
	}
	fmt.Printf("Row numbers of %s: %d", fname, count)
}