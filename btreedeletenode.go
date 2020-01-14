package student

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return nil
	} else if node.Data == "" {
		return root
	}
	if node.Right != nil {
		if node.Left == nil {
			if node.Parent == nil {
				node.Right.Parent = nil
				return node.Right
			} else {
				if node == node.Parent.Right {
					node.Right.Parent = node.Parent
					node.Parent.Right = node.Right
				} else {
					node.Right.Parent = node.Parent
					node.Parent.Left = node.Right
				}
			}
			return root
		} else {
			rNode := BTreeMin(node.Right)
			if node.Right == rNode {
				if rNode.Right != nil {
					rNode.Left = node.Left
					if rNode.Left != nil {
						rNode.Left.Parent = rNode
					}
				}
			} else {
				if rNode.Right != nil {
					rNode.Parent.Left = rNode.Right
					rNode.Right.Parent = rNode.Parent
				} else {
					rNode.Parent.Left = nil
				}
				node.Right.Parent = rNode
				rNode.Right = node.Right
			}
			rNode.Left = node.Left
			if rNode.Left != nil {
				rNode.Left.Parent = rNode
			}
			///////
			if node.Parent == nil {
				rNode.Parent = nil
				return rNode
			} else {
				if node == node.Parent.Right {
					rNode.Parent = node.Parent
					node.Parent.Right = rNode
				} else {
					rNode.Parent = node.Parent
					node.Parent.Left = rNode
				}
			}
			///////
			return root
		}

	} else if node.Left != nil {
		if node.Parent == nil {
			node.Left.Parent = nil
			return node.Left
		}
		node.Left.Parent = node.Parent
		if node == node.Parent.Right {
			node.Parent.Right = node.Left
		} else {
			node.Parent.Left = node.Left
		}
	} else {
		if node.Parent == nil {
			return nil
		} else {
			if node == node.Parent.Right {
				node.Parent.Right = nil
			} else {
				node.Parent.Left = nil
			}
		}
	}
	return root
}

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	for root.Left != nil {
		root = root.Left
	}
	return root
}
