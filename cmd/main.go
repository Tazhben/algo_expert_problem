package main

import (
	"Tazhben/algo_expert_problem/topic/tree"
	"fmt"
)

func main() {
	t8 := tree.BinaryTree{Value: 8}
	t9 := tree.BinaryTree{Value: 9}
	t6 := tree.BinaryTree{Value: 6}
	t7 := tree.BinaryTree{Value: 7}
	t5 := tree.BinaryTree{Value: 5}
	t4 := tree.BinaryTree{Value: 4, Left: &t8, Right: &t9}
	t3 := tree.BinaryTree{Value: 3, Left: &t6, Right: &t7}
	t2 := tree.BinaryTree{Value: 2, Left: &t4, Right: &t5}
	t1 := tree.BinaryTree{Value: 1, Left: &t2, Right: &t3}
	fmt.Print(tree.NodeDepths2(&t1))
}
