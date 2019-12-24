package student

// TreeNode q
type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

// BTreeTransplant q
func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if node == root {
		return rplc
	}

	if node.Parent.Left == node {
		node.Parent.Left = rplc
	} else if node.Parent.Right == node {
		node.Parent.Right = rplc
	} else {
		Print("Node not found in given root\n")
	}

	return root
}

// BTreeInsertData q
func BTreeInsertData(root *TreeNode, data string) {
	if data < root.Data {
		if root.Left == nil {
			root.Left = &TreeNode{Data: data, Parent: root}
		} else {
			BTreeInsertData(root.Left, data)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{Data: data, Parent: root}
		} else {
			BTreeInsertData(root.Right, data)
		}
	}
}

// BTreeSearchItem q
func BTreeSearchItem(root *TreeNode, data string) *TreeNode {

	if root != nil {
		if data < root.Data {
			return BTreeSearchItem(root.Left, data)
		}

		if data > root.Data {
			return BTreeSearchItem(root.Right, data)
		}
	}

	return root
}

// BTreeApplyInorder q
func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		BTreeApplyInorder(root.Left, f)
		f(root.Data)
		BTreeApplyInorder(root.Right, f)
	}
}
