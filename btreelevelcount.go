package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	lHeight := BTreeLevelCount(root.Left)
	rHeight := BTreeLevelCount(root.Right)

	if lHeight >= rHeight {
		return lHeight + 1
	} else {
		return rHeight + 1
	}
}
