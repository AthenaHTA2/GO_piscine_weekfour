package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil || root.Data == node.Data {
		return rplc
	}
	if root.Data > node.Data {
		swap := BTreeSearchItem(root.Left, node.Data)
		swap.Data = rplc.Data
		swap.Left = rplc.Left
		swap.Right = rplc.Right
	} else {
		swap := BTreeSearchItem(root.Right, node.Data)
		swap.Data = rplc.Data
		swap.Left = rplc.Left
		swap.Right = rplc.Right
	}
	return root
}
