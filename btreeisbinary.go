package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	isValid, _, _ := isValidBSTUtil(root)
	return isValid
}

func isValidBSTUtil(node *TreeNode) (bool, string, string) {
	if node.Left == nil && node.Right == nil {
		return true, node.Data, node.Data
	}

	min := node.Data
	max := node.Data

	isValidLeft := true
	var leftMin, leftMax string
	if node.Left != nil {
		isValidLeft, leftMin, leftMax = isValidBSTUtil(node.Left)

		if !isValidLeft {
			return false, "", ""
		}
		if node.Data <= leftMax {
			return false, "", ""
		}

		min = leftMin
	}

	isValidRight := true
	var rightMin, rightMax string

	if node.Right != nil {
		isValidRight, rightMin, rightMax = isValidBSTUtil(node.Right)

		if !isValidRight {
			return false, "", ""
		}

		if node.Data >= rightMin {
			return false, "", ""
		}
		max = rightMax
	}

	return true, min, max
}
