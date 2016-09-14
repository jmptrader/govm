str	newline		
str	welcome1	Welcome to the `choose` program.
str	welcome2	It was written in vsart/GoVM assembly language.
str	prompt1		Input the 1st value of C(x,y)
str	prompt2		Input the 2nd value of C(x,y)
str	result		The result is
str	error		Invalid first value.


dsp	newline
dsp	welcome1
dsp	welcome2
dsp	newline
dsp	newline

val	0	%0

dsp	prompt1
get	%1
cel	%1
cmp	%1	%0	Do not allow negative values.
jlt	error

dsp	prompt2
get	%2
cel	%2

dsp	newline

cmp	%2	%0
jlt	zero
jeq	one
cmp	%2	%1
jgt	zero
jeq	one

nop	Calculate x-y into %3
cpy	%1	%3
sub	%2	%3

val	1	%4

nop	Check if x-y > y, if so then swap %3 with %2
cmp	%2	%3
jlt	main

cpy	%2	%7
cpy	%3	%2
cpy	%7	%3

lab	main

mul	%1	%4
div	%3	%4

dec	%1
dec	%2
dec	%3
cmp	%3	%0
jgt	main

dsp	result
shw	%4

jmp	exit


lab	one
dsp	result
val	1	%0
shw	%0
jmp	exit


lab	zero
dsp	result
shw	%0
jmp	exit


lab	error
dsp	error
jmp	exit

lab	exit
dsp	newline
hlt
