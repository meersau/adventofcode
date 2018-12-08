package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var zweier, dreier int

	f, _ := os.Open(os.Args[1])
	s := bufio.NewScanner(f)
	for s.Scan() {
		str := s.Text()
		var found2, found3 bool
		for _, c := range strings.Split(str, "") {
			i := strings.Count(str, c)
			if i == 3 && !found3 {
				found3 = true
				dreier++
			}
			if i == 2 && !found2 {
				found2 = true
				zweier++
			}
		}
	}
	fmt.Println(zweier * dreier)
}
