package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"os"
)

// 65 - 90 A-Z
// 97 - 122 a-z
func main() {
	all, _ := ioutil.ReadFile(os.Args[1])
	allreacted := map[rune]int{}
	u := getUnits(all)
	for u := range u {
		mapf := func(r rune) rune {
			if r == u || r == u+32 {
				return -1
			}
			return r
		}
		newa := bytes.Map(mapf, all)
		reacted := react(newa)
		allreacted[u] = len(reacted)
	}

	min := math.MaxInt32
	for _, v := range allreacted {
		if v < min {
			min = v
		}
	}
	fmt.Println(min)
}

func getUnits(all []byte) map[rune]struct{} {
	units := make(map[rune]struct{})
	for _, r := range all {
		if r > 97 {
			r = r - 32
		}
		units[rune(r)] = struct{}{}
	}
	return units
}

func react(all []byte) []byte {
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
	return all
}
