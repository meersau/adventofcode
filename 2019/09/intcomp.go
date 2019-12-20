package main

import (
	"fmt"
	"strconv"
)

func intcomp(initmemory []int, inputs chan int, outputs chan int, amp string) {
	instpointer := 0
	relativebase := 0
	//inputcount := 0
	memory := make([]int, len(initmemory), 10000)
	copy(memory, initmemory)
	memory = memory[:cap(memory)]
	fmt.Println(len(memory))
	fmt.Println(cap(memory))
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
		fmt.Println("Amp:", amp, "Instruction:", instopstr, "Pointer:", instpointer, "Mod3:", modp3, "Mod2:", modp2, "Mod1:", modp1, "Opcode:", opcode)
		if opcode == 99 {
			fmt.Println(amp, "HALT")
			// return outputs
			// channel close?! ne, aber .... noch zu klären
			close(outputs)
			//close(inputs)
			break
		}

		var para1, para2 int
		if modp1 == 1 {
			//fmt.Println("Mod1:", modp1)
			para1 = memory[instpointer+1]
		} else if modp1 == 2 { // relative mode
			para1 = memory[relativebase+memory[instpointer+1]]
		} else { // Position Mode
			para1 = memory[memory[instpointer+1]]
		}

		// input 1 Para
		if opcode == 3 {
			if modp1 == 1 {
				//fmt.Println("Mod1:", modp1)
				memory[instpointer+1] = <-inputs
			} else if modp1 == 2 { // relative mode
				memory[relativebase+memory[instpointer+1]] = <-inputs
			} else { // Position Mode
				memory[memory[instpointer+1]] = <-inputs
			}
			//fmt.Println("Input", inputs[inputcount])
			// memory[memory[instpointer+1]] = <-inputs
			//fmt.Println(amp, "habe input")
			instpointer = instpointer + 2
			continue
		}

		// output one parameter
		if opcode == 4 {
			tooutput := para1
			fmt.Println(amp, " OUT: ", tooutput)
			outputs <- tooutput
			instpointer = instpointer + 2
			continue
		}
		// adjusts the relative base onyl one para
		if opcode == 9 {
			fmt.Println("Rel before", relativebase)
			relativebase = relativebase + para1
			fmt.Println("Rel after", relativebase)
			instpointer = instpointer + 2
			continue
		}

		if modp2 == 1 {
			//fmt.Println("Mod2:", modp2)
			para2 = memory[instpointer+2]
		} else if modp2 == 2 { // relative mode
			para2 = memory[relativebase+memory[instpointer+2]]
		} else { // Position Mode
			para2 = memory[memory[instpointer+2]]
		}

		var maddrpara3 int
		if modp3 == 1 {
			maddrpara3 = instpointer + 3
		} else if modp3 == 2 { // relative mode
			maddrpara3 = relativebase + memory[instpointer+3]
		} else { // Position Mode
			maddrpara3 = memory[instpointer+3]
		}

		// add
		if opcode == 1 {
			memory[maddrpara3] = para1 + para2
			instpointer = instpointer + 4
			continue
		}
		// multi
		if opcode == 2 {
			memory[maddrpara3] = para1 * para2
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

		// less-than
		if opcode == 7 {
			if para1 < para2 {
				memory[maddrpara3] = 1
			} else {
				memory[maddrpara3] = 0
			}
			instpointer = instpointer + 4
			continue
		}
		// equals
		if opcode == 8 {

			if para1 == para2 {
				memory[maddrpara3] = 1
			} else {
				memory[maddrpara3] = 0
			}

			instpointer = instpointer + 4
			continue
		}

	}
}
