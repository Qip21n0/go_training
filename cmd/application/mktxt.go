package main

import (
	"fmt"
	"os"
	"math/rand"
	"time"
	"strconv"
)

func main() {
	fname := "data/test.txt"
	var file *os.File
	var err error

	file, err = os.Create(fname)
	if err != nil {
		_ = fmt.Errorf("not open %s", fname)
	}

	var txt string
	rand.Seed(time.Now().UnixNano())
	n := 1000000
	for i := 0; i < n; i++ {
		txt = strconv.Itoa(rand.Intn(n)) + "\n"
		file.Write(([]byte)(txt))
	}
	file.Close()
}