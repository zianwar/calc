# calc
Command line utility for advanced math operations built with [Go](https://golang.org).  
Uses the [https://newton.now.sh/](https://newton.now.sh/) Math API under the hood.

### Installation:
1. Make sure you have [Go](http://golang.org/doc/install.html) installed
2. Install calc:
```
$ go get github.com/Zianwar/calc
```

### Usage:

```
Usage: calc [operation] [expression]

  • Where operation is defined as:

	 derive:     Derives expression.        eg:    x^2 + 2x   =>             2x + 2
	 zeroes:     Finds expression zeros.    eg:    x^2 + 2x   =>            [-2, 0]
	 cos:        Cosine of number.          eg:          pi   =>                 -1
	 sin:        Sine of number.            eg:           0   =>                  0
	 arccos:     Inverse Cosine of number   eg:           1   =>                  0
	 arctan:     Inverse Tangent of number  eg:           0   =>                  0
	 abs:        Absolute Value of number   eg:          -1   =>                  1
	 factor:     Factorise expression.      eg:    x^2 + 2x   =>           x(x + 2)
	 integrate:  integrates expression.      eg:    x^2 + 2x  =>  1/3 x^3 + x^2 + C
	 tan:        Tangent of number.         eg:           0   =>                  0
	 arcsin:     Inverse Sine of number     eg:           0   =>                  0
	 log:        Logarithm                  eg:         218   =>                  3
	 simplify:   Simplifies expression.     eg:  2^2 + 2(2)   =>                  8

  • And expression could be any number or mathematical equation.
```

### Examples

```
$ calc integrate "x"
  1/2 x^2

$ calc simplify "1+2-2"
  1
```
