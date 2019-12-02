package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type point struct {
	n   int
	x   int
	y   int
	inf bool
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	//var grid[][]
	s := bufio.NewScanner(f)
	//var n int
	coords := []point{}
	maxy := 0
	maxx := 0
	//closest := 0
	for s.Scan() {
		var x, y int
		_, e := fmt.Sscanf(s.Text(), "%d, %d", &x, &y)
		if e != nil {
			log.Fatal(e)
		}
		coords = append(coords, point{0, x, y, false})
		if x > maxx {
			maxx = x
		}
		if y > maxy {
			maxy = y
		}
		//n++
	}
	grid := make([][]int, maxx)
	for i := 0; i < maxx; i++ {
		grid[i] = make([]int, maxy)
	}
	for y := 0; y < maxy; y++ {
		for x := 0; x < maxx; x++ {
			min := math.MaxInt32
			minn := -1
			nocount := false
			for i, c := range coords {
				dist := getdist(point{0, x, y, false}, c)
				if dist < min {
					min = dist
					minn = i
					nocount = false
				} else if dist == min {
					nocount = true
				}
			}

			if x == 0 || y == 0 || x == maxx || y == maxy {
				coords[minn].inf = true

			}

			if !nocount {
				coords[minn].n++
			}
		}
	}
	var max int
	for _, c := range coords {

		if c.inf == true {
			continue
		}
		fmt.Println(c)
		if c.n > max {
			max = c.n
		}
	}
	fmt.Println(max)
}

func getdist(a, b point) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
