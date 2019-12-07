package main

import (
	"fmt"
	"strconv"
)

func intcomp(memory []int, inputs []int, outputs []int) []int {
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
		//fmt.Println("Instruction:", instopstr, "Pointer:", instpointer, "Mod3:", modp3, "Mod2:", modp2, "Mod1:", modp1, "Opcode:", opcode)
		if opcode == 99 {
			//fmt.Println("HALT")
			return outputs
			break
		}

		// input
		if opcode == 3 {
			//fmt.Println("Input", inputs[inputcount])
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
			outputs = append(outputs, tooutput)
			//fmt.Println("outputs innen:", outputs)
			//if tooutput != 0 {
			//fmt.Println("Fehler")
			//break
			//}
			instpointer = instpointer + 2
			continue
		}
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
		// add
		if opcode == 1 {
			memory[memory[instpointer+3]] = para1 + para2
			instpointer = instpointer + 4
			continue
		}
		// multi
		if opcode == 2 {
			memory[memory[instpointer+3]] = para1 * para2
			instpointer = instpointer + 4
			continue
		}

		// jump-if-true || jump-if-false
		if opcode == 5 {
			if para1 != 0 {
				instpointer = para2
				continue
			}
			instpointer = instpointer + 3
			continue

		}
		if opcode == 6 {
			if para1 == 0 {
				instpointer = para2
				continue
			}
			instpointer = instpointer + 3
			continue

		}

		// less-than || equals
		if opcode == 7 {
			if para1 < para2 {
				memory[memory[instpointer+3]] = 1
			} else {
				memory[memory[instpointer+3]] = 0
			}
			instpointer = instpointer + 4
			continue
		}

		if opcode == 8 {
			if para1 == para2 {
				memory[memory[instpointer+3]] = 1
			} else {
				memory[memory[instpointer+3]] = 0
			}

			instpointer = instpointer + 4
			continue
		}
	}
	return []int{}

}
