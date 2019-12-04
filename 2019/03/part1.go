package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type p struct {
	x, y int
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Scan()
	wire1 := strings.Split(s.Text(), ",")
	s.Scan()
	wire2 := strings.Split(s.Text(), ",")

	wp1 := makeL(wire1)
	wp2 := makeL(wire2)
	lines1 := createLines(wp1)
	lines2 := createLines(wp2)
	// fmt.Println("Lines: ", lines1, lines2)

	for _, l1 := range lines1 {
		for _, l2 := range lines2 {
			p, e := Intersection(l1, l2)
			if e != nil {
				continue
			}
			fmt.Println("Point of intersect", p)
		}
	}

}

type Line struct {
	slope float64
	yint  float64
}

func CreateLine(a, b p) Line {
	fmt.Println("create line")
	ys := float64(b.y - a.y)
	xs := float64(b.x - a.x)
	var slope float64
	if xs == 0 {
		slope = 0
	} else {
		slope = ys / xs
	}
	//fmt.Println(slope)
	yint := float64(a.y) - slope*float64(a.x)
	fmt.Println(yint)
	return Line{slope, yint}
}

func createLines(points []p) []Line {
	lines := make([]Line, len(points)-1)
	for i := range points {
		if i == len(points)-1 {
			break
		}
		fmt.Println(points[i], points[i+1])
		lines = append(lines, CreateLine(points[i], points[i+1]))
	}
	return lines
}

func Intersection(l1, l2 Line) (p, error) {
	if l1.slope == l2.slope {
		return p{}, errors.New("The lines do not intersect")
	}
	x := (l2.yint - l1.yint) / (l1.slope - l2.slope)
	y := l1.slope*x + l1.yint

	return p{int(x), int(y)}, nil
}

func makeL(wire []string) []p {
	pp := make([]p, 1)

	pp[0] = p{
		x: 0,
		y: 0,
	}

	for i, d := range wire {
		pp = append(pp, p{0, 0})
		var dir string
		var steps int
		fmt.Sscanf(d, "%1s%d", &dir, &steps)
		fmt.Println(dir, steps)
		switch dir {
		case "R":
			pr := p{
				x: pp[i].x + steps,
				y: pp[i].y,
			}
			pp[i+1] = pr
		case "L":
			pl := p{
				x: pp[i].x - steps,
				y: pp[i].y,
			}
			pp[i+1] = pl
		case "D":
			pd := p{
				x: pp[i].x,
				y: pp[i].y - steps,
			}
			pp[i+1] = pd
		case "U":
			pu := p{
				x: pp[i].x,
				y: pp[i].y + steps,
			}
			pp[i+1] = pu
		}
	}
	//fmt.Println(pp)
	return pp
}
