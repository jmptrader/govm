str	newline
str	separator	---------------------------------------------

nop	-------------------------------------------------------------
nop	STACK_GS

str	stack_in	stack.gs test
str	stack_msg	Output should be 8, 8, 8, 5, 5, 8

dsp	separator
dsp	stack_in
dsp	newline

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

dsp	stack_msg
dsp	newline

nop	-------------------------------------------------------------
nop	FUNCTION_GS

str	fn_in		function.gs test
str	fn_explanation	This program tests function calls.
str	fn_expected	You should see 5 4 3 2 1, this text and 8 7 6 5 4 3 2 1.

dsp	separator
dsp	fn_in
dsp	newline
dsp	fn_explanation

val	5	%0
psh	%0
cll	countdownAndPrint

dsp	fn_expected

val	8	%0
psh	%0
cll	countdownAndPrint
dsp	newline

hlt

lab	countdownAndPrint
pop	%0
lab	capLoop
shw	%0
dec	%0
cmz	%0
jgt	capLoop
ret
