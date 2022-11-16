package main

import "fmt"

func tral(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.left == nil && root.right == nil {
		return true
	}
	if root.left == nil || root.right == nil || root.left.val != root.right.val {
		return false
	}
	return tral(root.left) && tral(root.right)

}
func main() {
	//var node =
	var nodePointer *TreeNode = nil
	fmt.Println(tral(nodePointer))

}
