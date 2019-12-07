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
	inputsignal  int
	outputsignal int
	phase        int
}

func NewAmp(phase int) amp {
	return amp{
		phase: phase,
	}
}

func (a amp) RunAmp(input int, prog []int) []int {
	ownmem := make([]int, len(prog))
	outputs := make([]int, 0)
	copy(ownmem, prog)
	//fmt.Println(prog)
	//fmt.Println(ownmem)
	out := intcomp(ownmem, []int{a.phase, input}, outputs)
	//fmt.Println("Outputs:", outputs)
	return out
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
	totest := permutations([]int{0, 1, 2, 3, 4})
	for _, p := range totest {
		fmt.Println(p)
		t := thrust(p[0], p[1], p[2], p[3], p[4], inst)
		if t > max {
			max = t
		}
	}

	fmt.Println(max)
	// 4968420 <- ist falsch */

}

func thrust(p1, p2, p3, p4, p5 int, prog []int) int {
	a1 := NewAmp(p1)
	a2 := NewAmp(p2)
	a3 := NewAmp(p3)
	a4 := NewAmp(p4)
	a5 := NewAmp(p5)
	pro1 := make([]int, len(prog))
	copy(pro1, prog)
	pro2 := make([]int, len(prog))
	copy(pro2, prog)
	pro3 := make([]int, len(prog))
	copy(pro3, prog)
	pro4 := make([]int, len(prog))
	copy(pro4, prog)
	pro5 := make([]int, len(prog))
	copy(pro5, prog)

	o1 := a1.RunAmp(0, pro1)
	o2 := a2.RunAmp(o1[0], pro2)
	o3 := a3.RunAmp(o2[0], pro3)
	o4 := a4.RunAmp(o3[0], pro4)
	o5 := a5.RunAmp(o4[0], pro5)

	return o5[0]
}
