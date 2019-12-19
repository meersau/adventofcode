package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type amp struct {
	phase  int
	prog   []int
	input  chan int
	output chan int
	n      string
}

func NewAmp(phase int, prog []int, name string) amp {
	ownprog := make([]int, len(prog))
	copy(ownprog, prog)
	return amp{
		phase: phase,
		prog:  ownprog,
		n:     name}
}

func (a amp) RunAmp(input chan int) chan int {
	a.output = make(chan int)
	a.input = input
	go intcomp(a.prog, a.input, a.output, a.n)
	a.input <- a.phase
	return a.output
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

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

	var max int
	totest := permutations([]int{5, 6, 7, 8, 9})
	for _, p := range totest {
		t := thrust(p[0], p[1], p[2], p[3], p[4], inst)
		//t := thrust(5, 6, 7, 8, 9, inst)
		if t > max {
			max = t
		}
	}
	fmt.Println(max)
	// 4968420 <- ist falsch */
}

func thrust(p1, p2, p3, p4, p5 int, prog []int) int {
	a1 := NewAmp(p1, prog, "A1")
	a2 := NewAmp(p2, prog, "A2")
	a3 := NewAmp(p3, prog, "A3")
	a4 := NewAmp(p4, prog, "A4")
	a5 := NewAmp(p5, prog, "A5")

	input := make(chan int, 1)
	out1 := a1.RunAmp(input)
	o2 := a2.RunAmp(out1)
	o3 := a3.RunAmp(o2)
	o4 := a4.RunAmp(o3)
	o5 := a5.RunAmp(o4)
	input <- 0
	var rt int
	for {
		j, more := <-o5
		if more {
			rt = j
			input <- rt
		} else {
			close(input)
			return rt
		}
	}

	/* for feed := range o5 {
		rt = feed
		input <- rt
	}
	fmt.Println("Range fertig")
	close(input) */
	//return rt
}
