package main

import (
	"fmt"
)

var b2s = map[float64]string{
	hlt: "hlt",
	nop: "nop",

	cll: "cll",
	ret: "ret",

	val: "val",
	mov: "mov",
	psh: "psh",
	pop: "pop",

	add: "add",
	sub: "sub",
	mul: "mul",
	div: "div",
	inc: "inc",
	dec: "dec",

	flr: "flr",
	cel: "cel",
	abs: "abs",
	sqr: "sqr",
	sin: "sin",
	cos: "cos",
	tan: "tan",
	exp: "exp",
	log: "log",
	gam: "gam",
	pow: "pow",
	pwa: "pwa",

	jmp: "jmp",
	jlt: "jlt",
	jeq: "jeq",
	jgt: "jgt",

	cmp: "cmp",
	cmz: "cmz",

	shw: "shw",
	dsp: "dsp",
	get: "get",
}

func decode(code []float64) (ds []string) {
	ds = make([]string, len(code))
	count := 0

	for {
		if count >= len(code) {
			break
		}

		switch code[count] {
		case hlt:
			ds[count] = dbgLst("hlt")
			count += 1
		case nop:
			ds[count] = dbgLst("nop")
			count += 1

		case cll:
			ds[count] = "cll"
			ds[count+1] = dbgLst(dbgIdx(code[count+1]))
			count += 2
		case ret:
			ds[count] = dbgLst("ret")
			count += 1

		case val:
			ds[count] = "val"
			ds[count+1] = dbgLit(code[count+1])
			ds[count+2] = dbgLst(dbgReg(code[count+2]))
			count += 3
		case mov:
			ds[count] = "mov"
			ds[count+1] = dbgReg(code[count+1])
			ds[count+2] = dbgLst(dbgReg(code[count+2]))
			count += 3
		case psh:
			ds[count] = "psh"
			ds[count+1] = dbgLst(dbgReg(code[count+1]))
			count += 2
		case pop:
			ds[count] = "pop"
			ds[count+1] = dbgLst(dbgReg(code[count+1]))
			count += 2

		case add:
			ds[count] = "add"
			ds[count+1] = dbgReg(code[count+1])
			ds[count+2] = dbgLst(dbgReg(code[count+2]))
			count += 3
		case sub:
			ds[count] = "sub"
			ds[count+1] = dbgReg(code[count+1])
			ds[count+2] = dbgLst(dbgReg(code[count+2]))
			count += 3
		case mul:
			ds[count] = "mul"
			ds[count+1] = dbgReg(code[count+1])
			ds[count+2] = dbgLst(dbgReg(code[count+2]))
			count += 3
		case div:
			ds[count] = "div"
			ds[count+1] = dbgReg(code[count+1])
			ds[count+2] = dbgLst(dbgReg(code[count+2]))
			count += 3
		case inc:
			ds[count] = "inc"
			ds[count+1] = dbgLst(dbgReg(code[count+1]))
			count += 2
		case dec:
			ds[count] = "dec"
			ds[count+1] = dbgLst(dbgReg(code[count+1]))
			count += 2

		case flr:
			ds[count] = "flr"
			ds[count+1] = dbgLst(dbgReg(code[count+1]))
			count += 2
		case cel:
			ds[count] = "cel"
			ds[count+1] = dbgLst(dbgReg(code[count+1]))
			count += 2
		case abs:
			ds[count] = "abs"
			ds[count+1] = dbgLst(dbgReg(code[count+1]))
			count += 2
		case sqr:
			ds[count] = "sqr"
			ds[count+1] = dbgLst(dbgReg(code[count+1]))
			count += 2
		case sin:
			ds[count] = "sin"
			ds[count+1] = dbgLst(dbgReg(code[count+1]))
			count += 2
		case cos:
			ds[count] = "cos"
			ds[count+1] = dbgLst(dbgReg(code[count+1]))
			count += 2
		case tan:
			ds[count] = "tan"
			ds[count+1] = dbgLst(dbgReg(code[count+1]))
			count += 2
		case exp:
			ds[count] = "exp"
			ds[count+1] = dbgLst(dbgReg(code[count+1]))
			count += 2
		case log:
			ds[count] = "log"
			ds[count+1] = dbgLst(dbgReg(code[count+1]))
			count += 2
		case gam:
			ds[count] = "gam"
			ds[count+1] = dbgLst(dbgReg(code[count+1]))
			count += 2
		case pow:
			ds[count] = "pow"
			ds[count+1] = dbgReg(code[count+1])
			ds[count+2] = dbgLst(dbgReg(code[count+2]))
			count += 3
		case pwa:
			ds[count] = "pwa"
			ds[count+1] = dbgReg(code[count+1])
			ds[count+2] = dbgLst(dbgReg(code[count+2]))
			count += 3

		case jmp:
			ds[count] = "jmp"
			ds[count+1] = dbgLst(dbgIdx(code[count+1]))
			count += 2
		case jlt:
			ds[count] = "jlt"
			ds[count+1] = dbgLst(dbgIdx(code[count+1]))
			count += 2
		case jeq:
			ds[count] = "jeq"
			ds[count+1] = dbgLst(dbgIdx(code[count+1]))
			count += 2
		case jgt:
			ds[count] = "jgt"
			ds[count+1] = dbgLst(dbgIdx(code[count+1]))
			count += 2

		case cmp:
			ds[count] = "cmp"
			ds[count+1] = dbgReg(code[count+1])
			ds[count+2] = dbgLst(dbgReg(code[count+2]))
			count += 3
		case cmz:
			ds[count] = "cmz"
			ds[count+1] = dbgLst(dbgReg(code[count+1]))
			count += 2

		case shw:
			ds[count] = "shw"
			ds[count+1] = dbgLst(dbgReg(code[count+1]))
			count += 2
		case get:
			ds[count] = "get"
			ds[count+1] = dbgLst(dbgReg(code[count+1]))
			count += 2
		case dsp:
			ds[count] = "dsp"
			ds[count+1] = dbgLst(dbgIdx(code[count+1]))
			count += 2

		default:
			panic(fmt.Sprintf("Decoding error %f", code[count]))
		}
	}

	return
}

func dbgLit(c float64) (s string) {
	s = fmt.Sprintf("%f", c)
	return
}

func dbgReg(c float64) (s string) {
	s = fmt.Sprintf("%%%.0f", c)
	return
}

func dbgIdx(c float64) (s string) {
	s = fmt.Sprintf("%.0f", c)
	return
}

func dbgLst(msg string) (s string) {
	s = fmt.Sprintf("%s\n", msg)
	return
}
