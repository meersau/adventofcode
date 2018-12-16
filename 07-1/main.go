package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type node struct {
	stepname string
}

func (n *node) String() string {
	return fmt.Sprintf("%s", n.stepname)
}

type graph struct {
	nodes []*node
	edges map[node][]*node
}

func (g *graph) toposort() {
	visited := make(map[node]bool)
	for _, n := range g.nodes {
		visited[*n] = false
	}
	fmt.Println(visited)
	stack := make([]node, 0)
	for _, n := range g.nodes {
		if visited[*n] == false {
			g.toprec(n, visited, stack)
		}
	}
	fmt.Println(visited)

	fmt.Println("--", stack)
}

func (g *graph) toprec(n *node, visited map[node]bool, stack []node) {
	visited[*n] = true
	for _, no := range g.nodes {
		if visited[*no] == false {
			g.toprec(no, visited, stack)
		}
	}
	fmt.Println("append")
	fmt.Println(n)
	fmt.Println("++*", stack)
	stack = append(stack, *n)
}

func (g *graph) getNachfolger(n *node) int {
	return len(g.edges[*n])
}
func (g *graph) hasNode(n *node) bool {
	for _, no := range g.nodes {
		if no.stepname == n.stepname {
			return true
		}
	}
	return false
}
func (g *graph) AddNode(n *node) {
	if !g.hasNode(n) {
		g.nodes = append(g.nodes, n)
	}
}

func (g *graph) AddEdge(n1, n2 *node) {
	if g.edges == nil {
		g.edges = make(map[node][]*node)
	}
	g.edges[*n1] = append(g.edges[*n1], n2)
	// g.edges[*n2] = append(g.edges[*n2], n1)

}

func (g *graph) Print() {
	s := ""
	for i := 0; i < len(g.nodes); i++ {
		s += g.nodes[i].String() + " -> "
		near := g.edges[*g.nodes[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}
		//fmt.Printf("%d", g.getNachfolger(g.nodes[i]))
		s += "\n"
	}
	fmt.Println(s)
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	var gr graph
	for s.Scan() {
		n1 := &node{}
		n2 := &node{}
		fmt.Sscanf(s.Text(), "Step %s must be finished before step %s can begin.", &n1.stepname, &n2.stepname)
		gr.AddNode(n1)
		gr.AddNode(n2)
		gr.AddEdge(n1, n2)
	}
	gr.Print()
	gr.toposort()
}
