package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	lines := []string{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	//sort.Strings(lines)
	for i, s := range lines { // jedes wort
		for j, is := range lines { // mit jedem wort
			if i == j { // gleiche wörter ignorieren
				continue
			}
			//fmt.Printf("Vergleiche %s mit %s\n", s, is)
			var count, ug int
			for h, c := range s { // wort 1
				if c != rune(is[h]) {
					ug = h
					count++
				}
			}
			if count == 1 {

				fmt.Printf("%s %s %s\n", s, is, fmt.Sprintf("%s%s", s[:ug], s[ug+1:]))
				os.Exit(0)
			}
		}
	}
}
