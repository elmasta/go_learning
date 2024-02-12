package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if node.Parent != nil {
		if node.Data > node.Parent.Data {
			node.Parent.Right = rplc
		} else {
			node.Parent.Left = rplc
		}
		rplc.Parent = node.Parent
	}
	return root
}
