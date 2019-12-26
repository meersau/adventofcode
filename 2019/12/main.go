package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type moon struct {
	x, y, z    int // pos
	vx, vy, vz int // velocity
}

func (m moon) String() string {
	return fmt.Sprintf("pos=<x=%d, y=%d, z=%d>, vel=<x=%d, y=%d, z=%d>", m.x, m.y, m.z, m.vx, m.vy, m.vz)
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	moons := make([]moon, 0)
	s := bufio.NewScanner(f)
	for s.Scan() {
		m := moon{}
		fmt.Sscanf(s.Text(), "<x=%d, y=%d, z=%d>", &m.x, &m.y, &m.z)
		m.vx = 0
		m.vy = 0
		m.vz = 0
		moons = append(moons, m)
	}

	for i := 0; i < 1000; i++ {
		stepupdatevelo(moons)

		stepaddvelo(moons)
	}
	fmt.Println(moons)
	fmt.Println(engery(moons))
	// 111402 too high
	// 40453963 too high
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func engery(m []moon) int {
	var e int
	for _, mo := range m {
		e = e + ((abs(mo.x) + abs(mo.y) + abs(mo.z)) * (abs(mo.vx) + abs(mo.vy) + abs(mo.vz)))
	}
	return e
}

func stepaddvelo(moons []moon) {
	for i, m := range moons {
		moons[i].x = moons[i].x + m.vx
		moons[i].y = moons[i].y + m.vy
		moons[i].z = moons[i].z + m.vz
	}
}
func stepupdatevelo(m []moon) {
	//update the velocity of every moon by applying gravity
	//
	//fmt.Println("Update VELO")
	for i := 0; i < len(m); i++ {
		for j := i + 1; j < len(m); j++ {
			//fmt.Println(i, j)
			// i Ganymede has an x position of 3
			// j Callisto has a x position of 5
			// i Ganymede's x velocity changes by +1 (because j 5 > i 3)
			// j Callisto's x velocity changes by -1 (because i 3 < j 5).
			if m[i].x < m[j].x {
				m[i].vx++
				m[j].vx--
			} else if m[i].x > m[j].x {
				m[i].vx--
				m[j].vx++
			}

			if m[i].y < m[j].y {
				m[i].vy++
				m[j].vy--
			} else if m[i].y > m[j].y {
				m[i].vy--
				m[j].vy++
			}

			if m[i].z < m[j].z {
				m[i].vz++
				m[j].vz--
			} else if m[i].z > m[j].z {
				m[i].vz--
				m[j].vz++
			}

		}
	}
}
