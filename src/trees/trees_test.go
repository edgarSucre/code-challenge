package trees_test

import (
	"testing"

	"github.com/edgarsucre/challenge/src/trees"
)

func TestExistInTree(t *testing.T) {
	d := &trees.Node{Val: 4}
	e := &trees.Node{Val: 5}
	f := &trees.Node{Val: 6}
	b := &trees.Node{Val: 2, Left: d, Right: e}
	c := &trees.Node{Val: 3, Right: f}
	a := &trees.Node{Val: 1, Left: b, Right: c}

	ans := trees.ExistsInTreeBFS(a, 2)
	if !ans {
		t.Fail()
	}

	ans = trees.ExistsInTreeDFS(a, 8)
	if ans {
		t.Fail()
	}
}

func TestSumOfTree(t *testing.T) {
	d := &trees.Node{Val: 4}
	e := &trees.Node{Val: 2}
	f := &trees.Node{Val: 1}
	b := &trees.Node{Val: 11, Left: d, Right: e}
	c := &trees.Node{Val: 4, Right: f}
	a := &trees.Node{Val: 3, Left: b, Right: c}

	ans := trees.SumOfTreeDFS(a)
	if ans != 25 {
		t.Fail()
	}

}

func TestMinOfTree(t *testing.T) {
	d := &trees.Node{Val: 4}
	e := &trees.Node{Val: 2}
	f := &trees.Node{Val: 1}
	b := &trees.Node{Val: 11, Left: d, Right: e}
	c := &trees.Node{Val: 4, Right: f}
	a := &trees.Node{Val: 3, Left: b, Right: c}

	ans := trees.MinOfTreeBFS(a)
	if ans != 1 {
		t.Fail()
	}

	if ans != trees.MinOfTreeDFS(a) {
		t.Fail()
	}
}

func TestMaxPathOfTreeDFS(t *testing.T) {
	// d := &trees.Node{Val: 4}
	// e := &trees.Node{Val: -2}
	// f := &trees.Node{Val: 1}
	// b := &trees.Node{Val: 11, Left: d, Right: e}
	// c := &trees.Node{Val: 4, Right: f}
	// a := &trees.Node{Val: 3, Left: b, Right: c}

	// ans := trees.MaxPathOfTreeDFS(a)
	// if ans != 18 {
	// 	t.Fail()
	// }

	g := &trees.Node{Val: -1}
	h := &trees.Node{Val: -2}
	e := &trees.Node{Val: 0, Left: g}
	f := &trees.Node{Val: -13, Right: h}
	d := &trees.Node{Val: -3}
	c := &trees.Node{Val: -5, Right: f}
	b := &trees.Node{Val: -6, Left: d, Right: e}
	a := &trees.Node{Val: -1, Left: b, Right: c}

	ans := trees.MaxPathOfTreeDFS(a)
	if ans != -8 {
		t.Fail()
	}
}
