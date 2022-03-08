package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if root.Data > data {
		if root.Left == nil {
			root.Left = &TreeNode{Left: nil, Right: nil, Parent: root, Data: data}
		} else {
			BTreeInsertData(root.Left, data)
		}
	} else if root.Data < data {
		if root.Right == nil {
			root.Right = &TreeNode{Left: nil, Right: nil, Parent: root, Data: data}
		} else {
			BTreeInsertData(root.Right, data)
		}
	}
	return root
}
