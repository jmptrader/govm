package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func runtimeError(msg string, code, index float64) {
	exit(fmt.Sprintf("runtime: %s [%0.f @ %0.f]", msg, code, index))
}

func run(code []float64) {
	count := 0
	reader := bufio.NewReader(os.Stdin)

	flt := false
	feq := false
	fgt := false

	var reg [regCount]float64

	var callStack [callStackSize]int
	csp := 0

	var stack [stackSize]float64
	sp := 0

	halted := false

	for {
		if count >= len(code) || halted {
			break
		}

		// Debug
		// fmt.Println(b2s[code[count]])

		switch code[count] {
		case hlt:
			halted = true
		case nop:
			count += 1

		case cll:
			callStack[csp] = count + 2
			csp += 1
			count = int(code[count+1])
		case ret:
			csp -= 1
			count = callStack[csp]

		case val:
			reg[int(code[count+2])] = code[count+1]
			count += 3
		case mov:
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

		case rac:
			reg[int(code[count+1])] = float64(len(os.Args)) - 2
			count += 2
		case rad:
			val, err := strconv.ParseFloat(os.Args[int(reg[int(code[count+1])])+2], 64)
			if err != nil {
				runtimeError("invalid number", rad, float64(count))
				break
			}
			reg[int(code[count+2])] = val
			count += 3

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
		case inc:
			reg[int(code[count+1])] += 1
			count += 2
		case dec:
			reg[int(code[count+1])] -= 1
			count += 2

		case flr:
			reg[int(code[count+1])] = math.Floor(reg[int(code[count+1])])
			count += 2
		case cel:
			reg[int(code[count+1])] = math.Ceil(reg[int(code[count+1])])
			count += 2
		case abs:
			reg[int(code[count+1])] = math.Abs(reg[int(code[count+1])])
			count += 2
		case sqr:
			reg[int(code[count+1])] = math.Sqrt(reg[int(code[count+1])])
			count += 2
		case sin:
			reg[int(code[count+1])] = math.Sin(reg[int(code[count+1])])
			count += 2
		case cos:
			reg[int(code[count+1])] = math.Cos(reg[int(code[count+1])])
			count += 2
		case tan:
			reg[int(code[count+1])] = math.Tan(reg[int(code[count+1])])
			count += 2
		case exp:
			reg[int(code[count+1])] = math.Exp(reg[int(code[count+1])])
			count += 2
		case log:
			reg[int(code[count+1])] = math.Log(reg[int(code[count+1])])
			count += 2
		case gam:
			reg[int(code[count+1])] = math.Gamma(reg[int(code[count+1])])
			count += 2
		case pow:
			reg[int(code[count+2])] = math.Pow(reg[int(code[count+2])], reg[int(code[count+1])])
			count += 2
		case pwa:
			reg[int(code[count+2])] = math.Pow(reg[int(code[count+1])], reg[int(code[count+2])])
			count += 2

		case jmp:
			count = int(code[count+1])
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

		case shw:
			fmt.Println(reg[int(code[count+1])])
			count += 2
		case get:
			txt, err := reader.ReadString('\n')
			if err != nil {
				runtimeError("could not read the input string", get, float64(count))
				break
			}
			val, err := strconv.ParseFloat(strings.TrimSpace(txt), 64)
			if err != nil {
				runtimeError("invalid number", get, float64(count))
				break
			}
			reg[int(code[count+1])] = val
			count += 2
		case dsp:
			fmt.Println(data[int(code[count+1])])
			count += 2

		default:
			runtimeError("invalid instruction", code[count], float64(count))
		}
	}
}
