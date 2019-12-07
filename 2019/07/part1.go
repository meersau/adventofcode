package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
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
	fmt.Println(prog)
	fmt.Println(ownmem)
	out := intcomp(ownmem, []int{a.phase, input}, outputs)
	fmt.Println("Outputs:", outputs)
	return out
}

func getprog(file string) []int {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	i := bytes.Split(b, []byte{','})
	inst := make([]int, 0)
	for _, z := range i {
		i, _ := strconv.Atoi(string(z))
		inst = append(inst, i)
	}
	return inst

}

func main() {

	// prog := getprog(os.Args[0])
	// b := []byte("3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0")
	b := []byte("3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0")
	i := bytes.Split(b, []byte{','})
	prog := make([]int, 0)
	for _, z := range i {
		i, _ := strconv.Atoi(string(z))
		prog = append(prog, i)
	}
	a1 := NewAmp(0)
	a2 := NewAmp(1)
	a3 := NewAmp(2)
	a4 := NewAmp(3)
	a5 := NewAmp(4)

	o1 := a1.RunAmp(0, prog)
	o2 := a2.RunAmp(o1[0], prog)
	o3 := a3.RunAmp(o2[0], prog)
	o4 := a4.RunAmp(o3[0], prog)
	o5 := a5.RunAmp(o4[0], prog)

	fmt.Println(o5[0])
}
