package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	//b := []byte("109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99")
	// Large Number test
	//b := []byte("1102,34915192,34915192,7,4,7,99,0")
	//b := []byte("104,1125899906842624,99")

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
	in := make(chan int)
	out := make(chan int)
	go intcomp(inst, in, out, "A1")
	in <- 2
	oi := make([]int, 0)
	for o := range out {
		oi = append(oi, o)
	}
	close(in)
	fmt.Println(oi)
}
