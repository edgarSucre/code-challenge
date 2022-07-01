package dynamic

import (
	"fmt"
	"math"
	"strings"
)

func Tribonnacci(n int) int {
	var dpt func(int, map[int]int) int

	dpt = func(i int, memo map[int]int) int {
		if val, ok := memo[i]; ok {
			return val
		}

		if i < 2 {
			return 0
		}

		if i < 3 {
			return 1
		}

		r := dpt(i-3, memo) + dpt(i-2, memo) + dpt(i-1, memo)
		memo[i] = r

		return r
	}

	memo := make(map[int]int)
	return dpt(n, memo)
}

func SumPosible(target int, options []int) bool {
	var dpt func([]int, int, int, map[int]bool) bool
	dpt = func(list []int, t int, sum int, memo map[int]bool) bool {

		if _, ok := memo[sum]; ok {
			return false
		}

		if sum > t {
			return false
		}

		if sum == t {
			return true
		}

		for _, v := range list {
			if dpt(list, t, sum+v, memo) {
				return true
			}
		}

		memo[sum] = false

		return false
	}

	memo := make(map[int]bool)
	return dpt(options, target, 0, memo)
}

func MinChange(target int, n []int) int {

	var dps func([]int, int, int, map[int]bool) int
	dps = func(list []int, t int, sum int, visited map[int]bool) int {
		if _, ok := visited[sum]; ok {
			return -1
		}

		if sum > t {
			return -1
		}

		if sum == t {
			return 0
		}

		max := 99999999999
		min := max
		for _, v := range list {
			r := dps(list, t, sum+v, visited)
			if r > -1 && r < min {
				min = r
			}
		}

		if min < max {
			return min + 1
		}

		visited[sum] = false
		return -1
	}

	visited := make(map[int]bool)
	return dps(n, target, 0, visited)
}

func CountPaths(grid [][]string) int {
	memo := make(map[string]int)

	return countPathsDPS(grid, 0, 0, memo)
}

func countPathsDPS(grid [][]string, x, y int, memo map[string]int) int {
	if x > len(grid)-1 || y > len(grid[0])-1 {
		return 0
	}

	key := fmt.Sprintf("%d,%d", x, y)
	if f, ok := memo[key]; ok {
		return f
	}

	if grid[x][y] == "X" {
		return 0
	}

	if x == len(grid)-1 && y == len(grid[0])-1 {
		return 1
	}

	r := countPathsDPS(grid, x+1, y, memo) + countPathsDPS(grid, x, y+1, memo)
	memo[key] = r

	return r
}

func MaxPathSum(grid [][]int) int {
	memo := make(map[string]int)

	return maxPathSumDPT(grid, 0, 0, memo)
}

func maxPathSumDPT(grid [][]int, x, y int, memo map[string]int) int {
	if x > len(grid)-1 || y > len(grid[0])-1 {
		return 0
	}

	key := fmt.Sprintf("%d,%d", x, y)
	if f, ok := memo[key]; ok {
		return f
	}

	if x == len(grid)-1 && y == len(grid[0])-1 {
		return grid[x][y]
	}

	fromX := maxPathSumDPT(grid, x+1, y, memo)
	fromY := maxPathSumDPT(grid, x, y+1, memo)

	memo[key] = int(math.Max(float64(fromX), float64(fromY))) + grid[x][y]

	return memo[key]
}

func NonAdjacentSum(l []int) int {
	max := 0
	memo := make(map[string]int)
	for i, v := range l {
		r := NonAdjacenSumtDPT(getRemainingOptions(l, i), memo) + v
		if r > max {
			max = r
		}
	}

	return max
}

func NonAdjacenSumtDPT(l []int, memo map[string]int) int {
	key := getKey(l)
	if _, ok := memo[key]; ok {
		return memo[key]
	}

	max := 0
	for i, v := range l {
		r := NonAdjacenSumtDPT(getRemainingOptions(l, i), memo) + v
		if r > max {
			max = r
		}
	}

	memo[key] = max
	return max
}

func getKey(l []int) string {
	o := make([]string, len(l))
	for i, v := range l {
		o[i] = fmt.Sprint(v)
	}

	return strings.Join(o, "-")

}

func getRemainingOptions(l []int, t int) []int {
	r := make([]int, 0)
	if len(l) <= 2 {
		return r
	}
	if t == 0 {
		return l[2:]
	}
	if t == len(l)-1 {
		return l[:t-1]
	}
	r = append(r, l[:t-1]...)
	r = append(r, l[t+2:]...)
	return r
}
