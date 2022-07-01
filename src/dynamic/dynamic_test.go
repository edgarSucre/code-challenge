package dynamic_test

import (
	"testing"

	"github.com/edgarsucre/challenge/src/dynamic"
)

func TestTribonacci(t *testing.T) {
	r := dynamic.Tribonnacci(5)

	if r != 4 {
		t.Fail()
	}
}

func TestSumPosible(t *testing.T) {
	cases := []struct {
		target int
		list   []int
		result bool
	}{
		{8, []int{5, 12, 4}, true},
		{15, []int{6, 2, 10, 19}, false},
		{13, []int{6, 2, 1}, true},
		{103, []int{6, 20, 1}, true},
		{12, []int{}, false},
		{12, []int{12}, true},
		{0, []int{}, true},
		{271, []int{10, 8, 265, 24}, false},
		{2017, []int{4, 2, 10}, false},
		{13, []int{3, 5}, true},
	}

	for _, testCase := range cases {
		got := dynamic.SumPosible(testCase.target, testCase.list)
		if got != testCase.result {
			t.Errorf("Case %d-%v failed, expected %v, got %v", testCase.target, testCase.list, testCase.result, got)
		}
	}

}

func TestMinChange(t *testing.T) {

	cases := []struct {
		target int
		list   []int
		result int
	}{
		{8, []int{1, 5, 4, 12}, 2},
		{13, []int{1, 9, 5, 14, 30}, 5},
		{23, []int{2, 5, 7}, 4},
		{102, []int{1, 5, 10, 25}, 6},
		// {200, []int{1, 5, 10, 25}, 8},
		// {2017, []int{4, 2, 10}, -1},
		// {271, []int{10, 8, 265, 24}, -1},
		{0, []int{4, 2, 10}, 0},
		{0, []int{}, 0},
	}

	for _, testCase := range cases {
		got := dynamic.MinChange(testCase.target, testCase.list)
		if got != testCase.result {
			t.Errorf("Case %d-%v failed, expected %v, got %v", testCase.target, testCase.list, testCase.result, got)
		}
	}
}

func TestCountPaths(t *testing.T) {
	cases := []struct {
		grid   [][]string
		result int
	}{
		{
			[][]string{
				{"O", "O"},
				{"O", "O"},
			},
			2,
		},

		{
			[][]string{
				{"O", "O", "X"},
				{"O", "O", "O"},
				{"O", "O", "O"},
			},
			5,
		},

		{
			[][]string{
				{"O", "O", "O"},
				{"O", "O", "X"},
				{"O", "O", "O"},
			},
			3,
		},

		{
			[][]string{
				{"O", "O", "O"},
				{"O", "X", "X"},
				{"O", "O", "O"},
			},
			1,
		},

		{
			[][]string{
				{"O", "O", "X", "O", "O", "O"},
				{"O", "O", "X", "O", "O", "O"},
				{"X", "O", "X", "O", "O", "O"},
				{"X", "X", "X", "O", "O", "O"},
				{"O", "O", "O", "O", "O", "O"},
			},
			0,
		},

		{
			[][]string{
				{"O", "O", "X", "O", "O", "O"},
				{"O", "O", "O", "O", "O", "X"},
				{"X", "O", "O", "O", "O", "O"},
				{"X", "X", "X", "O", "O", "O"},
				{"O", "O", "O", "O", "O", "O"},
			},
			42,
		},

		{
			[][]string{
				{"O", "O", "X", "O", "O", "O"},
				{"O", "O", "O", "O", "O", "X"},
				{"X", "O", "O", "O", "O", "O"},
				{"X", "X", "X", "O", "O", "O"},
				{"O", "O", "O", "O", "O", "X"},
			},
			0,
		},

		{
			[][]string{
				{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
				{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
				{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
				{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
				{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
				{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
				{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
				{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
				{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
				{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
				{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
				{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
				{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
				{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
				{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
			},
			40116600,
		},

		{
			[][]string{
				{"O", "O", "X", "X", "O", "O", "O", "X", "O", "O", "O", "O", "O", "O", "O"},
				{"O", "O", "X", "X", "O", "O", "O", "X", "O", "O", "O", "O", "O", "O", "O"},
				{"O", "O", "O", "X", "O", "O", "O", "X", "O", "O", "O", "O", "O", "O", "O"},
				{"X", "O", "O", "O", "O", "O", "O", "X", "O", "O", "O", "O", "O", "O", "O"},
				{"X", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
				{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "X", "X", "O"},
				{"O", "O", "O", "O", "O", "O", "O", "X", "O", "O", "O", "O", "O", "X", "O"},
				{"O", "O", "O", "O", "O", "O", "O", "O", "X", "O", "O", "O", "O", "O", "O"},
				{"X", "X", "X", "O", "O", "O", "O", "O", "O", "X", "O", "O", "O", "O", "O"},
				{"O", "O", "O", "O", "X", "X", "O", "O", "O", "O", "X", "O", "O", "O", "O"},
				{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
				{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
				{"O", "O", "O", "O", "X", "X", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
				{"O", "O", "O", "O", "O", "O", "O", "O", "X", "O", "O", "O", "O", "O", "O"},
				{"O", "O", "O", "O", "O", "O", "O", "O", "X", "O", "O", "O", "O", "O", "O"},
			},
			3190434,
		},
	}

	for i, testCase := range cases {
		got := dynamic.CountPaths(testCase.grid)
		if got != testCase.result {
			t.Errorf("TestCase %d failed, Expected: %d, Got: %d", i, testCase.result, got)
		}
	}
}

func TestMaxPathSum(t *testing.T) {
	cases := []struct {
		grid   [][]int
		result int
	}{
		{
			[][]int{
				{1, 3, 12},
				{5, 1, 1},
				{3, 6, 1},
			},
			18,
		},

		{
			[][]int{
				{1, 2, 8, 1},
				{3, 1, 12, 10},
				{4, 0, 6, 3},
			},
			36,
		},

		{
			[][]int{
				{1, 2, 8, 1},
				{3, 10, 12, 10},
				{4, 0, 6, 3},
			},
			39,
		},

		{
			[][]int{
				{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			},
			27,
		},

		{
			[][]int{
				{1, 1, 3, 1, 1, 1, 1, 4, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				{1, 2, 1, 1, 6, 1, 1, 5, 1, 1, 0, 0, 1},
				{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 5, 1, 1, 1, 1, 0, 1, 1, 1, 1},
				{2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				{2, 1, 1, 1, 1, 8, 1, 1, 1, 1, 1, 1, 1},
				{2, 1, 3, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 9, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 3, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1, 1, 8, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			},
			56,
		},
	}

	for i, testCase := range cases {
		got := dynamic.MaxPathSum(testCase.grid)
		if got != testCase.result {
			t.Errorf("TestCase %d failed, Expected: %d, Got: %d", i, testCase.result, got)
		}
	}
}

func TestNonAdjacentSum(t *testing.T) {
	r := dynamic.NonAdjacentSum([]int{72, 62, 10, 6, 20, 19, 42, 46, 24, 78,
		30, 41, 75, 38, 23, 28, 66, 55, 12, 17,
		83, 80, 56, 68, 6, 22, 56, 96, 77, 98,
		61, 20, 0, 76, 53, 74, 8, 22, 92, 37,
		30, 41, 75, 38, 23, 28, 66, 55, 12, 17,
		72, 62, 10, 6, 20, 19, 42, 46, 24, 78,
		42})
	if r != 1465 {
		t.Fail()
	}
}
