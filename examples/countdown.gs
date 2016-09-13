nop Counts down from integer input


get	%0
val	0	%1

lab	main
shw	%0
dec	%0
cmp	%0	%1
jgt	main

lab	exit

val	3.1415	%0
shw	%0
hlt
