package Euclidean

import "testing"

func TestGCD(t *testing.T) {
	tests := []struct {
		x   int64
		y   int64
		gcd int64
		err error
	}{
		{9, 0, 0, ErrInvalidNum},
		{15, 15, 15, nil},
		{47, 30, 1, nil},
		{97, 2659, 1, nil},
		{195, 77, 1, nil},
		{20, 100, 20, nil},
		{624129, 2061517, 18913, nil},
	}
	for _, test := range tests {
		t.Log("--------------------------------")
		t.Logf("x: %d, y: %d\n", test.x, test.y)
		gcd, err := GCD(test.x, test.y)
		if err != nil {
			if err != test.err {
				t.Errorf("different error, want: %s, get: %s", test.err, err)
			}
			continue
		}
		if gcd != test.gcd {
			t.Errorf("unexpected gcd, want: %d, get: %d", test.gcd, gcd)
			continue
		}
		t.Logf("gcd: %d", gcd)
	}
}

func TestGCDWithTrace(t *testing.T) {
	tests := []struct {
		x   int64
		y   int64
		gcd int64
		err error
	}{
		{9, 0, 0, ErrInvalidNum},
		{15, 15, 15, nil},
		{47, 30, 1, nil},
		{97, 2659, 1, nil},
		{195, 77, 1, nil},
		{20, 100, 20, nil},
		{624129, 2061517, 18913, nil},
	}
	for _, test := range tests {
		t.Log("--------------------------------")
		t.Logf("x: %d, y: %d\n", test.x, test.y)
		trace := make([]int64, 0)
		gcd, trace, err := GCDWithTrace(test.x, test.y, trace)
		if err != nil {
			if err != test.err {
				t.Errorf("unexpected error, want: %s, get: %s", test.err, err)
			}
			continue
		}

		if gcd != test.gcd {
			t.Errorf("unexpected gcd, want: %d, get: %d", test.gcd, gcd)
			continue
		}
		t.Logf("gcd: %d", gcd)
		t.Log("trace: ", trace)

		// Bezout
		a, b, _ := Bezout(test.x, test.y)
		t.Log(a)
		t.Log(b)

		// not safe to calculate a*x+b*y here
		// there is a potential for overflow
		r := a*test.x + b*test.y
		if r != test.gcd {
			t.Errorf("failed to find a and b, a: %d, b: %d, want: %d, get: %d", a, b, test.gcd, r)
		}
	}
}
