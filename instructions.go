package main

// Dictionary
// [lit] - Double literal
// [reg] - Register
// [add] - Code address

// Operations always have the last argument as destination

// ATTENTION: hlt being equal to 0 is specified, anything else isn't
// This is so the golang spec of initializing things to 0 will initialize
// every instruction to a halt. Could've decided to be a nop instead, but hey.

const (
	hlt = iota // hlt				- Stops program execution

	val // float [lit] [reg] - Pushes value into register
	cpy // cpy [reg] [reg] - Copies the value from first register to second

	add // add [reg] [reg] - Adds first value to second
	sub // sub [reg] [reg] - Subtracts first value from second
	mul // mul [reg] [reg] - Multiplies the first value with the second

	shw // shw [reg] - Prints the value of reg to stdout
	get // get [reg] - Reads double from stdin into reg
	nop // nop - Does nothing

	flr // flr [reg] - Rounds the value of reg to floor(reg)
	cel // cel [reg] - Rounds the value of reg to ceil(reg)
	inc // inc [reg] - Increments 1 to the value of reg
	dec // dec [reg] - Decrements 1 to the value of dec

	jmp // jmp [add] - Jumps to code[add]

	cmp // cmp [reg] [reg] - Special operation
	jlt // jlt [add] - Special operation
	jeq // jeq [add] - Special operation
	jgt // jgt [add] - Special operation
)
