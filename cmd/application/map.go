package main

import "fmt"

func main() {
	var m1 = map[string]int{"Pochi": 5, "Kenta": 3, "Sally": 5, "Tommy": 7}
	fmt.Println(m1)

	m1["Kenta"] = 4
	delete(m1, "Tommy")
	fmt.Println(m1)

	fmt.Printf("Pochi = %d\n", m1["Pochi"])

	var m2 = make(map[string]int)
	m2["Tommy"] = 18
	m2["Johon"] = 22
	m2["田中"] = 32
	fmt.Println(m2)
}