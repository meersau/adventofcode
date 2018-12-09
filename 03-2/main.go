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

	var fabric [1000][1000][]int
	//var fabric [10][10][]int
	ids := []int{}
	over := map[int]bool{}
	re := regexp.MustCompile(`#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)
	//lines := []string{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		matches := re.FindStringSubmatch(s.Text())
		//fmt.Printf("%q", matches)
		var id, fromleft, fromtop, wide, tall int
		_, err := fmt.Sscanf(matches[1], "%d", &id)
		_, err = fmt.Sscanf(matches[2], "%d", &fromleft)
		_, err = fmt.Sscanf(matches[3], "%d", &fromtop)
		_, err = fmt.Sscanf(matches[4], "%d", &wide)
		_, err = fmt.Sscanf(matches[5], "%d", &tall)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("%s\n", s.Text())
		ids = append(ids, id)
		over[id] = false
		for i := fromleft; i < fromleft+wide; i++ {
			//fmt.Println("y")
			//fmt.Printf("left: %d w: %d", fromleft, wide)
			for j := fromtop; j < fromtop+tall; j++ {
				//fmt.Println("x")
				if fabric[i][j] == nil {
					fabric[i][j] = make([]int, 0)
				}
				fabric[i][j] = append(fabric[i][j], id)
			}
		}
	}
	//fmt.Println(fabric)

	for i := 0; i < len(fabric); i++ {
		for j := 0; j < len(fabric[0]); j++ {
			if len(fabric[i][j]) > 1 {
				for _, id := range fabric[i][j] {
					over[id] = true
				}
			}
		}
	}
	for i, v := range over {
		if !v {
			fmt.Println(i)
		}
	}
	//fmt.Println(over)
}
