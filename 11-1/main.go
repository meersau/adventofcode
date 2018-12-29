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
	var max int64
	var maxx, maxy int
	for i := 0; i < gc-3; i++ {
		for j := 0; j < gc-3; j++ {

			var sum int64
			for f := 0; f < 3; f++ {
				for g := 0; g < 3; g++ {
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

	fmt.Printf("%d,%d\n", maxx+1, maxy+1)
}
