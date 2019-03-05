package try_test

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	// a := 1
	// b := 1
	a, b := 1, 1
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	tmp := a
	a = b
	b = tmp

	t.Log(a, b)
}
