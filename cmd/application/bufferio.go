package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fname := `data/sample.dat`
	var file *os.File
	var err error

	file, err = os.Create(fname)
	if err != nil {
		_ = fmt.Errorf("not open %s", fname)
	}

	var txt string
	for i := 0; i < 5; i++ {
		txt = "This is a sample text." + "line[" + strconv.Itoa(i+1) + "]\n"
		file.Write(([]byte)(txt))
	}
	file.Close()

	file, err = os.Open(fname)
	if err != nil {
		_ = fmt.Errorf("not open %s", fname)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		if err := scan.Err(); err != nil {
			_ = fmt.Errorf("not read %s", fname)
			break
		}
		fmt.Println(scan.Text())
	}
}