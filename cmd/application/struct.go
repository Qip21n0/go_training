package main

import "fmt"

type Member struct {
	id int
	name string
	age int
}

type Coordinate struct {x, y int}

type Circle struct {
	Center Coordinate
	Radius int
}

func main() {
	var members = []Member{}

	members = append(members, Member{0xA010, "ポチ", 17})
	members = append(members, Member{0xA021, "犬田", 16})
	members = append(members, Member{0xB023, "pq", 21})

	mem := Member{0xC123, "ζ", 58}
	members = append(members, mem)

	fmt.Println(members)

	var pos = Coordinate{}
	pos.x, pos.y = 123, 45
	fmt.Println(pos)

	c := Circle{Coordinate{50, 50}, 70}
	fmt.Println(c)
}