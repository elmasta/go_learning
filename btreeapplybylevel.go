package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		// Create a queue and insert root
		queue := []*TreeNode{}
		queue = append(queue, root)
		// Process as long as queue is not empty
		for len(queue) > 0 {
			// Get the current size or length of the queue.
			// This indicates the total number of nodes that are part of current level
			sz := len(queue)
			for i := 0; i < sz; i++ {
				// Remove a node
				node := queue[0]
				queue = queue[1:]
				// Do the func
				f(node.Data)
				// Insert children of the node into the queue
				if node.Left != nil {
					queue = append(queue, node.Left)
				}
				if node.Right != nil {
					queue = append(queue, node.Right)
				}
			}
		}
	}
}
