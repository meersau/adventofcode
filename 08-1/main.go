package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type node struct {
	next   *node
	childs int
	pos    int
	mp     int
	meta   []int
}

func (n *node) String() string {
	return fmt.Sprintf("Node: %d - Meta: %v", n.pos, n.meta)
}

type tree struct {
	root *node
	c    int
}

func (t *tree) pr() {
	fmt.Println("C: ", t.c)
	cur := t.root

	for cur.next != nil {
		fmt.Printf("%v\n", cur)
		cur = cur.next
	}
	fmt.Printf("%v\n", cur)

}

func (t *tree) insertNode(n *node) {

	if t.root == nil {
		t.root = n
		t.c++
		return
	}
	cur := t.root
	for cur.next != nil {
		cur = cur.next
	}

	cur.next = n
	t.c++
	//fmt.Printf("New cir.netx: %v\n", n)
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	b := make([]int, 0)
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		i, _ := strconv.Atoi(s.Text())
		b = append(b, i)
	}
	t := &tree{}
	parsenodes(t, b, 0)

	fmt.Println("################")
	t.pr()
	//fmt.Printf("%v\n", t)
}

// 0 1 2 3 4  5  6  7 8 9 10 11 12 13 14 15
// 2 3 0 3 10 11 12 1 1 0 1  99 2  1  1  2
// A----------------------------------
//     B----------- C-----------
//				        D-----

func parsenodes(t *tree, b []int, pos int) int {
	n := &node{pos: pos, childs: b[pos], mp: b[pos+1]}
	fmt.Printf("--> Node %d hat %d Kinder und %d Meta\n", pos, b[pos], b[pos+1])
	pos = pos + 2
	for i := 0; i < n.childs; i++ {
		pos = parsenodes(t, b, pos)
		fmt.Println("Return parse: ", pos)
	}

	n.meta = make([]int, 0)
	if n.mp > 0 {
		for i := 0; i < n.mp; i++ {
			n.meta = append(n.meta, b[pos+i])
		}
	}
	t.insertNode(n)
	fmt.Printf("Insert Node pos: %d: %+v\n", pos, n)
	return pos + len(n.meta)
}
