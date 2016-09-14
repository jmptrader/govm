package main

import (
	"fmt"
)

var b2s = map[float64]string{
	hlt: "hlt",
	val: "val",
	cpy: "cpy",
	add: "add",
	sub: "sub",
	mul: "mul",
	div: "div",
	shw: "shw",
	dsp: "dsp",
	get: "get",
	nop: "nop",
	flr: "flr",
	cel: "cel",
	inc: "inc",
	dec: "dec",
	jmp: "jmp",
	cmp: "cmp",
	cmz: "cmz",
	jlt: "jlt",
	jeq: "jeq",
	jgt: "jgt",
}

func decode(code [maxCodeSize]float64) (ds [maxCodeSize]string) {
	count := 0

	for {
		if count >= maxCodeSize {
			break
		}

		switch code[count] {
		case cmp:
			ds[count] = "cmp"
			ds[count+1] = fmt.Sprintf("%%%f", code[count+1])
			ds[count+2] = fmt.Sprintf("%%%f\n", code[count+2])
			count += 3
		case cmz:
			ds[count] = "cmz"
			ds[count+1] = fmt.Sprintf("%%%f", code[count+1])
			count += 2

		case jmp:
			ds[count] = "jmp"
			ds[count+1] = fmt.Sprintf("%.0f\n", code[count+1])
			count += 2

		case jlt:
			ds[count] = "jlt"
			ds[count+1] = fmt.Sprintf("%.0f\n", code[count+1])
			count += 2
		case jeq:
			ds[count] = "jeq"
			ds[count+1] = fmt.Sprintf("%.0f\n", code[count+1])
			count += 2
		case jgt:
			ds[count] = "jgt"
			ds[count+1] = fmt.Sprintf("%.0f\n", code[count+1])
			count += 2

		case val:
			ds[count] = "val"
			ds[count+1] = fmt.Sprintf("%f", code[count+1])
			ds[count+2] = fmt.Sprintf("%%%f\n", code[count+2])
			count += 3
		case cpy:
			ds[count] = "cpy"
			ds[count+1] = fmt.Sprintf("%%%f", code[count+1])
			ds[count+2] = fmt.Sprintf("%%%f\n", code[count+2])
			count += 3
		case psh:
			ds[count] = "psh"
			ds[count+1] = fmt.Sprintf("%f", code[count+1])
			count += 2
		case pop:
			ds[count] = "pop"
			ds[count+1] = fmt.Sprintf("%%%f", code[count+1])
			count += 2
		case add:
			ds[count] = "add"
			ds[count+1] = fmt.Sprintf("%%%f", code[count+1])
			ds[count+2] = fmt.Sprintf("%%%f\n", code[count+2])
			count += 3
		case sub:
			ds[count] = "sub"
			ds[count+1] = fmt.Sprintf("%%%f", code[count+1])
			ds[count+2] = fmt.Sprintf("%%%f\n", code[count+2])
			count += 3
		case mul:
			ds[count] = "mul"
			ds[count+1] = fmt.Sprintf("%%%f", code[count+1])
			ds[count+2] = fmt.Sprintf("%%%f\n", code[count+2])
			count += 3
		case div:
			ds[count] = "div"
			ds[count+1] = fmt.Sprintf("%%%f", code[count+1])
			ds[count+2] = fmt.Sprintf("%%%f\n", code[count+2])
			count += 3
		case shw:
			ds[count] = "shw"
			ds[count+1] = fmt.Sprintf("%%%f\n", code[count+1])
			count += 2
		case dsp:
			ds[count] = "dsp"
			ds[count+1] = fmt.Sprintf("%.0f\n", code[count+1])
			count += 2
		case get:
			ds[count] = "get"
			ds[count+1] = fmt.Sprintf("%%%f\n", code[count+1])
			count += 2
		case nop:
			ds[count] = "nop\n"
			count += 1
		case hlt:
			ds[count] = "hlt\n"
			count += 1
		case flr:
			ds[count] = "flr"
			ds[count+1] = fmt.Sprintf("%%%f\n", code[count+1])
			count += 2
		case cel:
			ds[count] = "cel"
			ds[count+1] = fmt.Sprintf("%%%f\n", code[count+1])
			count += 2
		case inc:
			ds[count] = "inc"
			ds[count+1] = fmt.Sprintf("%%%f\n", code[count+1])
			count += 2
		case dec:
			ds[count] = "dec"
			ds[count+1] = fmt.Sprintf("%%%f\n", code[count+1])
			count += 2
		default:
			panic(fmt.Sprintf("Decoding error %f", code[count]))
		}
	}

	return
}
