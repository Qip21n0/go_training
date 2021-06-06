package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fname := "data/sample5.dat"

	txt := "Tjis is a sample text."
	werr := ioutil.WriteFile(fname, ([]byte)(txt), 0666)
	if werr != nil {
		_ = fmt.Errorf("not write %s", fname)
	}

	data, rerr := ioutil.ReadFile(fname)
	if rerr != nil {
		_ = fmt.Errorf("not read %s", fname)
	} else {
		fmt.Println(string(data))
	}
}