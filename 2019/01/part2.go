package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func bm(m int) int {
	e := m / 3
	// runden
	e = e - 2
	return e
}

func ff(e int) int {
	var sf int
	f := bm(e)
	fmt.Println("+ ", f)
	for f >= 1 {
		fmt.Println("+ ", f)
		sf = f + sf
		f = bm(f)
	}

	return sf
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var su int

	s := bufio.NewScanner(f)
	for s.Scan() {
		m, _ := strconv.Atoi(s.Text())
		e := bm(m)
		fmt.Println("fm", e)
		sff := ff(e)
		su = su + e + sff
	}
	fmt.Println(su)
}
