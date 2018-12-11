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
		fmt.Sscanf(s.Text(), "%d, %d", &x, &y)
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

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			min := math.MaxInt32
			minn := 0
			for i, c := range coords {
				dist := getdist(c, point{0, x, y, false})
				if dist < min {
					min = dist
					minn = i

				}
			}
			if x == 0 || y == 0 || x == len(grid) || y == len(grid[x]) {
				coords[minn].inf = true
			}
			coords[minn].n++
		}
	}
	var max int
	for _, c := range coords {
		if c.inf == true {
			continue
		}
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
		return x * -1
	}
	return x
}
