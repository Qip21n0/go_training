package main

import (
	"fmt"
	"os"
	"math/rand"
	"time"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Input two args.")
		os.Exit(1)
	}
	n, _ := strconv.Atoi(os.Args[1])
	fname := os.Args[2]
	var file *os.File
	var err error

	file, err = os.Create(fname)
	if err != nil {
		_ = fmt.Errorf("not open %s", fname)
	}

	var txt string
	rand.Seed(time.Now().UnixNano())
	
	for i := 0; i < n; i++ {
		txt = strconv.Itoa(rand.Intn(n)) + "\n"
		file.Write(([]byte)(txt))
	}
	file.Close()
}