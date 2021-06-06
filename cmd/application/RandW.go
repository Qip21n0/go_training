package main

import (
	"fmt"
	"os"
)

func main() {
	fname := `data/sample.txt`
	var file *os.File
	var err error

	file, err = os.Create(fname)
	if err != nil {
		_ = fmt.Errorf("not open %s", fname)
	}

	txt := "This is a sample text."
	file.Write(([]byte)(txt))
	file.Close()

	BUFSIZE := 1024

	file, err = os.Open(fname)
	if err != nil {
		_ = fmt.Errorf("not open %s", fname)
	}
	defer file.Close()

	buf := make([]byte, BUFSIZE)
	for {
		n, err := file.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			_ = fmt.Errorf("Error while reading %s", fname)
			break
		}
		fmt.Printf(string(buf[:n]))
	}
}