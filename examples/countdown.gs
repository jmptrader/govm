nop Counts down from integer input


get	%0
val	0	%1

main:
shw	%0
dec	%0
cmp	%0	%1
jgt	main

exit:

val	3.1415	%0
shw	%0
hlt
