str	explanation	This program tests function calls.
str	expected	You should see 5 4 3 2 1, this text and 8 7 6 5 4 3 2 1.

dsp	explanation

val	5	%0
psh	%0
cll	countdownAndPrint
dsp	expected
val	8	%0
psh	%0
cll	countdownAndPrint
hlt

def	countdownAndPrint
pop	%0
lab	capLoop
shw	%0
dec	%0
cmz	%0
jgt	capLoop
ret
