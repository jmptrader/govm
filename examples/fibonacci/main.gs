str	welcome0	Fibonacci program.
str	welcome1	Enter a number and I'll tell you Fib(n)
str	result		Fib(n) is
str	newline
str	needargs	fib: Expected exactly one numeric argument.

val	1	%1
rac	%2
cmp	%1	%2
jgt	needargs
jlt	needargs

dsp	welcome0
dsp	welcome1

dec	%1
rad	%1	%0
cel	%0

nop get	%0

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
