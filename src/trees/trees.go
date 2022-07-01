package trees

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func ExistsInTreeBFS(root *Node, target int) bool {
	if root == nil {
		return false
	}

	queue := []*Node{root}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Val == target {
			return true
		}

		if current.Left != nil {
			queue = append(queue, current.Left)
		}

		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}

	return false
}

func ExistsInTreeDFS(root *Node, target int) bool {
	if root == nil {
		return false
	}

	if root.Val == target {
		return true
	}

	return ExistsInTreeDFS(root.Left, target) || ExistsInTreeDFS(root.Right, target)
}

func SumOfTreeDFS(root *Node) int {
	if root == nil {
		return 0
	}

	return SumOfTreeDFS(root.Left) + SumOfTreeDFS(root.Right) + root.Val
}

func MinOfTreeBFS(root *Node) int {
	if root == nil {
		return -1
	}

	maxUint := ^uint(0)
	min := int(maxUint >> 1)

	queue := []*Node{root}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Val < min {
			min = current.Val
		}

		if current.Left != nil {
			queue = append(queue, current.Left)
		}

		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}

	return min
}

func MinOfTreeDFS(root *Node) int {
	maxUint := ^uint(0)
	min := int(maxUint >> 1)

	var dfs func(*Node)

	dfs = func(r *Node) {
		if r == nil {
			return
		}

		if r.Val < min {
			min = r.Val
		}

		dfs(r.Left)
		dfs(r.Right)
	}

	dfs(root)

	return min
}

func MaxPathOfTreeDFS(root *Node) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	if root.Left == nil {
		return MaxPathOfTreeDFS(root.Right) + root.Val
	}

	if root.Right == nil {
		return MaxPathOfTreeDFS(root.Left) + root.Val
	}

	left := MaxPathOfTreeDFS(root.Left)
	right := MaxPathOfTreeDFS(root.Right)

	if left >= right {
		return left + root.Val
	}
	return right + root.Val
}
