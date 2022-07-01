package strings

import (
	"fmt"
	"sort"
	"strings"
)

/*
 * Replace consecutive character repetitions by the sum of the repetitions and the
 * character: in = SBBK => S2BK
 *
 */
func ReplaceRepetition(s string) string {
	ins := strings.Split(s, "")
	counts := make([]int, len(ins))

	var index int

	for i, v := range ins {
		if i == 0 {
			counts[i]++
			continue
		}

		if v == ins[index] {
			counts[index]++
		} else {
			counts[i]++
			index = i
		}
	}

	var out string
	for i, v := range ins {
		t := counts[i]
		if t == 0 {
			continue
		} else if t == 1 {
			out = fmt.Sprintf("%s%s", out, v)
		} else {
			out = fmt.Sprintf("%s%v%s", out, t, v)
		}
	}

	return out
}

func SortCharacters(s string) string {
	sl := strings.Split(s, "")
	sort.Strings(sl)
	s = strings.Join(sl, "")

	return s
}

func AreAnagrams(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	asl := strings.Split(a, "")
	bsl := strings.Split(b, "")

	sort.Strings(asl)
	sort.Strings(bsl)

	a = strings.Join(asl, "")
	b = strings.Join(bsl, "")

	return a == b
}

func MakeAnagram(a, b string) int {
	if AreAnagrams(a, b) {
		return 0
	}

	asl := strings.Split(a, "")
	bsl := strings.Split(b, "")
	countA := make(map[string]int)
	countB := make(map[string]int)
	guide := make(map[string]interface{})

	for _, v := range asl {
		countA[v]++
		guide[v] = ""
	}

	for _, v := range bsl {
		countB[v]++
		guide[v] = ""
	}

	var sum int
	for k := range guide {
		r := countA[k] - countB[k]
		if r == 0 {
			continue
		} else if r > 0 {
			sum += r
		} else {
			r = r * -1
			sum += r
		}
	}

	return sum
}

//is valid id all letter have the same frecuency
func IsValid(s string) string {
	// Write your code here
	sls := strings.Split(s, "")
	counter := make(map[string]int)

	for _, v := range sls {
		counter[v]++
	}

	duplicateCounter := make(map[int]int)
	for _, v := range counter {
		duplicateCounter[v]++
	}

	if len(duplicateCounter) > 2 {
		return "NO"
	} else if len(duplicateCounter) == 1 {
		return "YES"
	} else {
		for i, v := range duplicateCounter {
			if v == 1 {
				if i == 1 || duplicateCounter[i-1] > 0 {
					return "YES"
				}
			}
		}
		return "NO"
	}
}

func IsPalindrome(s string) bool {
	return s == InvertStringRecursive(s)
}

func CountSpecialString(s string, n int) int {
	subs := strings.Split(s, "")

	// sliceIsHomogeneous := func(z []string) bool {
	// 	a := strings.Join(z, "")
	// 	a = strings.TrimSpace(strings.Join(strings.Split(a, z[0]), ""))
	// 	return len(a) == 0
	// }

	// slicesAreEqual := func(x, y []string) bool {
	// 	return sliceIsHomogeneous(x) && sliceIsHomogeneous(y) && x[0] == y[0]
	// }

	// isSpecial := func(c []string) bool {
	// 	m := len(c)
	// 	z := m / 2
	// 	x := c[0:z]
	// 	var y []string

	// 	if m%2 == 0 {
	// 		y = c[z:]
	// 	} else {
	// 		y = c[z+1:]
	// 	}
	// 	return slicesAreEqual(x, y)
	// }

	var sum int
	for i := 2; i <= n; i++ {
		m := i / 2
		for j := 0; j < m; j++ {
			a := j
			b := a + 1

			if i == 2 {
				if subs[a] != subs[b] {
					continue
				}
				sum++
			} else if i == 3 {
				if i%2 == 0 {
					b = a + 2
					if subs[a] != subs[b] {
						continue
					}
					sum++
				}
			} else {

				var (
					c int
					d int
				)
				if i%2 == 0 {
					c = b + 1
					d = c + i
				} else {
					c = b + 2
					d = c + 1
				}

				if subs[a] == subs[b] && subs[a] == subs[c] && subs[b] == subs[d] {
					sum++
				}
			}
		}
	}

	return sum + len(subs)
}

func InvertStringRecursive(s string) string {
	if s == "" {
		return s
	}

	return InvertStringRecursive(s[1:]) + s[0:1]
}

func InvertStringLoop(s string) string {
	list := strings.Split(s, "")
	out := make([]string, len(list))

	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = list[j], list[i]
	}

	return strings.Join(out, "")
}

func ConvertDecimalToBinary(d int) string {
	var rec func(int, string) string

	rec = func(in int, result string) string {

		if in == 0 {
			return result
		}

		out := fmt.Sprintf("%v%s", in%2, result)
		return rec(in/2, out)
	}

	return rec(d, "")
}
