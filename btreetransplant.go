package student

type TreeNode struct {
	Data                string
	Left, Right, Parent *TreeNode
}

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if node.Data == "" {
		return root
	} else if node.Parent == nil {
		root = rplc
	} else if node == node.Parent.Left {
		node.Parent.Left = rplc
	} else {
		node.Parent.Right = rplc
	}
	if rplc != nil {
		rplc.Parent = node.Parent
	}
	return root
}

func BTreeInsertData(root *TreeNode, data string) {
	if root == nil {
		root = &TreeNode{Data: data}
	}
	if data < root.Data {
		if root.Left != nil {
			BTreeInsertData(root.Left, data)
		} else {
			root.Left = &TreeNode{Data: data, Parent: root}
		}
	} else {
		if root.Right != nil {
			BTreeInsertData(root.Right, data)
		} else {
			root.Right = &TreeNode{Data: data, Parent: root}
		}
	}
}

func BTreeSearchItem(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{}
	}
	if data > root.Data {
		return BTreeSearchItem(root.Right, data)
	} else if data < root.Data {
		return BTreeSearchItem(root.Left, data)
	} else {
		return root
	}
}

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if root.Left != nil {
		BTreeApplyInorder(root.Left, f)
	}
	f(root.Data)
	if root.Right != nil {
		BTreeApplyInorder(root.Right, f)
	}
	return
}
