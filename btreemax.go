package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	if root.Right != nil {
		toReturn := BTreeMax(root.Right)
		return toReturn
	}
	return root
}
