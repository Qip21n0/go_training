package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Input two args.")
		os.Exit(1)
	}

	ifile, ierr := os.Open(os.Args[1])
	if ierr != nil {
		_ = fmt.Errorf("not open %s", os.Args[1])
	}
	ofile, oerr := os.Create(os.Args[2])
	if oerr != nil {
		_ = fmt.Errorf("not open %s", os.Args[2])
	}

	r := bufio.NewReader(ifile)
	w := bufio.NewWriter(ofile)

	for {
		c, err := r.ReadByte()
		if err != nil {
			break
		}
		oerr := w.WriteByte(c)
		if oerr != nil {
			fmt.Println(oerr)
			break
		}
	}
	err := w.Flush()
	if err != nil {
		fmt.Println(oerr)
	}
}