package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type p struct {
	x, y int
}

type robot struct {
	panel    map[p]int
	curpan   *p
	curdir   string
	curcol   int
	inchan   chan int
	outchan  chan int
	topaint  int
	turn     int
	countpan int
}

func makenewdir(curdir string, newdir int) string {
	switch curdir {
	case "U":
		if newdir == 0 { // left 90 degrees
			return "L"
		}
		// (1) right 90 degrees
		return "R"

	case "R":
		if newdir == 0 { // left 90 degrees
			return "U"
		}
		// (1) right 90 degrees
		return "D"
	case "L":
		if newdir == 0 { // left 90 degrees
			return "D"
		} else { // (1) right 90 degrees
			return "U"
		}
	case "D":
		if newdir == 0 { // left 90 degrees
			return "R"
		} else { // (1) right 90 degrees
			return "L"
		}
	}
	return ""
}
func makestep(pan *p, dir string) {
	switch dir {
	case "U":
		pan.x = pan.x + 1
	case "L":
		pan.y = pan.y - 1
	case "R":
		pan.y = pan.y + 1
	case "D":
		pan.x = pan.x - 1
	}
}

func (r *robot) nextstep() {
	r.topaint = <-r.outchan
	r.turn = <-r.outchan

	if r.curcol == r.topaint {
		fmt.Println("Male gleiche Farbe")
	} else {
		fmt.Println("Male neue farbe")
		r.panel[*r.curpan] = r.topaint
	}
	r.curdir = makenewdir(r.curdir, r.turn)
	makestep(r.curpan, r.curdir)
	fmt.Println(r.curpan)
	r.inchan <- 0
}
func (r *robot) step(out chan int, in chan int) {
	// First, it will output a value indicating the
	// color to paint the panel the robot is over:
	// 0 means to paint the panel black, and
	// 1 means to paint the panel white.

	// Second, it will output a value indicating the direction the robot should turn:
	// 0 means it should turn left 90 degrees, and 1 means it should turn right 90 degrees.
	// one step forward
}

func main() {
	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	i := bytes.Split(b, []byte{','})
	inst := make([]int, 0)
	for _, z := range i {
		i, _ := strconv.Atoi(string(z))
		inst = append(inst, i)
	}

	in := make(chan int)
	out := make(chan int)
	done := make(chan bool)

	rob := &robot{
		panel:   make(map[p]int),
		curpan:  &p{0, 0},
		curcol:  0,
		inchan:  in,
		outchan: out,
	}

	go intcomp(inst, in, out, "AX", done)
	// start
	in <- 0
	for {
		select {
		case <-done:
			fmt.Println(rob.countpan)
			break
		default:
			rob.nextstep()
		}
	}

	close(in)
}
