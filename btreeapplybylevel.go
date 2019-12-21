package student

// BTreeApplyByLevel q
func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {

	height := BTreeGetHeight(root)

	for level := 1; level <= height; level++ {
		BTreeApplyToGivenLevel(root, level, f)
	}
}

// BTreeApplyToGivenLevel q
func BTreeApplyToGivenLevel(root *TreeNode, level int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	if level == 1 {
		f(root.Data)
	} else if level > 1 {
		BTreeApplyToGivenLevel(root.Left, level-1, f)
		BTreeApplyToGivenLevel(root.Right, level-1, f)
	}
}

// BTreeGetHeight q
func BTreeGetHeight(root *TreeNode) int {

	if root == nil {
		return 0
	}

	lHeight := 1 + BTreeGetHeight(root.Left)
	rHeight := 1 + BTreeGetHeight(root.Right)

	if lHeight < rHeight {
		return rHeight
	}

	return lHeight
}
