package Euclidean

import "fmt"

var (
	ErrInvalidNum   = fmt.Errorf("invalid number")
	ErrInvalidTrace = fmt.Errorf("invalid trace")
)

// calculate greatest common divisor of x and y
// using euclidean algorithm
// return GCD of x and y and the quotients during calculating
func gcd(x, y int64, trace []int64) (int64, []int64, error) {
	if x == 0 || y == 0 {
		return 0, nil, ErrInvalidNum
	}

	q := x / y
	r := x % y
	trace = append(trace, q)
	if r == 0 {
		return y, trace, nil
	} else {
		return gcd(y, r, trace)
	}
}

// calculate the integer tuple (a, b) satisfying a*x + b*y = gcd(x, y)
func bezout(trace []int64) (int64, int64, error) {
	l := len(trace)
	if l == 0 {
		return 0, 0, ErrInvalidTrace
	}

	// x|y
	if l == 1 {
		return 1, -trace[l-1] + 1, nil
	}

	// x < y && x|y
	if l == 2 && trace[0] == 0 {
		return -trace[l-1] + 1, 1, nil
	}

	// !(x|y)
	a := int64(1)
	b := -trace[l-2]
	for i := l - 2; i > 0 && trace[i-1] != 0; i-- {
		a, b = b, a+b*(-trace[i-1])
	}

	if trace[0] == 0 {
		// x < y
		return b, a, nil
	} else {
		// x > y
		return a, b, nil
	}
}
