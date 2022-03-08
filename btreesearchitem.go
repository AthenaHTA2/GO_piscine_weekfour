package piscine

func BTreeSearchItem(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Data == data {
		return root
	}
	if data < root.Data {
		return BTreeSearchItem(root.Left, data)
	}
	return BTreeSearchItem(root.Right, data)
}
