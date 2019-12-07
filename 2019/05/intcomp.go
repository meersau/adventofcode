package main

import (
	"fmt"
	"strconv"
)

func intcomp(memory []int, inputs []int) {
	instpointer := 0
	inputcount := 0

	// fmt.Println(instopstr, modp3, modp2, modp1, opcode)
	for {
		var modp3, modp2, modp1, opcode int

		instopstr := strconv.Itoa(memory[instpointer])
		if len(instopstr) == 1 {
			opcode = memory[instpointer]
		} else {
			if len(instopstr) == 5 {
				fmt.Sscanf(instopstr, "%1d%1d%1d%2d", &modp3, &modp2, &modp1, &opcode)
			} else if len(instopstr) == 4 {
				fmt.Sscanf(instopstr, "%1d%1d%2d", &modp2, &modp1, &opcode)
			} else if len(instopstr) == 3 {
				fmt.Sscanf(instopstr, "%1d%2d", &modp1, &opcode)
			} else {
				fmt.Sscanf(instopstr, "%2d", &opcode)
			}
		}
		fmt.Println("Instruction:", instopstr, "Pointer:", instpointer, "Mod3:", modp3, "Mod2:", modp2, "Mod1:", modp1, "Opcode:", opcode)
		if opcode == 99 {
			fmt.Println("HALT")
			break
		}
		// add
		if opcode == 1 {
			var para1, para2 int
			if modp1 == 1 {
				//fmt.Println("Mod1:", modp1)
				para1 = memory[instpointer+1]
			} else {
				para1 = memory[memory[instpointer+1]]
			}
			if modp2 == 1 {
				//fmt.Println("Mod2:", modp2)
				para2 = memory[instpointer+2]
			} else {
				para2 = memory[memory[instpointer+2]]
			}
			//fmt.Println("Para1:", para1, "Para2:", para2)
			//fmt.Println("Vor:", memory[memory[instpointer+3]])
			memory[memory[instpointer+3]] = para1 + para2
			//fmt.Println("Danach", memory[memory[instpointer+3]])
			instpointer = instpointer + 4
			continue
		}
		// multi
		if opcode == 2 {
			var para1, para2 int
			if modp1 == 1 {
				para1 = memory[instpointer+1]
			} else {
				para1 = memory[memory[instpointer+1]]
			}
			if modp2 == 1 {
				para2 = memory[instpointer+2]
			} else {
				para2 = memory[memory[instpointer+2]]
			}
			memory[memory[instpointer+3]] = para1 * para2
			instpointer = instpointer + 4
			continue
		}

		// input
		if opcode == 3 {
			fmt.Println("Input", inputs[inputcount])
			memory[memory[instpointer+1]] = inputs[inputcount]
			inputcount++
			instpointer = instpointer + 2
			continue
		}

		// output
		if opcode == 4 {
			var tooutput int
			if modp1 == 1 {
				tooutput = memory[instpointer+1]
			} else {
				tooutput = memory[memory[instpointer+1]]
			}

			fmt.Println("OUT: ", tooutput)
			//if tooutput != 0 {
			//fmt.Println("Fehler")
			//break
			//}
			instpointer = instpointer + 2
			continue
		}
	}
}
