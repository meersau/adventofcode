package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type instruction struct {
	sorted bool
	steps  *step
}

type step struct {
	name   string
	before *step
	//next   *step
}

func (s *step) String() string {
	if s.before == nil {
		return fmt.Sprintf("Step %s, vor diesem muss nix\n", s.name)
	}
	return fmt.Sprintf("Step %s, vor diesem muss %s\n", s.name, s.before.name)
}
func (s *step) Next() *step {
	if s == nil {
		return nil
	}
	return s.before
}

func (s *step) Before() *step {
	if s == nil {
		return nil
	}
	return s.before
}

type steps struct {
	start *step
}

func (s *steps) Insert(newstep *step) {
	fmt.Println(newstep)
	if s.start == nil {
		s.start = newstep
		fmt.Printf("Add %s as Start\n", s.start.name)
		return
	}
	if newstep.before.name == s.start.name {
		fmt.Printf("Neuer Step %s kommt vor %s", newstep.before.name, s.start.name)
		if s.start.before != nil {
			fmt.Println("Gibt schon einen :-(")
			tempbefore := s.start.before
			s.start.before = newstep
			newstep.before = tempbefore
		}

		return
	}
	stepbefore := s.Step(newstep.before.name)
	if stepbefore == nil {
		fmt.Printf("Kein Step %s gefunden: %v\n", newstep.before.name, newstep)
		tempstart := s.start
		s.start = nil
		s.start = newstep
		fmt.Printf("Neu Start Step %s\n", s.start.name)
		s.start.before = tempstart
		return
	}
	fmt.Printf("ACHTUNG ÜBERSCHREIBE %v mit %v\n", stepbefore.before, newstep)
	stepbefore.before = newstep
}
func (ss *steps) Step(name string) *step {
	if ss == nil {
		return nil
	}
	if ss.start.name == name {
		return ss.start
	}
	cst := ss.start.Next()
	for cst != nil {
		//fmt.Println("Found next", cst)
		// if cst.name == name || cst.before.name == name || cst.next.name == name {
		if cst.name == name {
			return cst
		}
		cst = cst.before
	}
	return nil
}

func (ss *steps) PrintAll() {
	cst := ss.start
	for cst != nil {
		fmt.Printf("%s ", cst.name)
		fmt.Printf("%v", cst)
		cst = cst.Next()
	}
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	var ss steps
	for s.Scan() {
		var mustfin, canbegin string
		fmt.Sscanf(s.Text(), "Step %s must be finished before step %s can begin.", &mustfin, &canbegin)
		ss.Insert(&step{
			name: mustfin,
			before: &step{
				name: canbegin,
			},
		})
	}
	ss.PrintAll()
}
