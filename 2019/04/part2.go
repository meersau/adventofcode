package main

import (
	"fmt"
	"strconv"
)

const inputdata = "356261-846303"

func main() {
	var count int
	for i := 356261; i <= 846303; i++ {
		if neverdec(i) && isok(i) {
			count++
		}

	}
	fmt.Println(count)
}

func isok(i int) bool {
	si := strconv.Itoa(i)
	last := '0'
	count := 1
	r := false
	for _, c := range si {
		if last == c {
			count++
		} else {
			if count == 2 {
				r = true
			}
			count = 1
		}
		last = c
	}
	if count == 2 {
		return true
	}
	return r
}
func same(i int) bool {
	ht := i / 100000
	zt := (i - (ht * 100000)) / 10000
	t := (i - (ht*100000 + zt*10000)) / 1000
	h := (i - (ht*100000 + zt*10000 + t*1000)) / 100
	z := (i - (ht*100000 + zt*10000 + t*1000 + h*100)) / 10
	e := (i - (ht*100000 + zt*10000 + t*1000 + h*100 + z*10))
	return ht == zt || zt == t || t == h || h == z || z == e
}

func neverdec(i int) bool {
	ht := i / 100000
	zt := (i - (ht * 100000)) / 10000
	t := (i - (ht*100000 + zt*10000)) / 1000
	h := (i - (ht*100000 + zt*10000 + t*1000)) / 100
	z := (i - (ht*100000 + zt*10000 + t*1000 + h*100)) / 10
	e := (i - (ht*100000 + zt*10000 + t*1000 + h*100 + z*10))
	if zt < ht {
		return false
	}
	if t < zt {
		return false
	}
	if h < t {
		return false
	}
	if z < h {
		return false
	}
	if e < z {
		return false
	}
	return true
}
