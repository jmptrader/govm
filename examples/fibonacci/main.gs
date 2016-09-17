str	welcome0	Fibonacci program.
str	welcome1	Enter a number and I'll tell you Fib(n)
str	result		Fib(n) is
str	newline

dsp	welcome0
dsp	welcome1
get	%0
psh	%0
cll	fib
dsp	newline
dsp	result
shw	%0
dsp	newline
hlt
