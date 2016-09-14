package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func run(code [maxCodeSize]float64) {
	count := 0
	reader := bufio.NewReader(os.Stdin)

	flt := false
	feq := false
	fgt := false

	var reg [regCount]float64
	var stack [stackSize]float64
	sp := 0

	for {
		if count >= maxCodeSize {
			break
		}

		// Debug
		// fmt.Println(b2s[code[count]])

		switch code[count] {
		case dsp:
			fmt.Println(data[int(code[count+1])])
			count += 2

		case jmp:
			count = int(code[count+1])

		case cmp:
			if reg[int(code[count+1])] < reg[int(code[count+2])] {
				flt = true
				feq = false
				fgt = false
			} else if reg[int(code[count+1])] == reg[int(code[count+2])] {
				flt = false
				feq = true
				fgt = false
			} else {
				flt = false
				feq = false
				fgt = true
			}
			count += 3

		case cmz:
			if reg[int(code[count+1])] < 0 {
				flt = true
				feq = false
				fgt = false
			} else if reg[int(code[count+1])] == 0 {
				flt = false
				feq = true
				fgt = false
			} else {
				flt = false
				feq = false
				fgt = true
			}
			count += 2

		case jlt:
			if flt {
				count = int(code[count+1])
			} else {
				count += 2
			}
		case jeq:
			if feq {
				count = int(code[count+1])
			} else {
				count += 2
			}
		case jgt:
			if fgt {
				count = int(code[count+1])
			} else {
				count += 2
			}

		case val:
			reg[int(code[count+2])] = code[count+1]
			count += 3
		case cpy:
			reg[int(code[count+2])] = reg[int(code[count+1])]
			count += 3
		case psh:
			stack[sp] = reg[int(code[count+1])]
			sp += 1
			count += 2
		case pop:
			sp -= 1
			reg[int(code[count+1])] = stack[sp]
			count += 2
		case add:
			reg[int(code[count+2])] += reg[int(code[count+1])]
			count += 3
		case sub:
			reg[int(code[count+2])] -= reg[int(code[count+1])]
			count += 3
		case mul:
			reg[int(code[count+2])] *= reg[int(code[count+1])]
			count += 3
		case div:
			reg[int(code[count+2])] /= reg[int(code[count+1])]
			count += 3
		case shw:
			fmt.Println(reg[int(code[count+1])])
			count += 2
		case get:
			txt, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("govm: could not read the input string")
				break
			}
			val, err := strconv.ParseFloat(strings.TrimSpace(txt), 64)
			if err != nil {
				fmt.Println("govm: invalid number")
				break
			}
			reg[int(code[count+1])] = val
			count += 2
		case nop:
			count += 1
		case hlt:
			count = maxCodeSize
		case flr:
			reg[int(code[count+1])] = math.Floor(reg[int(code[count+1])])
			count += 2
		case cel:
			reg[int(code[count+1])] = math.Ceil(reg[int(code[count+1])])
			count += 2
		case inc:
			reg[int(code[count+1])] += 1
			count += 2
		case dec:
			reg[int(code[count+1])] -= 1
			count += 2
		default:
			panic(fmt.Sprintf("Unexpected instruction code. %.0f\n", code[count]))
		}
	}
}
