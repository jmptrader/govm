str	unif	Uniform Distribution
str	norm	Normal Distribution

now	%0
sed	%0

dsp	unif
val	10	%1
unif:
rnd	%0
shw	%0
dec	%1
cmz	%1
jne	unif

dsp	norm
val	10	%1
norm:
rnn	%0
shw	%0
dec	%1
cmz	%1
jne	norm

hlt
