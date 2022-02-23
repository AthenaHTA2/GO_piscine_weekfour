package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
	} else if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else {
		if root.Left == nil {
			return root.Right
		} else {
			temp := root.Left
			for temp.Right != nil {
				temp = temp.Right
			}
			root.Data = temp.Data
			root.Left = BTreeDeleteNode(root.Left, node)
		}
	}
	return root
}
