package tree

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func BranchSums(root *BinaryTree) []int {
	return Rec(root, make([]int, 0), 0)
}

func Rec(root *BinaryTree, all []int, s int) []int {
	if root.Left == nil && root.Right == nil {
		all = append(all, root.Value)
		return all
	}
	s += root.Value

	Rec(root.Right, all, s)
	Rec(root.Left, all, s)
	return all
}