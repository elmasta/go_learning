package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	// enter the realm of black magic and child abuse
	if node.Right != nil && node.Left != nil {
		// two childs (double the childs, double the troubles)
		// find the tiniest on the right
		tiniestNode := BTreeMin(node.Right)
		if tiniestNode.Right != nil {
			// if tiniest has right child (just stop reproducing already!)
			temp := tiniestNode.Right
			tiniestNode.Parent.Left = temp
			temp.Parent = tiniestNode.Parent
		} else {
			tiniestNode.Parent.Left = nil
		}
		if node.Parent != nil {
			// check if parent exist
			if node.Parent.Data > node.Data {
				// if parent > current
				node.Parent.Left = tiniestNode
			} else {
				// if parent < current
				node.Parent.Right = tiniestNode
			}
		}
		tiniestNode.Parent = node.Parent
		tiniestNode.Right = node.Right
		tiniestNode.Left = node.Left
		node.Left.Parent = tiniestNode
		node.Right.Parent = tiniestNode
		return tiniestNode
	} else if node.Right != nil {
		// one child right
		if node.Parent != nil {
			// check if parent exist
			if node.Parent.Data > node.Data {
				// if parent > current
				node.Parent.Left = node.Right
			} else {
				// if parent < current
				node.Parent.Right = node.Right
			}
			node.Right.Parent = node.Parent
			node.Parent = nil
			node.Right = nil
		} else {
			// if top of tree
			temp := node.Right
			node.Right = nil
			return temp
		}
	} else if node.Left != nil {
		// one child left
		if node.Parent != nil {
			// check if parent exist
			if node.Parent.Data > node.Data {
				// if parent > current
				node.Parent.Left = node.Left
			} else {
				// if parent < current
				node.Parent.Right = node.Left
			}
			node.Left.Parent = node.Parent
			node.Parent = nil
			node.Left = nil
		} else {
			// if top of tree
			temp := node.Left
			node.Left = nil
			return temp
		}
	} else if node.Parent != nil {
		// no childs one parent
		if node.Parent.Data > node.Data {
			// if parent > current
			node.Parent.Left = nil
		} else {
			// if parent < current
			node.Parent.Right = nil
		}
		node.Parent = nil
	} else {
		// only one element
		return nil
	}
	return root
}
