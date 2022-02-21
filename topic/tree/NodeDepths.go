package tree

func NodeDepths(root *BinaryTree) int {
	return F(root, 0, -1)
}

func F(root *BinaryTree, sum int, k int) int {
	if root == nil {
		return 0
	}
	k++
	if root.Left == nil && root.Right == nil {
		sum = sum + k
		return sum
	}
	if root.Right != nil {
		sum = F(root.Right, sum, k)
	}

	if root.Left != nil {
		sum = F(root.Left, sum, k)
	}
	sum = sum + k
	return sum
}
