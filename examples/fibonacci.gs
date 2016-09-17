str	string0	The value of the fibonacci sequence for the
str	string1	-th value of the series is

str	argerr	expected exactly one argument
str	sgnerr	value needs to be non-negative

val	1	%0	Check if received exactly one argument
rac	%1
cmp	%1	%0
jne	argerr

val	0	%1	Get argument and check if it's not negative
rad	%1	%0
flr	%0
cmz	%0
jlt	sgnerr

dsp	string0		Display pre-computation messages
shw	%0
dsp	string1

jeq	retone		Special cases for n = 0 or n = 1
val	1	%1
cmp	%0	%1
jeq	retone

nop			Main computation loop
val	0	%1	a := 1
val	1	%2	b := 1
val	1	%3	count := 2
main:
swp	%1	%2
add	%1	%2
inc	%3
cmp	%3	%0
jne	main

shw	%2
hlt

retone:
val	1	%0
shw	%0
hlt

argerr:
dsp	argerr
hlt

sgnerr:
dsp	sgnerr
hlt
