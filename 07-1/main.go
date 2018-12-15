package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type step struct {
	name string
	next *step
}

func (s *step) String() string {
	if s.next == nil {
		return fmt.Sprintf("Step %s, kommr keiner mehr", s.name)
	}
	return fmt.Sprintf("Step %s, hat als nächstes %s", s.name, s.next.name)
}
func (s *step) Next() *step {
	if s == nil {
		return nil
	}
	return s.next
}

type steps struct {
	start *step
}

func (s *steps) Insert(newstep *step) {

	//fmt.Println(newstep)
	if s.start == nil {
		fmt.Println("START")
		s.start = newstep
		return
	}

	if s.start.name == newstep.name {
		fmt.Println("Muss START ersetzen")
		if s.start.next == nil {
			s.start.next = newstep.next
			return
		}
		if s.start.next.name < newstep.next.name {
			fmt.Println(s.start.next.name, "<", newstep.next.name)
			if s.start.next.next == nil {
				fmt.Printf("%s hat keinen nächsten\n", s.start.next.name)
				s.start.next.next = newstep.next
				return
			}
			newstep.next.next = s.start.next.next
			s.start.next.next = newstep.next

		} else {
			newstep.next.next = s.start.next
			s.start.next = newstep.next

		}
		return
	}
	if newstep.next.name == s.start.name {
		fmt.Printf("START GLEICH NEW\n")
		return
	}

	foundstep := s.Step(newstep.name)
	if foundstep == nil {
		fmt.Println("Not yte implemented")
		return
	}
	if foundstep.next == nil {
		fmt.Println("Gefundener Step hat keinen Nächsten")
		fmt.Println("Neu: ", newstep)
		fmt.Println("Gefunden: ", foundstep)
		gibtesdennächsten := s.Step(newstep.next.name)
		fmt.Println("Nächster: ", gibtesdennächsten)
		vorgaenger := s.VorStep(gibtesdennächsten.name)
		fmt.Println("VORGAENGER", vorgaenger)
		return
	}

	if foundstep.next.name == newstep.next.name {
		return
	}

	if foundstep.next.name < newstep.next.name {
		fmt.Println(foundstep.next.name, "<", newstep.next.name)
		if foundstep.next.next == nil {
			fmt.Printf("%s hat keinen nächsten\n", foundstep.next.name)
			foundstep.next.next = newstep.next
			return
		}
		newstep.next.next = foundstep.next.next
		foundstep.next.next = newstep.next
		return
	}

	if foundstep.next.name > newstep.next.name {
		fmt.Println(foundstep.next.name, ">", newstep.next.name)
		newstep.next.next = foundstep.next
		foundstep.next = newstep.next
		return

	}

	fmt.Println("der keine ahnung was fall")
	fmt.Println("Neustep: ", newstep)
	fmt.Println("found: ", foundstep)
	return

}
func (ss *steps) VorStep(name string) *step {
	if ss == nil {
		return nil
	}
	if ss.start.name == name {
		fmt.Println("START HAT KEINE VORGÄNGER")
		return nil
	}
	cst := ss.start.Next()
	var vorg *step
	for cst != nil {
		if cst.name == name {
			return vorg
		}
		vorg = cst
		cst = cst.next
	}
	return nil
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
		if cst.name == name {
			return cst
		}
		cst = cst.next
	}
	return nil
}

func (ss *steps) PrintAll() {
	cst := ss.start
	for cst != nil {
		fmt.Printf("%s ", cst.name)
		//fmt.Printf("%v", cst)
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
			next: &step{
				name: canbegin,
			},
		})
		fmt.Println("==========================================")
		ss.PrintAll()
		fmt.Println()
		fmt.Println("==========================================")
	}
	ss.PrintAll()
}
