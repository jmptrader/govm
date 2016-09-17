str	intro	This file tests the instructions implemented in PC010.
str	expect0	Should be approximately 1.21
str	expect1	Should be approximately 1.4142
str	expect2	Should be approximately 0.14112
str	expect3	Should be approximately -0.9899
str	expect4	Should be approximately -0.1425
str	expect5	Should be approximately 0.1353
str	expect6	Should be approximately -0.1053
str	expect7	Should be approximately 7.756
str	expect8	Should be approximately 9.974
str	expect9	Should be approximately 10.761

val	-1.21	%0
abs	%0
shw	%0
dsp	expect0

val	2	%0
sqr	%0
shw	%0
dsp	expect1

val	3	%0
sin	%0
shw	%0
dsp	expect2

val	3	%0
cos	%0
shw	%0
dsp	expect3

val	3	%0
tan	%0
shw	%0
dsp	expect4

val	-2	%0
exp	%0
shw	%0
dsp	expect5

val	0.9	%0
log	%0
shw	%0
dsp	expect6

val	4.2	%0
gam	%0
shw	%0
dsp	expect7

val	3.1	%0
val	2.1	%1
pow	%0	%1
shw	%1
dsp	expect8

val	3.1	%0
val	2.1	%1
pwa	%0	%1
shw	%1
dsp	expect9

hlt
