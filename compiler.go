package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func compilerError(msg, file, str string, line int) {
	exit(fmt.Sprintf("compiler: %s [%s] at line %d in %s", msg, str, line, file))
}

func encoderError(msg string) {
	exit(fmt.Sprintf("encoder: %s", msg))
}

func decoderError(msg string) {
	exit(fmt.Sprintf("decoder: %s", msg))
}

// We use this so that if a compilation error occurr we know the faulty line
type labelInfo struct {
	name string
	file string
	line int
}

var labels map[string]int
var labelsPending map[int]labelInfo

var data []string
var dataCurr int
var dataMap map[string]int

func init() {
	labels = make(map[string]int)
	labelsPending = make(map[int]labelInfo)
	data = make([]string, 0)
	dataCurr = 0
	dataMap = make(map[string]int)
}

func parser(f *os.File, fileName string, start int) (code []float64, count int) {
	code = make([]float64, 0)

	reader := bufio.NewReader(f)

	lineNumber := 0
	count = start

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				compilerError("failed line read", fileName, "ReadString()", lineNumber)
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
			data = append(data, strings.Join(ins[2:len(ins)], " "))
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
			code = append(code, getLiteral(ins[1], fileName, lineNumber))
			code = append(code, getRegister(ins[2], fileName, lineNumber))
			count += 3
		case "mov":
			code = append(code, mov)
			code = append(code, getRegister(ins[1], fileName, lineNumber))
			code = append(code, getRegister(ins[2], fileName, lineNumber))
			count += 3
		case "swp":
			code = append(code, swp)
			code = append(code, getRegister(ins[1], fileName, lineNumber))
			code = append(code, getRegister(ins[2], fileName, lineNumber))
			count += 3
		case "psh":
			code = append(code, psh)
			code = append(code, getRegister(ins[1], fileName, lineNumber))
			count += 2
		case "pop":
			code = append(code, pop)
			code = append(code, getRegister(ins[1], fileName, lineNumber))
			count += 2

		case "rac":
			code = append(code, rac)
			code = append(code, getRegister(ins[1], fileName, lineNumber))
			count += 2
		case "rad":
			code = append(code, rad)
			code = append(code, getRegister(ins[1], fileName, lineNumber))
			code = append(code, getRegister(ins[2], fileName, lineNumber))
			count += 3

		case "add":
			code = append(code, add)
			code = append(code, getRegister(ins[1], fileName, lineNumber))
			code = append(code, getRegister(ins[2], fileName, lineNumber))
			count += 3
		case "sub":
			code = append(code, sub)
			code = append(code, getRegister(ins[1], fileName, lineNumber))
			code = append(code, getRegister(ins[2], fileName, lineNumber))
			count += 3
		case "mul":
			code = append(code, mul)
			code = append(code, getRegister(ins[1], fileName, lineNumber))
			code = append(code, getRegister(ins[2], fileName, lineNumber))
			count += 3
		case "div":
			code = append(code, div)
			code = append(code, getRegister(ins[1], fileName, lineNumber))
			code = append(code, getRegister(ins[2], fileName, lineNumber))
			count += 3
		case "inc":
			code = append(code, inc)
			code = append(code, getRegister(ins[1], fileName, lineNumber))
			count += 2
		case "dec":
			code = append(code, dec)
			code = append(code, getRegister(ins[1], fileName, lineNumber))
			count += 2

		case "flr":
			code = append(code, flr)
			code = append(code, getRegister(ins[1], fileName, lineNumber))
			count += 2
		case "cel":
			code = append(code, cel)
			code = append(code, getRegister(ins[1], fileName, lineNumber))
			count += 2
		case "abs":
			code = append(code, abs)
			code = append(code, getRegister(ins[1], fileName, lineNumber))
			count += 2
		case "sqr":
			code = append(code, sqr)
			code = append(code, getRegister(ins[1], fileName, lineNumber))
			count += 2
		case "sin":
			code = append(code, sin)
			code = append(code, getRegister(ins[1], fileName, lineNumber))
			count += 2
		case "cos":
			code = append(code, cos)
			code = append(code, getRegister(ins[1], fileName, lineNumber))
			count += 2
		case "tan":
			code = append(code, tan)
			code = append(code, getRegister(ins[1], fileName, lineNumber))
			count += 2
		case "exp":
			code = append(code, exp)
			code = append(code, getRegister(ins[1], fileName, lineNumber))
			count += 2
		case "log":
			code = append(code, log)
			code = append(code, getRegister(ins[1], fileName, lineNumber))
			count += 2
		case "gam":
			code = append(code, gam)
			code = append(code, getRegister(ins[1], fileName, lineNumber))
			count += 2
		case "pow":
			code = append(code, pow)
			code = append(code, getRegister(ins[1], fileName, lineNumber))
			code = append(code, getRegister(ins[2], fileName, lineNumber))
			count += 3
		case "pwa":
			code = append(code, pwa)
			code = append(code, getRegister(ins[1], fileName, lineNumber))
			code = append(code, getRegister(ins[2], fileName, lineNumber))
			count += 3

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
		case "jge":
			code = append(code, jge)
			code = append(code, 0.0) // Placeholder
			labelsPending[count+1] = labelInfo{ins[1], fileName, lineNumber}
			count += 2
		case "jle":
			code = append(code, jle)
			code = append(code, 0.0) // Placeholder
			labelsPending[count+1] = labelInfo{ins[1], fileName, lineNumber}
			count += 2
		case "jne":
			code = append(code, jne)
			code = append(code, 0.0) // Placeholder
			labelsPending[count+1] = labelInfo{ins[1], fileName, lineNumber}
			count += 2

		case "cmp":
			code = append(code, cmp)
			code = append(code, getRegister(ins[1], fileName, lineNumber))
			code = append(code, getRegister(ins[2], fileName, lineNumber))
			count += 3
		case "cmz":
			code = append(code, cmz)
			code = append(code, getRegister(ins[1], fileName, lineNumber))
			count += 2

		case "shw":
			code = append(code, shw)
			code = append(code, getRegister(ins[1], fileName, lineNumber))
			count += 2
		case "get":
			code = append(code, get)
			code = append(code, getRegister(ins[1], fileName, lineNumber))
			count += 2
		case "dsp":
			code = append(code, dsp)
			code = append(code, float64(dataMap[ins[1]]))
			count += 2

		case "dty":
			code = append(code, dty)
			code = append(code, getRegister(ins[1], fileName, lineNumber))
			count += 2
		case "dtm":
			code = append(code, dtm)
			code = append(code, getRegister(ins[1], fileName, lineNumber))
			count += 2
		case "dtd":
			code = append(code, dtd)
			code = append(code, getRegister(ins[1], fileName, lineNumber))
			count += 2
		case "tmh":
			code = append(code, tmh)
			code = append(code, getRegister(ins[1], fileName, lineNumber))
			count += 2
		case "tmm":
			code = append(code, tmm)
			code = append(code, getRegister(ins[1], fileName, lineNumber))
			count += 2
		case "tms":
			code = append(code, tms)
			code = append(code, getRegister(ins[1], fileName, lineNumber))
			count += 2
		case "now":
			code = append(code, now)
			code = append(code, getRegister(ins[1], fileName, lineNumber))
			count += 2

		case "sed":
			code = append(code, sed)
			code = append(code, getRegister(ins[1], fileName, lineNumber))
			count += 2
		case "rnd":
			code = append(code, rnd)
			code = append(code, getRegister(ins[1], fileName, lineNumber))
			count += 2
		case "rnn":
			code = append(code, rnn)
			code = append(code, getRegister(ins[1], fileName, lineNumber))
			count += 2

		default:
			if ins[0][len(ins[0])-1] == ':' {
				labName := ins[0][:len(ins[0])-1]
				if labels[labName] != 0 {
					compilerError("label redefinition", fileName, labName, lineNumber)
				}
				labels[labName] = count
				break
			}

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

func getLiteral(ins, file string, line int) float64 {
	lit, err := strconv.ParseFloat(ins, 64)
	if err != nil {
		compilerError("invalid literal", file, ins, line)
	}
	return lit
}

func getRegister(ins, file string, line int) float64 {
	tins := ins[1:len(ins)]
	reg, err := strconv.ParseFloat(tins, 64)
	if err != nil {
		compilerError("invalid register", file, ins, line)
	}
	return float64(reg)
}

func writeFile(code []float64, fileName string) {
	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		encoderError("could not create binary file")
	}

	// Header
	var headerValue int64
	headerValue = 10002
	err = binary.Write(file, binary.LittleEndian, headerValue)
	if err != nil {
		encoderError("could not write header")
	}

	// Code size
	var codeSize int64
	codeSize = int64(len(code))
	err = binary.Write(file, binary.LittleEndian, codeSize)
	if err != nil {
		encoderError("could not write code size")
	}

	// Code
	for i := 0; i < len(code); i++ {
		err = binary.Write(file, binary.LittleEndian, code[i])
		if err != nil {
			encoderError("could not write code entry")
		}
	}

	// Data size
	var dataStackSize int64
	dataStackSize = int64(dataCurr)
	err = binary.Write(file, binary.LittleEndian, dataStackSize)
	if err != nil {
		encoderError("could not write data size")
	}

	// Data
	for i := 0; i < dataCurr; i++ {
		_, err = fmt.Fprintln(file, data[i])
		if err != nil {
			encoderError("could not write data entry")
		}
	}
}

func loadFile(fileName string) (code []float64) {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		decoderError("could not open file")
	}

	var headerValue int64
	err = binary.Read(file, binary.LittleEndian, &headerValue)
	if err != nil {
		decoderError("could not read header")
	}
	if headerValue != 10002 {
		decoderError("invalid header value")
	}

	var codeSize int64
	err = binary.Read(file, binary.LittleEndian, &codeSize)
	if err != nil {
		decoderError("could not read code size")
	}
	code = make([]float64, codeSize)
	for i := 0; i < int(codeSize); i++ {
		err = binary.Read(file, binary.LittleEndian, &code[i])
		if err != nil {
			decoderError("could not read code entry")
		}
	}

	var dataStackSize int64
	err = binary.Read(file, binary.LittleEndian, &dataStackSize)
	if err != nil {
		decoderError("could not read data size")
	}

	r := bufio.NewReader(file)
	for i := 0; i < int(dataStackSize); i++ {
		str, err := r.ReadString('\n')
		if err != nil {
			decoderError("could not read data entry")
		}
		data = append(data, str[:len(str)-1])
	}

	return
}
