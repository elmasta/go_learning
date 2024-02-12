package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	clone := root
	for clone != nil {
		if data >= clone.Data {
			if clone.Right == nil {
				clone.Right = &TreeNode{Data: data, Parent: clone}
				break
			}
			clone = clone.Right
		} else {
			if clone.Left == nil {
				clone.Left = &TreeNode{Data: data, Parent: clone}
				break
			}
			clone = clone.Left
		}
	}
	return root
}
