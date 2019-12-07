package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	//b := []byte("3,9,8,9,10,9,4,9,99,-1,8")
	//b := []byte("3,3,1108,-1,8,3,4,3,99")
	// less
	//b := []byte("3,9,7,9,10,9,4,9,99,-1,8")
	// b := []byte("3,3,1107,-1,8,3,4,3,99")
	// jump test
	//	Here are some jump tests that take an input, then output 0 if the input was zero or 1 if the input was non-zero:

	// b := []byte("3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9")
	//b := []byte("3,3,1105,-1,9,1101,0,0,12,4,12,99,1")
	// b := []byte("3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99")

	i := bytes.Split(b, []byte{','})
	inst := make([]int, 0)
	for _, z := range i {
		i, _ := strconv.Atoi(string(z))
		inst = append(inst, i)
	}

	intcomp(inst, []int{5})

}
