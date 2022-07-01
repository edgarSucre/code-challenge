package graph_test

import (
	"testing"

	"github.com/edgarsucre/challenge/src/graph"
)

func TestWidthOfTree(t *testing.T) {
	// r2 := &graph.Node{Val: 4}
	// r1 := &graph.Node{Val: 3, Right: r2}
	// l := &graph.Node{Val: 2, Right: r1}
	// n := &graph.Node{Val: 1, Left: l}
	//+++++++++++++++++++++++++++++++++++++++++
	// l := &graph.Node{Val: 2}
	// r := &graph.Node{Val: 3}
	// n := &graph.Node{Val: 1, Left: l, Right: r}
	//+++++++++++++++++++++++++++++++++++++++++
	// l1 := &graph.Node{Val: 4}
	// l := &graph.Node{Val: 2, Left: l1}
	// r := &graph.Node{Val: 3}
	// n := &graph.Node{Val: 1, Left: l, Right: r}
	//++++++++++++++++++++++++++++++++++++++++++
	// r := &graph.Node{Val: 3}
	// l := &graph.Node{Val: 2, Right: r}
	// n := &graph.Node{Val: 1, Left: l}
	//+++++++++++++++++++++++++++++++++++++++++
	r := &graph.Node{Val: 3}
	l1 := &graph.Node{Val: 4, Right: r}
	l := &graph.Node{Val: 2, Left: l1}
	n := &graph.Node{Val: 1, Left: l}

	ans := graph.WidthOfTreeDFS(n)
	if ans != 2 {
		t.Fail()
	}

}

func TestHasPath(t *testing.T) {
	nodes := make(map[string][]string)
	nodes["f"] = []string{"g", "i"}
	nodes["g"] = []string{"h"}
	nodes["h"] = []string{}
	nodes["i"] = []string{"g", "k"}
	nodes["j"] = []string{"i"}
	nodes["k"] = []string{}

	ans := graph.HasPathDFS(nodes, "f", "k")
	if !ans {
		t.Fail()
	}

	ans2 := graph.HasPathBFS(nodes, "f", "h")
	if !ans2 {
		t.Fail()
	}
}

func TestCountConnectedComponents(t *testing.T) {
	nodes := make(map[int][]int)
	nodes[0] = []int{8, 1, 5}
	nodes[1] = []int{0}
	nodes[5] = []int{0, 8}
	nodes[8] = []int{0, 5}
	nodes[2] = []int{3, 4}
	nodes[3] = []int{2, 4}
	nodes[4] = []int{3, 2}

	ans := graph.CountConnectedComponents(nodes)
	if ans != 2 {
		t.Fail()
	}
}

func TestLargestComponent(t *testing.T) {
	nodes := make(map[int][]int)
	nodes[0] = []int{8, 1, 5}
	nodes[1] = []int{0}
	nodes[5] = []int{0, 8}
	nodes[8] = []int{0, 5}
	nodes[2] = []int{3, 4}
	nodes[3] = []int{2, 4}
	nodes[4] = []int{3, 2}

	ans := graph.LargestComponent(nodes)
	if ans != 4 {
		t.Fail()
	}
}

func TestShortestRoute(t *testing.T) {
	edges := [][2]string{
		{
			"w", "x",
		},
		{
			"x", "y",
		},
		{
			"z", "y",
		},
		{
			"z", "v",
		},
		{
			"w", "v",
		},
	}

	ans := graph.ShortestRoute(edges, "w", "z")
	if ans != 2 {
		t.Fail()
	}
}

func TestFindIslands(t *testing.T) {
	nodes := [][]string{
		{"w", "l", "w", "w", "w"},
		{"w", "l", "w", "w", "w"},
		{"w", "w", "w", "l", "w"},
		{"w", "w", "l", "l", "w"},
		{"l", "w", "w", "l", "l"},
		{"l", "l", "w", "w", "w"},
	}

	ans := graph.FindIslands(nodes)
	if ans != 3 {
		t.Fail()
	}
}

func TestMinimunIslandSize(t *testing.T) {
	nodes := [][]string{
		{"w", "l", "w", "w", "w"},
		{"w", "l", "w", "w", "w"},
		{"w", "w", "w", "l", "w"},
		{"w", "w", "l", "l", "w"},
		{"l", "w", "w", "l", "l"},
		{"l", "l", "w", "w", "w"},
	}

	ans := graph.MinimunIslandSize(nodes)
	if ans != 2 {
		t.Fail()
	}
}

func TestBestBridge(t *testing.T) {
	nodes := [][]string{
		{"W", "W", "W", "L", "L"},
		{"L", "L", "W", "W", "L"},
		{"L", "L", "L", "W", "L"},
		{"W", "L", "W", "W", "W"},
		{"W", "W", "W", "W", "W"},
		{"W", "W", "W", "W", "W"},
	}

	ans := graph.BestBrige(nodes)
	if ans != 1 {
		t.Fail()
	}

	nodes = [][]string{
		{"W", "W", "W", "W", "W"},
		{"W", "W", "W", "W", "W"},
		{"L", "L", "W", "W", "L"},
		{"W", "L", "W", "W", "L"},
		{"W", "W", "W", "L", "L"},
		{"W", "W", "W", "W", "W"},
	}

	ans = graph.BestBrige(nodes)
	if ans != 2 {
		t.Fail()
	}

	nodes = [][]string{
		{"W", "W", "W", "W", "W"},
		{"W", "W", "W", "L", "W"},
		{"L", "W", "W", "W", "W"},
	}

	ans = graph.BestBrige(nodes)
	if ans != 3 {
		t.Fail()
	}

	nodes = [][]string{
		{"W", "W", "W", "W", "W", "W", "W", "W"},
		{"W", "W", "W", "W", "W", "W", "W", "W"},
		{"W", "W", "W", "W", "W", "W", "W", "W"},
		{"W", "W", "W", "W", "W", "L", "W", "W"},
		{"W", "W", "W", "W", "L", "L", "W", "W"},
		{"W", "W", "W", "W", "L", "L", "L", "W"},
		{"W", "W", "W", "W", "W", "L", "L", "L"},
		{"L", "W", "W", "W", "W", "L", "L", "L"},
		{"L", "L", "L", "W", "W", "W", "W", "W"},
		{"W", "W", "W", "W", "W", "W", "W", "W"},
	}

	ans = graph.BestBrige(nodes)
	if ans != 3 {
		t.Fail()
	}

	nodes = [][]string{
		{"L", "L", "L", "L", "L", "L", "L", "L"},
		{"L", "W", "W", "W", "W", "W", "W", "L"},
		{"L", "W", "W", "W", "W", "W", "W", "L"},
		{"L", "W", "W", "W", "W", "W", "W", "L"},
		{"L", "W", "W", "W", "W", "W", "W", "L"},
		{"L", "W", "W", "W", "W", "W", "W", "L"},
		{"L", "W", "W", "L", "W", "W", "W", "L"},
		{"L", "W", "W", "W", "W", "W", "W", "L"},
		{"L", "W", "W", "W", "W", "W", "W", "L"},
		{"L", "W", "W", "W", "W", "W", "W", "L"},
		{"L", "W", "W", "W", "W", "W", "W", "L"},
		{"L", "L", "L", "L", "L", "L", "L", "L"},
	}

	ans = graph.BestBrige(nodes)
	if ans != 2 {
		t.Fail()
	}

	nodes = [][]string{
		{"W", "L", "W", "W", "W", "W", "W", "W"},
		{"W", "L", "W", "W", "W", "W", "W", "W"},
		{"W", "W", "W", "W", "W", "W", "W", "W"},
		{"W", "W", "W", "W", "W", "W", "W", "W"},
		{"W", "W", "W", "W", "W", "W", "W", "W"},
		{"W", "W", "W", "W", "W", "W", "L", "W"},
		{"W", "W", "W", "W", "W", "W", "L", "L"},
		{"W", "W", "W", "W", "W", "W", "W", "L"},
	}

	ans = graph.BestBrige(nodes)
	if ans != 8 {
		t.Fail()
	}
}

func TestHasCycle(t *testing.T) {
	nodes := map[string][]string{
		"a": {"b"},
		"b": {"c"},
		"c": {"a"},
	}

	if !graph.HasCycle(nodes) {
		t.Fail()
	}

	nodes = map[string][]string{
		"a": {"b", "c"},
		"b": {"c"},
		"c": {"d"},
		"d": {},
	}

	if graph.HasCycle(nodes) {
		t.Fail()
	}
}
