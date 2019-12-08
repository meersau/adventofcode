package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type amp struct {
	inputsignal  int
	outputsignal int
	phase        int
	inchan       chan int
	outchan      chan int
	memory       []int
}

func NewAmp(phase int, mem []int) amp {
	in := make(chan int)
	out := make(chan int)
	mymem := make([]int, len(mem))
	copy(mymem, mem)
	return amp{
		phase:   phase,
		inchan:  in,
		outchan: out,
		memory:  mymem,
	}
}

func (a amp) RunAmp() {
	intcomp(a.memory, a.inchan, a.outchan)
}

func (a amp) SendInput(in int) {
	a.inchan <- in
}

func (a amp) GetOut() int {
	return <-a.outchan
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

	/* b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	} */
	b := []byte("3,9,8,9,10,9,4,9,99,-1,8")
	i := bytes.Split(b, []byte{','})
	inst := make([]int, 0)
	for _, z := range i {
		i, _ := strconv.Atoi(string(z))
		inst = append(inst, i)
	}

	in := make(chan int)
	out := make(chan int)

	go intcomp(inst, in, out)

	in <- 8
	fmt.Println(<-out)

	/* var max int
	totest := permutations([]int{0, 1, 2, 3, 4})
	for _, p := range totest {
		fmt.Println(p)
		t := thrust(p[0], p[1], p[2], p[3], p[4], inst)
		if t > max {
			max = t
		}
	}
	fmt.Println(max)
	*/ // 4968420 <- ist falsch */

}

func thrust(p1, p2, p3, p4, p5, startsig int, prog []int) {

	a1 := NewAmp(p1, prog)
	a2 := NewAmp(p2, prog)
	a3 := NewAmp(p3, prog)
	a4 := NewAmp(p4, prog)
	a5 := NewAmp(p5, prog)

	a1.RunAmp()
	a2.RunAmp()
	a3.RunAmp()
	a4.RunAmp()
	a5.RunAmp()

}
