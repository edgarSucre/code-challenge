package array

import (
	"math"
	"sort"
)

func abosolute(c int) int {
	if c < 0 {
		return c * -1
	}
	return c
}

func StoneSmasher(target int, stones []int) int {
	n := len(stones)

	if n == 0 {
		return target
	}

	if target == 0 {
		if n >= 2 {
			return StoneSmasher(stones[n-1]-stones[n-2], stones[:n-2])
		} else {
			return StoneSmasher(abosolute(target-stones[n-1]), []int{})
		}
	} else {
		return StoneSmasher(abosolute(target-stones[n-1]), stones[:n-1])
	}
}

func Directions(steps []string) []int {
	n := len(steps)
	if n == 0 {
		return []int{}
	}

	guide := make(map[int]int)
	guideOrder := make([]int, n)
	r := [2]int{0, int(math.Pow(float64(2), float64(n)))}

	for i, v := range steps {
		media := (r[0] + r[1]) / 2

		if v == "R" {
			r[0] = media
		} else {
			r[1] = media
		}

		j := i + 1
		guide[media] = j
		guideOrder[i] = media
	}

	sort.Ints(guideOrder)
	for i, v := range guideOrder {
		guideOrder[i] = guide[v]
	}

	return guideOrder
}

func BinarySearch(list []int, target int) int {
	if len(list) == 1 {
		if list[0] != target {
			return -1
		}
		return 0
	}

	middle := len(list) / 2
	if list[middle] == target {
		return middle
	}

	if target > list[middle] {
		return BinarySearch(list[middle:], target) + middle
	}

	return BinarySearch(list[:middle], target)
}

func MergeSort(list []int) []int {
	if len(list) == 1 {
		return list
	}

	if len(list) == 2 {
		newList := make([]int, 2)
		if list[0] >= list[1] {
			newList[0], newList[1] = list[1], list[0]
			return newList
		}
		return list
	}

	middle := len(list) / 2
	listA := MergeSort(list[:middle])
	listB := MergeSort(list[middle:])

	merged := mergeInOrder(listA, listB)

	return merged
}

func mergeInOrder(a, b []int) []int {
	if len(a) == 0 {
		return b
	}

	if len(b) == 0 {
		return a
	}

	sorted := make([]int, 1)
	if a[0] <= b[0] {
		sorted[0] = a[0]
		sorted = append(sorted, mergeInOrder(a[1:], b)...)
	} else {
		sorted[0] = b[0]
		sorted = append(sorted, mergeInOrder(a, b[1:])...)
	}
	return sorted
}

// func minVal(a, b int) int {
// 	if a >= b {
// 		return b
// 	}
// 	return a
// }

// func intSliceToString(list []int) string {
// 	var b strings.Builder
// 	for _, v := range list {
// 		fmt.Fprintf(&b, "%d", v)
// 	}

// 	return b.String()
// }

func CountingSort(list []int, max int) []int {
	count := make([]int, max)

	//counting instances
	for _, v := range list {
		count[v]++
	}

	//adding pre-sum & shifting to the right
	for i, sum := 0, 0; i < max; i++ {
		sum, count[i] = sum+count[i], sum
	}

	sorted := make([]int, len(list))
	for _, v := range list {
		sorted[count[v]] = v
		count[v]++
	}

	return sorted
}

func SortUniqueElements(list []int) []int {
	positionIndex := make(map[int]int)
	min := list[0]
	max := list[0]

	for i, v := range list {
		positionIndex[v] = i

		if v < min {
			min = v
		}

		if v > max {
			max = v
		}
	}

	sorted := make([]int, len(list))
	for i, j := 0, min; j <= max; j++ {
		if _, ok := positionIndex[j]; ok {
			sorted[i] = j
			i++
		}
	}

	return sorted
}

func FindCommonSubjects(enrollments [][]string) map[[2]string][]string {
	byStudent := make(map[string][]string)

	for _, enrollment := range enrollments {
		byStudent[enrollment[0]] = append(byStudent[enrollment[0]], enrollment[1])
	}

	intersection := func(a, b []string) []string {
		mapA := make(map[string]bool)
		for _, v := range a {
			mapA[v] = true
		}

		out := make([]string, 0)
		for _, v := range b {
			if _, ok := mapA[v]; ok {
				out = append(out, v)
			}
		}
		return out
	}

	out := make(map[[2]string][]string)

	for studentA, subjectsA := range byStudent {
		for studentB, subjectsB := range byStudent {
			if studentA == studentB {
				continue
			}

			key := [2]string{studentA, studentB}
			if _, ok := out[key]; ok {
				continue
			}

			out[key] = intersection(subjectsA, subjectsB)
		}
	}

	return out
}
