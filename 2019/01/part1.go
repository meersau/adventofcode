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
	fmt.Printf("m %d f: %d\n", m, e)
	return e
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
		su = su + e
	}
	fmt.Println(su)
}
