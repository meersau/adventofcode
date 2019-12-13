package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	bf := bufio.NewReader(f)

	wide := 25
	tall := 6
	min := math.MaxInt64
	ml := 0
	//layermit := 0
	le := 0
	lz := 0
	// for l := 0; l < len(b)/31; l = l + wide + tall {

	anzahllayer := bf.Size() / (tall + wide)
	// image := make([][][]int, anzahllayer)
	for layer := 0; layer < anzahllayer; layer++ {
		c0 := 0
		einser := 0
		zweier := 0
		for w := 0; w < wide; w++ {
			for t := 0; t < tall; t++ {
				by, _ := bf.ReadByte()
				digitstr := string(by)
				//fmt.Println("index:", layer+w+t, b[layer+w+t])
				digit, _ := strconv.Atoi(digitstr)
				if digit == 0 {
					c0++
				}
				if digit == 1 {
					einser++
				}
				if digit == 2 {
					zweier++
				}
			}
		}
		//fmt.Printf("Layer %d hat Nuller: %d das sind Einser: %d und Zweier: %d\n", layer, c0, einser, zweier)
		if c0 < min {
			min = c0
			le = einser
			lz = zweier
		}
	}

	fmt.Println(ml, le, lz)
	fmt.Println(le * lz)

}
