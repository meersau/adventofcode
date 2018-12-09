package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var fabric [1000][1000]int
	//var fabric [10][10]int
	re := regexp.MustCompile(`#\d+ @ (\d+),(\d+): (\d+)x(\d+)`)
	//lines := []string{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		matches := re.FindStringSubmatch(s.Text())
		//fmt.Printf("%q", matches)
		var fromleft, fromtop, wide, tall int
		_, err := fmt.Sscanf(matches[1], "%d", &fromleft)
		_, err = fmt.Sscanf(matches[2], "%d", &fromtop)
		_, err = fmt.Sscanf(matches[3], "%d", &wide)
		_, err = fmt.Sscanf(matches[4], "%d", &tall)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("%s\n", s.Text())
		for i := fromleft; i < fromleft+wide; i++ {
			//fmt.Println("y")
			//fmt.Printf("left: %d w: %d", fromleft, wide)
			for j := fromtop; j < fromtop+tall; j++ {
				//fmt.Println("x")
				fabric[i][j]++
			}
		}
	}
	//fmt.Println(fabric)
	var over int
	for i := 0; i < len(fabric); i++ {
		//fmt.Println("")
		for j := 0; j < len(fabric[0]); j++ {
			if fabric[i][j] > 1 {
				over++
			}
			//fmt.Printf("%d", fabric[i][j])
		}
	}
	fmt.Println(over)
}
