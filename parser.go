package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func compilerError(msg, file, str string, line int) {
	exit(fmt.Sprintf("compiler: %s [%s] at line %d in %s", msg, str, line, file))
}

// We use this so that if a compilation error occurr we know the faulty line
type labelInfo struct {
	name string
	file string
	line int
}

var labels map[string]int
var labelsPending map[int]labelInfo

func init() {
	labels = make(map[string]int)
	labelsPending = make(map[int]labelInfo)
}

var data [dataSize]string

func parser(f *os.File, fileName string, start int) (code []float64, count int) {
	code = make([]float64, 0)

	reader := bufio.NewReader(f)

	lineNumber := 0
	count = start

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

		ins := strings.Fields(line)

		if len(ins) == 0 {
			continue
		}

		switch ins[0] {
		// Pseudo-instructions

		case "str":
			data[dataCurr] = strings.Join(ins[2:len(ins)], " ")
			dataMap[ins[1]] = dataCurr
			dataCurr += 1

		case "lab":
			if labels[ins[1]] != 0 {
				compilerError("label redefinition", fileName, ins[1], lineNumber)
			}
			labels[ins[1]] = count

		// Instructions

		case "hlt":
			code = append(code, hlt)
			count += 1
		case "nop":
			break

		case "cll":
			code = append(code, cll)
			code = append(code, 0.0) // Placeholder
			labelsPending[count+1] = labelInfo{ins[1], fileName, lineNumber}
			count += 2
		case "ret":
			code = append(code, ret)
			count += 1

		case "val":
			code = append(code, val)
			code = append(code, getLiteral(ins[1]))
			code = append(code, getRegister(ins[2]))
			count += 3
		case "mov":
			code = append(code, mov)
			code = append(code, getRegister(ins[1]))
			code = append(code, getRegister(ins[2]))
			count += 3
		case "psh":
			code = append(code, psh)
			code = append(code, getRegister(ins[1]))
			count += 2
		case "pop":
			code = append(code, pop)
			code = append(code, getRegister(ins[1]))
			count += 2

		case "add":
			code = append(code, add)
			code = append(code, getRegister(ins[1]))
			code = append(code, getRegister(ins[2]))
			count += 3
		case "sub":
			code = append(code, sub)
			code = append(code, getRegister(ins[1]))
			code = append(code, getRegister(ins[2]))
			count += 3
		case "mul":
			code = append(code, mul)
			code = append(code, getRegister(ins[1]))
			code = append(code, getRegister(ins[2]))
			count += 3
		case "div":
			code = append(code, div)
			code = append(code, getRegister(ins[1]))
			code = append(code, getRegister(ins[2]))
			count += 3
		case "inc":
			code = append(code, inc)
			code = append(code, getRegister(ins[1]))
			count += 2
		case "dec":
			code = append(code, dec)
			code = append(code, getRegister(ins[1]))
			count += 2

		case "flr":
			code = append(code, flr)
			code = append(code, getRegister(ins[1]))
			count += 2
		case "cel":
			code = append(code, cel)
			code = append(code, getRegister(ins[1]))
			count += 2

		case "jmp":
			code = append(code, jmp)
			code = append(code, 0.0) // Placeholder
			labelsPending[count+1] = labelInfo{ins[1], fileName, lineNumber}
			count += 2
		case "jlt":
			code = append(code, jlt)
			code = append(code, 0.0) // Placeholder
			labelsPending[count+1] = labelInfo{ins[1], fileName, lineNumber}
			count += 2
		case "jeq":
			code = append(code, jeq)
			code = append(code, 0.0) // Placeholder
			labelsPending[count+1] = labelInfo{ins[1], fileName, lineNumber}
			count += 2
		case "jgt":
			code = append(code, jgt)
			code = append(code, 0.0) // Placeholder
			labelsPending[count+1] = labelInfo{ins[1], fileName, lineNumber}
			count += 2

		case "cmp":
			code = append(code, cmp)
			code = append(code, getRegister(ins[1]))
			code = append(code, getRegister(ins[2]))
			count += 3
		case "cmz":
			code = append(code, cmz)
			code = append(code, getRegister(ins[1]))
			count += 2

		case "shw":
			code = append(code, shw)
			code = append(code, getRegister(ins[1]))
			count += 2
		case "get":
			code = append(code, get)
			code = append(code, getRegister(ins[1]))
			count += 2
		case "dsp":
			code = append(code, dsp)
			code = append(code, float64(dataMap[ins[1]]))
			count += 2

		default:
			compilerError("invalid instruction", fileName, ins[0], lineNumber)
		}
	}

	return
}

func evaluateLabels(code []float64) {
	for index, info := range labelsPending {
		code[index] = float64(labels[info.name])
		if code[index] == 0 {
			compilerError("undefined label", info.file, info.name, info.line)
		}
	}
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
