str	result0		The value of the Fibonacci series for the
str	result1		-th element is
str	newline
str	needargs	fib: Expected exactly one numeric argument.
str	positive	fib: Only accepts non-negative values.

val	1	%1
rac	%2
cmp	%1	%2
jgt	needargs
jlt	needargs

dec	%1
rad	%1	%0
cel	%0

cmz	%0
jlt	positive

dsp	result0
shw	%0
dsp	result1

psh	%0
cll	fib
shw	%0
dsp	newline
hlt

needargs:
dsp	needargs
hlt

positive:
dsp	positive
hlt
