package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var data [dataSize]string

func parser(f *os.File) (code [maxCodeSize]float64) {
	reader := bufio.NewReader(f)

	labels := make(map[string]int)
	labelsPending := make(map[int]string)

	lineNumber := 0
	count := 0

	dataMap := make(map[string]int)
	dataCurr := 0

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		lineNumber += 1

		instructions := strings.Fields(line)

		if len(instructions) == 0 {
			continue
		}

		switch instructions[0] {
		case "lab":
			labels[instructions[1]] = count
		case "def":
			labels[instructions[1]] = count

		case "cll":
			code[count] = cll
			labelsPending[count+1] = instructions[1]
			count += 2
		case "ret":
			code[count] = ret
			count += 1

		case "str":
			data[dataCurr] = strings.Join(instructions[2:len(instructions)], " ")
			dataMap[instructions[1]] = dataCurr
			dataCurr += 1
		case "dsp":
			code[count] = dsp
			code[count+1] = float64(dataMap[instructions[1]])
			count += 2

		case "jmp":
			code[count] = jmp
			labelsPending[count+1] = instructions[1]
			count += 2
		case "jlt":
			code[count] = jlt
			labelsPending[count+1] = instructions[1]
			count += 2
		case "jeq":
			code[count] = jeq
			labelsPending[count+1] = instructions[1]
			count += 2
		case "jgt":
			code[count] = jgt
			labelsPending[count+1] = instructions[1]
			count += 2

		case "cmp":
			code[count] = cmp
			code[count+1] = getRegister(instructions[1])
			code[count+2] = getRegister(instructions[2])
			count += 3
		case "cmz":
			code[count] = cmz
			code[count+1] = getRegister(instructions[1])
			count += 2

		case "val":
			code[count] = val
			code[count+1] = getLiteral(instructions[1])
			code[count+2] = getRegister(instructions[2])
			count += 3
		case "cpy":
			code[count] = cpy
			code[count+1] = getRegister(instructions[1])
			code[count+2] = getRegister(instructions[2])
			count += 3
		case "psh":
			code[count] = psh
			code[count+1] = getRegister(instructions[1])
			count += 2
		case "pop":
			code[count] = pop
			code[count+1] = getRegister(instructions[1])
			count += 2

		case "add":
			code[count] = add
			code[count+1] = getRegister(instructions[1])
			code[count+2] = getRegister(instructions[2])
			count += 3
		case "sub":
			code[count] = sub
			code[count+1] = getRegister(instructions[1])
			code[count+2] = getRegister(instructions[2])
			count += 3
		case "mul":
			code[count] = mul
			code[count+1] = getRegister(instructions[1])
			code[count+2] = getRegister(instructions[2])
			count += 3
		case "div":
			code[count] = div
			code[count+1] = getRegister(instructions[1])
			code[count+2] = getRegister(instructions[2])
			count += 3
		case "shw":
			code[count] = shw
			code[count+1] = getRegister(instructions[1])
			count += 2
		case "get":
			code[count] = get
			code[count+1] = getRegister(instructions[1])
			count += 2
		case "nop":
			code[count] = nop
			count += 1
		case "hlt":
			code[count] = hlt
			count += 1
		case "flr":
			code[count] = flr
			code[count+1] = getRegister(instructions[1])
			count += 2
		case "cel":
			code[count] = cel
			code[count+1] = getRegister(instructions[1])
			count += 2
		case "inc":
			code[count] = inc
			code[count+1] = getRegister(instructions[1])
			count += 2
		case "dec":
			code[count] = dec
			code[count+1] = getRegister(instructions[1])
			count += 2
		default:
			exit(fmt.Sprintf("invalid '%s' at line %d", instructions[0], lineNumber))
		}
	}

	for index, label := range labelsPending {
		code[index] = float64(labels[label])
	}

	return
}

func getLiteral(ins string) float64 {
	lit, err := strconv.ParseFloat(ins, 64)
	if err != nil {
		panic(err)
	}
	return lit
}

func getRegister(ins string) float64 {
	tins := ins[1:len(ins)]
	reg, err := strconv.ParseFloat(tins, 64)
	if err != nil {
		exit(fmt.Sprintf("%s is an invalid register", ins))
	}
	return float64(reg)
}
