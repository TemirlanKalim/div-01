package student

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	level := BTreeLevel(root)
	for i := 1; i <= level; i++ {
		BTreeApplyByLevel2(root, f, i)
	}
}

func BTreeApplyByLevel2(root *TreeNode, f func(...interface{}) (int, error), level int) {
	if root == nil || level < 1 {
		return
	}
	if level == 1 {
		f(root.Data)
		return
	}
	if root.Left != nil {
		BTreeApplyByLevel2(root.Left, f, level-1)
	}
	if root.Right != nil {
		BTreeApplyByLevel2(root.Right, f, level-1)
	}
}

func BTreeLevel(root *TreeNode) int {
	if root == nil {
		return 0
	}
	countA := 0
	countB := 0
	if root.Left != nil {
		countA = BTreeLevel(root.Left)
	}
	if root.Right != nil {
		countB = BTreeLevel(root.Right)
	}
	if countA > countB {
		countB = countA
	}
	return countB + 1
}
