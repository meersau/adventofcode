package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
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
	re := regexp.MustCompile(`\[(.*)\] (.*)`)
	tl := timeline{}
	for s.Scan() {
		matches := re.FindStringSubmatch(s.Text())
		t, err := time.Parse("2006-01-02 15:04", matches[1])
		if err != nil {
			log.Fatal(err)
		}
		if matches[2] == "falls asleep" {
			tl = append(tl, lineentry{t, 0, true, false})
			continue
		}
		if matches[2] == "wakes up" {
			tl = append(tl, lineentry{t, 0, false, false})
			continue
		}
		var i int
		fmt.Sscanf(matches[2], "Guard #%d begins shift", &i)
		tl = append(tl, lineentry{t, i, false, true})
	}
	sort.Sort(tl)

	tasleep := map[int]int{}
	min := 0
	g := 0
	var startt, lastasleep time.Time
	for _, e := range tl {
		//fmt.Println("Schleife: ", in)
		if e.start {
			fmt.Printf("Guard %d start at %v\n", e.guard, e.timeentry)
			min = 0
			startt = e.timeentry
			g = e.guard
			continue
		}
		if e.asleep {
			lastasleep = e.timeentry
			continue
		}
		if e.asleep == false {
			min = min + int(e.timeentry.Sub(startt).Minutes())
			fmt.Println(e.timeentry.Sub(startt).Minutes())
			tasleep[g] = tasleep[g] + min
			startt = e.timeentry
			continue
		}

	}

	fmt.Println(tasleep)

	// mg := 0
	// for in, i := range tl {
	// 	if in == 0 && i.guard == 0 {
	// 		fmt.Println("sollte nicht passieren")
	// 		os.Exit(1)
	// 	}

	// 	if in == 0 {
	// 		mg = i.guard
	// 		continue
	// 	}
	// 	if i.guard != 0 && mg != i.guard {
	// 		mg = i.guard
	// 		continue
	// 	}
	// 	i.guard = mg

	// }

}
