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





fib:
pop	%0

cmz	%0
jeq	fibZero
val	1	%1
cmp	%0	%1
jeq	fibOne

dec	%0		R0 := n-1
psh	%0		STACK: [n-1 | ...]
psh	%0		STACK: [n-1 | n-1 | ...]
cll	fib		STACK: [n-1 | ...] ; R0 := fib(n-1)
pop	%1		STACK: [...] R1 := n-1
dec	%1		R1 := n-2
psh	%0		STACK: [fib(n-1) | ...]
psh	%1		STACK: [fib(n-1) | n-2 | ...]
cll	fib		STACK: [fib(n-1) | ...] ; R0 := fib(n-2)
pop	%1		R1 := fib(n-1)
add	%1	%0	R0 := fib(n-1) + fib(n-2)
ret

fibZero:
val	0	%0
ret

fibOne:
val	1	%0
ret
