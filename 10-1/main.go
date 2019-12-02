package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type point struct {
	X  int
	Y  int
	VX int
	VY int
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	points := make([]*point, 0)

	s := bufio.NewScanner(f)
	for s.Scan() {

		p := &point{}
		fmt.Sscanf(s.Text(), "position=<%3d,  %3d> velocity=< %3d,  %3d>", &p.X, &p.Y, &p.VX, &p.VY)
		fmt.Println(p)
		points = append(points, p)

	}

}

func ppoint(points []*point) {

}
