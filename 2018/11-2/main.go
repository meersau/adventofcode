package main

import (
	"fmt"
	"os"
	"strconv"
)

const gc = 300

func main() {
	var grid [gc][gc]int64
	sn, _ := strconv.Atoi(os.Args[1])

	for i := 0; i < gc; i++ {
		for j := 0; j < gc; j++ {
			x := i + 1
			y := j + 1
			rackid := x + 10
			pl := (rackid * y) + sn
			hunder := pl * rackid
			h := (hunder / 100) % 10

			grid[i][j] = int64(h - 5)
		}
	}

	var max, sum int64

	var x, y, maxx, maxy, s int
	for i := 0; i < gc; i++ {

		x, y, sum = sumsize(grid, i)
		if sum > max {
			max = sum
			maxx = x
			maxy = y
			s = i
			fmt.Println("Neuer max:", x, y, s)
		}
	}
	fmt.Printf("%d,%d,%d\n", maxx+1, maxy+1, s)
}

func sumsize(grid [gc][gc]int64, size int) (int, int, int64) {
	var max int64
	var maxx, maxy int
	for i := 0; i < gc-size; i++ {
		for j := 0; j < gc-size; j++ {

			var sum int64
			for f := 0; f < size; f++ {
				for g := 0; g < size; g++ {
					//fmt.Println("f:", f, g)
					sum += grid[i+f][j+g]
				}
			}
			if sum > max {
				max = sum
				maxx = i
				maxy = j
			}
		}
	}
	return maxx, maxy, max
}
