str	newline		
str	welcome1	Welcome to the `mean` program.
str	welcome2	It was written in vsart/GoVM assembly language.
str	welcome3	Use Ctrl+C to stop execution.
str	prompt		Add a new number
str	result		The current mean is

dsp	newline
dsp	welcome1
dsp	welcome2
dsp	welcome3
dsp	newline

dsp	prompt
get	%5		This will hold the mean
val	1	%6	This will hold the number of values

lab	mainLoop

cll	clearScreen
dsp	result
shw	%5
dsp	prompt
get	%1

mul	%6	%5	mean := mean * n
add	%1	%5	mean := mean + new
inc	%6		n := n + 1
div	%6	%5	mean := mean / n

jmp	mainLoop
hlt

def	clearScreen
val	40	%0
lab	clearScreenLoop
dsp	newline
dec	%0
cmz	%0
jgt	clearScreenLoop
ret
