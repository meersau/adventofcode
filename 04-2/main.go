package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	re := regexp.MustCompile(`\[\d{4}-\d{2}-\d{2} \d{2}:(\d{2})\] (.*)`)
	var id int
	var startmin int
	whichmin := map[int][59]int{}
	guardsleep := map[int]int{}
	for s.Scan() {
		matches := re.FindStringSubmatch(s.Text())

		minutes, _ := strconv.Atoi(matches[1])
		action := matches[2]

		if strings.Contains(action, "begins shift") {
			fmt.Sscanf(matches[2], "Guard #%d begins shift", &id)
			continue
		}

		if strings.Contains(action, "falls asleep") {
			startmin = minutes
			continue
		}

		if strings.Contains(action, "wakes up") {
			guardsleep[id] = guardsleep[id] + (minutes - startmin)
			min := whichmin[id]
			for i := startmin; i < minutes; i++ {
				min[i]++
			}
			whichmin[id] = min

		}
	}

	// fmt.Println(whichmin)
	var max, mid, cix int

	for id, mins := range whichmin {
		for ix, c := range mins {
			if c > max {
				max = c
				mid = id
				cix = ix
			}
		}
	}
	fmt.Println(mid * cix)
}
