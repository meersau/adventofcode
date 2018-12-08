package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	a    int
	freq map[int]int
)

func main() {

	freq = make(map[int]int)
	freq[0] = 0
	var c int
	for {
		loop()
		c++
		fmt.Println(c)
	}
}

func loop() {
	f, _ := os.Open(os.Args[1])
	s := bufio.NewScanner(f)
	for s.Scan() {
		l := s.Text()
		i, e := strconv.Atoi(l)
		if e != nil {
			log.Fatal(e)
		}
		a = a + i
		freq[a]++
		//fmt.Printf("%v\n", freq)
		//fmt.Println(a)
		c, ok := freq[a]
		if ok && c == 2 {
			fmt.Println(a)
			os.Exit(0)
		}
		//fmt.Println(a)
	}
	f.Close()
}
