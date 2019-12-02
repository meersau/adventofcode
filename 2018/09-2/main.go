package main

import (
	"container/ring"
	"fmt"
	"log"
	"os"
)

// input: 428 players; last marble is worth 72061 points
// test:
// 10 players; last marble is worth 1618 points: high score is 8317
// 13 players; last marble is worth 7999 points: high score is 146373
// 17 players; last marble is worth 1104 points: high score is 2764
// 21 players; last marble is worth 6111 points: high score is 54718
// 30 players; last marble is worth 5807 points: high score is 37305

func main() {
	var players, lastmarbelpoints int
	_, e := fmt.Sscanf(os.Args[1], "%d players; last marble is worth %d points", &players, &lastmarbelpoints)
	if e != nil {
		log.Fatal(e)
	}
	ppoints := make([]int, players, players)

	r := ring.New(1)
	r.Value = 0

	curplay := 0
	for i := 1; i < lastmarbelpoints*100; i++ {
		//fmt.Println("Marbel: ", i, "Player: ", curplay)
		if curplay == players {
			fmt.Println("Reset players")
			curplay = 0
		}
		if i%23 == 0 {
			ppoints[curplay] = ppoints[curplay] + i
			r = r.Move(-8)
			re := r.Unlink(1)
			r = r.Move(1)
			ppoints[curplay] = ppoints[curplay] + re.Value.(int)
			fmt.Println("M Points: ", re.Value.(int))
			fmt.Printf("Mod 23: %d Marbel remove: %d\n", i, re.Value.(int))

		} else {
			nr := ring.New(1)
			nr.Value = i
			// fmt.Printf("nr: %+v\n", nr)
			// fmt.Printf("r: %+v\n", r)
			r = r.Move(1)
			r.Link(nr)
			r = r.Move(1)
		}
		curplay++
	}
	var max int
	for _, r := range ppoints {
		if r > max {
			max = r
		}
	}
	fmt.Println(max)
}
