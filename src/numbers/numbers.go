package numbers

/*
 * Find the closest number to n, that is also divisible by m. Can be positive or negative
 * In case of having two numbers, return the absulte bigger.
 * Example: n=15, m=6 => 18. Posible solutions are 12 and -18; -18 having the bigger
 * absolute value.
 */
func ClosestDivisible(n, m int) int {

	i, j := n, n

	for {
		i += 1
		j -= 1
		if i%m == 0 && j%m == 0 {
			if n < 0 {
				return j
			}
			return i
		} else if i%m == 0 {
			return i
		} else if j%m == 0 {
			return j
		}
	}
}

func RoundToNearestMultiple(n, m int) int {
	low := n / m * m
	high := low + m

	if n-low > high-n {
		return high
	}
	return low
}

func SumOfNumbers(n int) int {
	if n == 1 {
		return 1
	}

	return SumOfNumbers(n-1) + n
}

func Fibonacci(n int) int {
	var fib func(int) int

	mem := make(map[int]int)
	fib = func(in int) int {
		if in <= 2 {
			return 1
		}

		if v, ok := mem[in]; ok {
			return v
		}

		ans := fib(in-1) + fib(in-2)
		mem[in] = ans

		return ans
	}

	return fib(n)
}

func LongestSum(n []int) int {
	longest := n[0]
	for i, j := 0, 1; j < len(n); j++ {
		sum := sumArray(n[i : j+1])
		if sum > longest {
			longest = sum
		} else {
			i = j
		}
	}

	return longest
}

func LongestSumIndices(n []int) (int, int) {
	longest, start, end := n[0], 0, 0
	for i, j := 0, 1; j < len(n)-1; j++ {
		sum := sumArray(n[i : j+1])
		if sum > longest {
			start = i
			end = j
			longest = sum
		} else {
			i = j
		}
	}

	return start, end
}

func sumArray(n []int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}
	return sum
}
