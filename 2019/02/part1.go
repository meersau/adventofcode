package main_

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type ops func(int, int, int, []int)

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

	o := make([]ops, 3)
	o[1] = func(i, j, k int, z []int) {
		z[z[k]] = z[z[i]] + z[z[j]]
	}
	o[2] = func(i, j, k int, z []int) {
		z[z[k]] = z[z[i]] * z[z[j]]
	}

	inst[1] = 12
	inst[2] = 2
	for g := 0; g < len(inst); g = g + 4 {
		if inst[g] == 99 {
			fmt.Println(inst[0])

			break
		}
		o[inst[g]](g+1, g+2, g+3, inst)
	}

}
