package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type lineentry struct {
	timeentry time.Time
	guard     int
	asleep    bool
	start     bool
}

type timeline []lineentry

// Forward request for length
func (p timeline) Len() int {
	return len(p)
}

// Define compare
func (p timeline) Less(i, j int) bool {
	return p[i].timeentry.Before(p[j].timeentry)
}

// Define swap over an array
func (p timeline) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

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

	fmt.Println(guardsleep)
	var max, mid int

	for k, v := range guardsleep {
		if v > max {
			max = v
			mid = k
		}
	}
	fmt.Println(mid)
	max = 0
	min := 0
	for k, v := range whichmin[mid] {
		if v > max {
			max = v
			min = k
		}
	}
	fmt.Println(mid * min)
}
