package collection_test

import (
	"testing"

	col "github.com/edgarsucre/challenge/src/collection"
)

func TestContains(t *testing.T) {

	data := []int{1, 2, 3, 4}
	collection := col.New(data)

	if collection.Contains(5) {
		t.Fail()
	}

	if !collection.Contains(2) {
		t.Fail()
	}
}

func TestFilter(t *testing.T) {
	data := []int{1, 2, 3, 4}
	collection := col.New(data)

	square := collection.Filter(func(val int) bool {
		return val%2 == 0
	})

	if !square.Contains(2) {
		t.Fail()
	}

	if square.Contains(1) {
		t.Fail()
	}
}

func TestPush(t *testing.T) {
	data := []int{1, 2, 3, 4}
	collection := col.New(data)

	collection.Push(5)

	if !collection.Contains(5) {
		t.Fail()
	}
}

func TestPop(t *testing.T) {
	data := []int{1, 2, 3, 4}
	collection := col.New(data)

	removed := collection.Pop()
	if removed != 4 {
		t.Fail()
	}

	if collection.Contains(4) {
		t.Fail()
	}

	if collection.Size() != 3 {
		t.Fail()
	}

	data2 := []int{1, 2, 3, 4, 4}
	collection2 := col.New(data2)

	collection2.Pop()

	if !collection2.Contains(4) {
		t.Fail()
	}

	collection2.Pop()

	if collection2.Contains(4) {
		t.Fail()
	}

}
