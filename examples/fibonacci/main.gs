str	welcome0	Fibonacci program.
str	welcome1	Enter a number and I'll tell you Fib(n)
str	result		Fib(n) is
str	newline
str	needargs	fib: Expected exactly one numeric argument.

nop	val	1	%1
nop	rac	%2
nop	cmp	%1	%2
nop	jgt	needargs
nop	jlt	needargs

dsp	welcome0
dsp	welcome1

nop	dec	%1
nop	rad	%1	%0
nop	cel	%0

get	%0

psh	%0
cll	fib
dsp	newline
dsp	result
shw	%0
dsp	newline
hlt

needargs:
dsp	needargs
hlt
