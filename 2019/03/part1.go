package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	weg1 := mw(wire1)
	weg2 := mw(wire2)
	//fmt.Printf("%#v", weg1)
	//pweg(weg1)
	//pweg(weg2)

	l := math.MaxInt64
	d := 0
	for k := range weg1 {
		_, ok := weg2[k]
		if ok {
			d = getdist(p{0, 0}, k)
			//fmt.Println(d)
			//fmt.Println(d, l)
			if d < l {
				l = d
				//fmt.Println("neu", l)

			}
		}
	}
	fmt.Println(l)

}

func getdist(a, b p) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func mw(wire1 []string) map[p]bool {
	weg1 := make(map[p]bool)

	lastx := 0
	lasty := 0
	for i := 0; i < len(wire1); i++ {
		d := wire1[i]
		var dir string
		var steps int
		fmt.Sscanf(d, "%1s%d", &dir, &steps)
		// fmt.Println("Dir", dir, "Steps:", steps)
		switch dir {
		case "R":
			// fmt.Println("R")
			for j := 0; j < steps; j++ {
				lastx++
				weg1[p{lastx, lasty}] = true
			}
		case "L":
			// fmt.Println("L")

			for j := 0; j < steps; j++ {
				lastx--
				weg1[p{lastx, lasty}] = true
			}
		case "D":
			// fmt.Println("D")

			for j := 0; j < steps; j++ {
				lasty--
				weg1[p{lastx, lasty}] = true
			}
		case "U":
			// fmt.Println("U")

			for j := 0; j < steps; j++ {
				lasty++
				weg1[p{lastx, lasty}] = true
			}
		}
	}
	return weg1
}

func pweg(weg map[p]bool) {
	mapp := make([][]string, 30)
	for i := 0; i < 30; i++ {
		mapp[i] = make([]string, 30)
	}
	for i := 0; i < 30; i++ {
		for j := 0; j < 30; j++ {
			mapp[i][j] = "."
		}
	}
	for k := range weg {
		mapp[k.x][k.y] = "-"
	}
	for i := 0; i < 30; i++ {
		for j := 0; j < 30; j++ {
			fmt.Printf("%v", mapp[i][j])
		}
		fmt.Println()
	}
}

/* func makeL(wire []string) map[p]bool {
	weg := make(map[p]bool)

	weg[p{
		x: 0,
		y: 0,
	}] = true

	for i, d := range wire {
		var dir string
		var steps int
		fmt.Sscanf(d, "%1s%d", &dir, &steps)
		fmt.Println(dir, steps)
		switch dir {
		case "R":
			for j := 0; j < steps; j++ {
				k := p{}
			}
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
*/
