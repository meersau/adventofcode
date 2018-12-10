package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

// 65 - 90 A-Z
// 97 - 122 a-z
func main() {
	all, _ := ioutil.ReadFile(os.Args[1])
	all = bytes.TrimRight(all, "\n")
	var pos int
	for {
		if pos+1 >= len(all) {
			break
		}
		b1 := all[pos]
		b2 := all[pos+1]
		//fmt.Printf("Check %s %s\n", string(b1), string(b2))
		diff := int(b1) - int(b2)
		//fmt.Println(diff)
		if diff == 32 || diff == -32 {
			//fmt.Println("double", string(all[pos]), string(all[pos+1]))
			all = append(all[:pos], all[pos+2:]...)
			pos = 0
			continue
		}
		pos++

	}
	fmt.Println(string(all), len(all))
}
