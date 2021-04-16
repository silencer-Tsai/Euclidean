package Euclidean

import "testing"

func TestGCD(t *testing.T) {
	tests := []struct {
		x      int64
		y      int64
		result int64
		err    error
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
		result, trace, err := gcd(test.x, test.y, trace)
		if err != nil {
			if err != test.err {
				t.Errorf("different error, want: %s, get: %s", test.err, err)
			}
			continue
		}

		if result != test.result {
			t.Errorf("different result, want: %d, get: %d", test.result, result)
			continue
		}
		t.Logf("gcd: %d", result)
		t.Log("trace: ", trace)

		// bezout
		a, b, _ := bezout(trace)
		t.Log(a)
		t.Log(b)
		r := a*test.x + b*test.y
		if r != test.result {
			t.Errorf("failed to find a and b, a: %d, b: %d, want: %d, get: %d", a, b, test.result, r)
		}
	}
}
