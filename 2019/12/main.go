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
	sts := steptosame(moons)
	fmt.Println(sts)
	fmt.Println(lcm(lcm(sts[0], sts[1]), sts[2]))
	// to low 919590 p2
	// 111402 too high
	// 40453963 too high
}

/*
const gcd = (a, b) => a ? gcd(b % a, a) : b
const lcm = (a, b) => a * b / gcd(a, b)

*/

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func checkx(m, init []moon) bool {
	for i := 0; i < len(m); i++ {
		if m[i].x != init[i].x || m[i].vx != init[i].vx {
			return false
		}
	}
	return true
}
func checky(m, init []moon) bool {
	for i := 0; i < len(m); i++ {
		if m[i].y != init[i].y || m[i].vy != init[i].vy {
			return false
		}
	}
	return true
}

func checkz(m, init []moon) bool {
	for i := 0; i < len(m); i++ {
		if m[i].z != init[i].z || m[i].vz != init[i].vz {
			return false
		}
	}
	return true
}

func steptosame(moons []moon) []int {
	init := make([]moon, len(moons))
	copy(init, moons)
	steps := make([]int, 3)

	var stopx, stopy, stopz bool
	// var stepx, stepy, stepz int
	step := 0
	for !(stopx && stopy && stopz) {
		stepupdatevelo(moons)
		stepaddvelo(moons)
		step++

		if !stopx && checkx(moons, init) {
			steps[0] = step //+ 1
			stopx = true
		}
		if !stopy && checky(moons, init) {
			steps[1] = step //+ 1
			stopy = true
			fmt.Println("Only one")
		}
		if !stopz && checkz(moons, init) {
			steps[2] = step //+ 1
			stopz = true
		}
	}

	return steps
}

func part1(moons []moon) int {
	for i := 0; i < 1000; i++ {
		stepupdatevelo(moons)
		stepaddvelo(moons)
	}
	return engery(moons)
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
