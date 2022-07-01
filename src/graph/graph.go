package graph

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func WidthOfTreeDFS(root *Node) int {
	var (
		mostLeft  int
		mostRight int
		dfs       func(*Node, int, int) (int, int)
	)

	dfs = func(n *Node, pl, pr int) (left int, right int) {
		if n == nil {
			return left, right
		}

		lLeft, lRight := dfs(n.Left, 1, 0)
		rLeft, rRight := dfs(n.Right, 0, 1)

		left = max(lLeft, rLeft) + pl - pr
		right = max(rRight, lRight) + pr - pl

		return left, right
	}

	mostLeft, mostRight = dfs(root, 0, 0)

	return mostLeft + mostRight
}

func WidhtofTreeNodeBFS(root *Node) int {
	ans := 1
	queue := []*Node{root}
	for len(queue) > 0 {
		next := make([]*Node, 0)
		for _, n := range queue {
			if n.Left != nil {
				next = append(next, n.Left)
			}
			if n.Right != nil {
				next = append(next, n.Right)
			}
		}
		if len(next) > 0 {
			ans = max(ans, next[len(next)-1].Val-next[0].Val+1)
		}
		queue = next
	}

	return ans
}

func HasPathDFS(nodes map[string][]string, src, dst string) bool {
	if src == dst {
		return true
	}

	for _, n := range nodes[src] {
		if HasPathDFS(nodes, n, dst) {
			return true
		}
	}

	return false
}

func HasPathBFS(nodes map[string][]string, src, dst string) bool {
	queue := []string{src}
	for len(queue) > 0 {
		current := queue[0]
		if current == dst {
			return true
		}

		queue = queue[1:]
		queue = append(queue, nodes[current]...)
	}

	return false
}

func CountConnectedComponents(nodes map[int][]int) int {
	var (
		count int
		dfs   func(int, map[int][]int) bool
	)

	visited := make(map[int]bool)

	dfs = func(current int, list map[int][]int) bool {
		if _, ok := visited[current]; ok {
			return false
		}

		visited[current] = true

		for _, neighbor := range list[current] {
			dfs(neighbor, list)
		}

		return true
	}

	for k := range nodes {
		if dfs(k, nodes) {
			count++
		}
	}

	return count
}

func LargestComponent(nodes map[int][]int) int {
	var (
		dfs func(int, map[int][]int) int
		ans int
	)

	visited := make(map[int]bool)

	dfs = func(current int, list map[int][]int) int {
		if _, ok := visited[current]; ok {
			return 0
		}

		visited[current] = true

		currentCount := 1
		for _, neighbor := range list[current] {
			currentCount += dfs(neighbor, list)
		}

		return currentCount
	}

	for k := range nodes {
		r := dfs(k, nodes)
		ans = max(ans, r)
	}

	return ans
}

type routeNode struct {
	node string
	size int
}

func ShortestRoute(edges [][2]string, start, end string) int {
	nodes := edgesToNodes(edges)
	queue := make([]routeNode, 1)
	queue[0] = routeNode{start, 0}
	visited := make(map[string]bool)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.node == end {
			return current.size
		}

		if _, ok := visited[current.node]; ok {
			continue
		}

		visited[current.node] = true

		for _, v := range nodes[current.node] {
			queue = append(queue, routeNode{v, current.size + 1})
		}
	}

	return -1
}

func edgesToNodes(edges [][2]string) map[string][]string {
	nodes := make(map[string][]string)

	for _, v := range edges {
		if _, ok := nodes[v[0]]; !ok {
			//nodes[v[0]] = append(nodes[v[0]], v[1])
			nodes[v[0]] = []string{}
		}

		if _, ok := nodes[v[1]]; !ok {
			nodes[v[1]] = []string{}
		}

		nodes[v[0]] = append(nodes[v[0]], v[1])
		nodes[v[1]] = append(nodes[v[1]], v[0])
	}

	return nodes
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func FindIslands(grid [][]string) int {
	var (
		count int
		dfs   func([][]string, [2]int) bool
	)

	visited := make(map[[2]int]bool)

	dfs = func(g [][]string, node [2]int) bool {
		if _, ok := visited[node]; ok {
			return false
		}
		visited[node] = true

		row, colum := node[0], node[1]

		rob := row < 0 || row >= len(g)
		cob := colum < 0 || colum >= len(g[0])
		if rob || cob {
			return false
		}

		if g[row][colum] == "w" {
			return false
		}

		dfs(g, [2]int{row - 1, colum})
		dfs(g, [2]int{row + 1, colum})
		dfs(g, [2]int{row, colum - 1})
		dfs(g, [2]int{row, colum + 1})

		return true
	}

	for i, rows := range grid {
		for j := range rows {
			if dfs(grid, [2]int{i, j}) {
				count++
			}
		}
	}

	return count
}

func MinimunIslandSize(grid [][]string) int {
	var (
		dfs func([][]string, [2]int) int
	)

	minimun := len(grid) * len(grid[0])
	visited := make(map[[2]int]bool)
	dfs = func(g [][]string, node [2]int) int {
		if _, ok := visited[node]; ok {
			return 0
		}
		visited[node] = true

		row, colum := node[0], node[1]
		rob := row < 0 || row >= len(g)
		cob := colum < 0 || colum >= len(g[0])
		if rob || cob {
			return 0
		}

		if g[row][colum] == "w" {
			return 0
		}

		count := 1
		count += dfs(g, [2]int{row - 1, colum}) + dfs(g, [2]int{row + 1, colum})
		count += dfs(g, [2]int{row, colum - 1}) + dfs(g, [2]int{row, colum + 1})

		return count
	}

	for i, rows := range grid {
		for j := range rows {
			size := dfs(grid, [2]int{i, j})
			if size > 0 {
				minimun = min(minimun, size)
			}
		}
	}

	return minimun
}

func BestBrige(grid [][]string) int {
	var traverse func([][]string, int, int, map[[2]int]bool) map[[2]int]bool
	traverse = func(graph [][]string, x int, y int, visited map[[2]int]bool) map[[2]int]bool {
		if isOutOfBounds(graph, x, y) {
			return visited
		}

		if grid[x][y] == "W" {
			return visited
		}

		key := [2]int{x, y}
		if visited[key] {
			return visited
		}
		visited[key] = true

		traverse(grid, x-1, y, visited)
		traverse(grid, x+1, y, visited)
		traverse(grid, x, y-1, visited)
		traverse(grid, x, y+1, visited)

		return visited
	}

	mainIsland := make(map[[2]int]bool)
out:
	for x, row := range grid {
		for y := range row {
			temp := traverse(grid, x, y, mainIsland)
			if len(temp) > 0 {
				break out
			}
		}
	}

	queue := make([][3]int, 0)
	for k := range mainIsland {
		queue = append(queue, [3]int{k[0], k[1], 0})
	}

	visited := make(map[[2]int]bool)
	for len(queue) > 0 {
		x, y, distance := queue[0][0], queue[0][1], queue[0][2]
		queue = queue[1:]

		if grid[x][y] == "L" && !mainIsland[[2]int{x, y}] {
			return distance - 1
		}

		deltas := [4][2]int{
			{-1, 0},
			{1, 0},
			{0, -1},
			{0, 1},
		}

		for _, d := range deltas {
			dX := x + d[0]
			dY := y + d[1]
			if !visited[[2]int{dX, dY}] && !isOutOfBounds(grid, dX, dY) {
				queue = append(queue, [3]int{dX, dY, distance + 1})
				visited[[2]int{dX, dY}] = true
			}
		}
	}

	return 0
}

func isOutOfBounds(grid [][]string, x int, y int) bool {
	xIssOff := x < 0 || x >= len(grid)
	yIssOff := y < 0 || y >= len(grid[0])
	return xIssOff || yIssOff
}

func HasCycle(graph map[string][]string) bool {
	var dpt func(map[string][]string, string, map[string]bool, map[string]bool) bool

	dpt = func(nodes map[string][]string, target string, visiting, visited map[string]bool) bool {
		if visited[target] {
			return false
		}

		if visiting[target] {
			return true
		}
		visiting[target] = true

		for _, v := range nodes[target] {
			if dpt(nodes, v, visiting, visited) {
				return true
			}
		}

		visited[target] = true

		return false
	}

	visiting, visited := make(map[string]bool), make(map[string]bool)
	for k := range graph {
		if dpt(graph, k, visiting, visited) {
			return true
		}
	}

	// dpt = func(nodes map[string][]string, target string, origin string, times int) bool {
	// 	if origin == target {
	// 		if times > 0 {
	// 			return true
	// 		}
	// 		times++
	// 	}

	// 	for _, v := range nodes[target] {
	// 		if dpt(nodes, v, origin, times) {
	// 			return true
	// 		}
	// 	}
	// 	return false
	// }

	// //visited := make(map[string]bool)
	// for k := range graph {
	// 	if dpt(graph, k, k, 0) {
	// 		return true
	// 	}
	// }

	return false
}
