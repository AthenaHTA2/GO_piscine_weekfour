package piscine

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	} else {
		f(root.Data)
		BTreeApplyPreorder(root.Left, f)
		BTreeApplyPreorder(root.Right, f)
	}
}
