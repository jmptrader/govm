str	msg	Output should be 8, 8, 8, 5, 5, 8

val	5	%0
val	8	%1

psh	%1
psh	%0
psh	%0
psh	%1
psh	%1
psh	%1

val 	6	%3

lab	popAndShow

pop	%2
shw	%2

dec	%3
cmz	%3
jgt	popAndShow

dsp	msg
hlt
