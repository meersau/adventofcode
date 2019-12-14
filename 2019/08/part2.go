package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// 0 is black, 1 is white, and 2 is transparent
// first layer in front and the last layer in back
func main() {

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	bf := bufio.NewReader(f)

	wide := 25
	tall := 6
	//wide := 2
	//tall := 2
	//layermit := 0

	anzahllayer := bf.Size() / (tall + wide)
	fmt.Println("Anzal lay", anzahllayer)
	image := make([][]int, tall)
	for i := 0; i < tall; i++ {
		image[i] = make([]int, wide)
	}
	//fmt.Println(image)

	for layer := 0; layer < anzahllayer; layer++ {
		/* 		c0 := 0
		   		einser := 0
		   		zweier := 0
		*/
		//lay := make([][]int, wide)
		for t := 0; t < tall; t++ {

			for w := 0; w < wide; w++ {

				by, _ := bf.ReadByte()
				digitstr := string(by)
				digit, _ := strconv.Atoi(digitstr)
				if layer == 0 {
					image[t][w] = digit
					continue
				}
				if image[t][w] == 0 || image[t][w] == 1 {
					continue
				}

				image[t][w] = digit

			}
		}
		//image = append(image, lay)
		//fmt.Printf("Layer %d hat Nuller: %d das sind Einser: %d und Zweier: %d\n", layer, c0, einser, zweier)

	}
	//fmt.Println(image)
	for _, r := range image {
		for _, w := range r {
			if w == 0 {
				fmt.Printf("\033[1;30m%s\033[0m", "X") // black
			} else {
				fmt.Printf("\033[1;37m%s\033[0m", "x")
			}
		}
		fmt.Println()
	}
	//fmt.Println(image)
}

/* var (
	Black   = Color("\033[1;30m%s\033[0m")
	Red     = Color("\033[1;31m%s\033[0m")
	Green   = Color("\033[1;32m%s\033[0m")
	Yellow  = Color("\033[1;33m%s\033[0m")
	Purple  = Color("\033[1;34m%s\033[0m")
	Magenta = Color("\033[1;35m%s\033[0m")
	Teal    = Color("\033[1;36m%s\033[0m")
	White   = Color("\033[1;37m%s\033[0m")
)
*/
