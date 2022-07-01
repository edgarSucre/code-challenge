package numbers_test

import (
	"testing"

	"github.com/edgarsucre/challenge/src/numbers"
)

func TestClosestDivisible(t *testing.T) {
	cases := []map[string]int{
		{
			"n": -15,
			"m": 6,
			"o": -18,
		},
		{
			"n": 13,
			"m": 4,
			"o": 12,
		},
	}

	for _, c := range cases {
		got := numbers.ClosestDivisible(c["n"], c["m"])
		expected := c["o"]
		if got != expected {
			t.Errorf("expected: %v, got: %v", expected, got)
		}
	}

}

func TestRoundToNearestMultiple(t *testing.T) {
	cases := []map[string]int{
		{
			"n": 152,
			"m": 10,
			"o": 150,
		},
		{
			"n": 7,
			"m": 10,
			"o": 10,
		},
		{
			"n": 20,
			"m": 10,
			"o": 20,
		},
	}

	for _, c := range cases {
		got := numbers.RoundToNearestMultiple(c["n"], c["m"])
		expected := c["o"]
		if got != expected {
			t.Errorf("expected: %v, got: %v", expected, got)
		}
	}
}

func TestSumOfNumbers(t *testing.T) {
	o := numbers.SumOfNumbers(10)
	if o != 55 {
		t.Fail()
	}
}

func TestFibonacci(t *testing.T) {
	o := numbers.Fibonacci(10)
	if o != 55 {
		t.Fail()
	}

	o = numbers.Fibonacci(32)
	if o != 2178309 {
		t.Fail()
	}
}

func TestLongestSum(t *testing.T) {
	r := numbers.LongestSum([]int{-3, 5, -20, 10, 1, 5, -30})
	if r != 16 {
		t.Fail()
	}
}

func TestLongestSumIndices(t *testing.T) {
	list := []int{-3, 5, -20, 10, 1, 5, -30}
	x, y := numbers.LongestSumIndices(list)
	if x != 3 && y != 5 {
		t.Fail()
	}
}
