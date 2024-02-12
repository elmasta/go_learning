package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	if root.Left != nil {
		toReturn := BTreeMin(root.Left)
		return toReturn
	}
	return root
}
