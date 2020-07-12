package Week_03

// 前序遍历 root (index=0) 寻找 中序遍历 root (index=k)

func buildTree(preorder []int, inorder []int) *TreeNode {
	for k := range inorder {
		if inorder[k] == preorder[0] {//中序遍历 root (index=k)
			return &TreeNode{
				Val: preorder[0],
				//Val: inorder[root],
				Left:  buildTree(preorder[1:k+1], inorder[0:k]),
				Right: buildTree(preorder[k+1:], inorder[k+1:]),
			}
		}
	}
	return nil
}

