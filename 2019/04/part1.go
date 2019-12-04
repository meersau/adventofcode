package main

import "fmt"

const inputdata = "356261-846303"

func main() {
	var count int
	for i := 356261; i <= 846303; i++ {
		if neverdec(i) && same(i) {
			count++
		}
	}
	fmt.Println(count)
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
	if e < h {
		return false
	}
	return true
}
