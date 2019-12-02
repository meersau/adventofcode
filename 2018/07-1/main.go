package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type node struct {
	stepname string
}

type nodes []node

func (n nodes) Len() int           { return len(n) }
func (n nodes) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n nodes) Less(i, j int) bool { return n[i].stepname < n[j].stepname }

func (n *node) String() string {
	return n.stepname
}

type graph struct {
	nodes []*node
	edges map[node][]*node
}

func (g *graph) getNach(k string) []node {
	nachfolgers := nodes{}
	//fmt.Println("Edges:", g.edges)
	for re, ed := range g.edges {
		for _, n := range ed {
			if n.stepname == k {
				//fmt.Println("HABE NACHFOLGER", n)
				nachfolgers = append(nachfolgers, re)
			}
		}
	}

	//sort.Sort(nachfolgers)

	return nachfolgers
}
func (g *graph) getNode(k string) *node {
	for _, n := range g.nodes {
		if n.stepname == k {
			return n
		}
	}
	return nil
}

func (g *graph) getEdes(k string) []*node {
	for _, n := range g.nodes {
		if n.stepname == k {
			return g.edges[*n]
		}
	}
	return nil
}

func (g *graph) toposort() {
	nachfolgercount := make(map[node]int)

	queue := nodes{}
	toporder := []node{}
	for _, node := range g.nodes {
		nachfolgercount[*node] = g.getNachfolger(node)
		if nachfolgercount[*node] == 0 {
			queue = append(queue, *node)
		}
	}

	// To store the keys in slice in sorted order
	//var keys []string
	//for k := range nachfolgercount {
	//	fmt.Println("Key for", k)
	//	keys = append(keys, k.stepname)
	//}
	//sort.Strings(keys)
	/* for i := len(keys)/2 - 1; i >= 0; i-- {
		opp := len(keys) - 1 - i
		keys[i], keys[opp] = keys[opp], keys[i]
	} */
	//fmt.Println("keys:", keys)

	//fmt.Println(nachfolgercount)
	fmt.Println(queue)
	fmt.Println("##")
	sort.Sort(queue)
	for len(queue) > 0 {
		fmt.Println("**************")
		var u node
		u, queue = queue[0], queue[1:]
		fmt.Println("Node for toporder", u)

		toporder = append(toporder, u)
		fmt.Println("-----------------------")
		//---> edges nicht nodes
		//for _, n := range g.nodes {

		for _, en := range g.getNach(u.stepname) {
			//fmt.Println("Count nachfolgers", len(g.getNach(u.stepname)))
			//fmt.Printf("Node im innerloop %v", no)
			nachfolgercount[en]--
			if nachfolgercount[en] == 0 {
				queue = append(queue, en)
			}
		}
		//}
		sort.Sort(queue)
		//fmt.Println("NewQu", queue)

	}

	for _, p := range toporder {
		fmt.Printf("%s", p.stepname)
	}
	fmt.Print("\n")

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
	//g.edges[*n1] = append(g.edges[*n1], n2)
	g.edges[*n2] = append(g.edges[*n2], n1)

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
	//gr.Print()
	gr.toposort()
}
