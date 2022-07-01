package array_test

import (
	"fmt"
	"testing"

	"github.com/edgarsucre/challenge/src/array"
)

func TestStoneSmasher(t *testing.T) {
	res := array.StoneSmasher(0, []int{1, 2, 3, 6, 7, 7})
	if res != 0 {
		t.Fail()
	}
}

func TestDirections(t *testing.T) {
	res := array.Directions([]string{"L", "R", "L", "R", "R", "L"})
	if res == nil {
		t.Fail()
	}
}

func TestBinarySearch(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, v := range list {
		ans := array.BinarySearch(list, v)
		if ans != i {
			t.Fail()
		}
	}
}

func TestMergeSort(t *testing.T) {
	o := array.MergeSort([]int{2, 1, 8, 9, 7})
	if o[0] != 1 {
		t.Fail()
	}
}

func TestCountingSort(t *testing.T) {
	o := array.CountingSort([]int{2, 1, 8, 9, 7}, 10)
	if o[0] != 1 {
		t.Fail()
	}
}

func TestSortUniqueElements(t *testing.T) {
	o := array.SortUniqueElements([]int{2, 1, 8, 9, 7})
	if o[0] != 1 {
		t.Fail()
	}
}

func TestFindCommonSubjects(t *testing.T) {
	enrollments := [][]string{
		{"02", "matematica"},
		{"85", "espanol"},
		{"30", "economia"},
		{"30", "matematica"},
	}

	o := array.FindCommonSubjects(enrollments)
	fmt.Println(o)
}
