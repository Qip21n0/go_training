package main

import (
	"fmt"
	"log"
	"os"
)

type Point struct{X, Y int}

const fname = "data/points.data"

func Load() []Point {
	var points = []Point{}
	file, err := os.OpenFile(fname, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	var x, y int
	for {
		n, err := fmt.Fscanf(file, "%d %d", &x, &y)
		if n == 0 || err != nil {
			break
		} else {
			points = append(points, Point{x, y})
		}
	}

	file.Close()
	return points
}

func Save(points []Point) int {
	file, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range points {
		fmt.Fprintf(file, "%d %d\n", p.X, p.Y)
	}

	file.Close()
	return len(points)
}

func main() {
	var points = []Point{}
	points = append(points, Point{5, 5})
	points = append(points, Point{7, 8})
	points = append(points, Point{4, 3})
	points = append(points, Point{7, 4})
	points = append(points, Point{9, 5})

	fmt.Println("Writing Data")
	for i, p := range points {
		fmt.Printf("point[%d] = (%d, %d)\n", i+1, p.X, p.Y)
	}
	Save(points)

	data := Load()

	fmt.Println("Reading Data")
	for i, p := range data {
		fmt.Printf("point[%d] = (%d, %d)\n", i+1, p.X, p.Y)
	}
}