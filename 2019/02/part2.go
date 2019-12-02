package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type ops func(int, int, int, []int)

var o []ops

func main() {
	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	i := bytes.Split(b, []byte{','})
	inst := make([]int, 0)
	for _, z := range i {
		i, _ := strconv.Atoi(string(z))
		inst = append(inst, i)
	}

	o = make([]ops, 3)
	o[1] = func(i, j, k int, z []int) {
		z[z[k]] = z[z[i]] + z[z[j]]
	}
	o[2] = func(i, j, k int, z []int) {
		z[z[k]] = z[z[i]] * z[z[j]]
	}

	for n := 0; n < 100; n++ {
		for v := 0; v < 100; v++ {
			if testnv(inst, n, v) == 19690720 {
				fmt.Println(100*n + v)
				break
			}
		}
	}
}

func testnv(start []int, noun, verb int) int {
	c := make([]int, len(start))
	copy(c, start)
	c[1] = noun
	c[2] = verb
	for g := 0; g < len(c); g = g + 4 {
		if c[g] == 99 {
			return c[0]
			break
		}
		o[c[g]](g+1, g+2, g+3, c)
	}
	return 0
}
