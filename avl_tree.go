package main

import "fmt"

type AVLNode struct {
	value  uint32
	height uint32
	left   *AVLNode
	right  *AVLNode
}

func Insert(root *AVLNode, value uint32) *AVLNode {
	if root == nil {
		root = &AVLNode{
			value:  value,
			height: 1,
		}
		return root
	}
	if value == root.value {
		return root
	}
	if value < root.value { //新值插入左子树的情况
		root.left = Insert(root.left, value)
		if getHeight(root.left)-getHeight(root.right) > 1 { //当前树深度已失衡
			if value < root.left.value { //新插入节点在左子树根节点的左子树中
				root = LLRotate(root)
			} else { //新插入节点在左子树根节点的右子树中
				root = LRRotate(root)
			}
		}
	} else { //新值插入右子树的情况
		root.right = Insert(root.right, value)
		if getHeight(root.right)-getHeight(root.left) > 1 { //当前树深度已失衡
			if value < root.right.value { //新插入节点在右子树根节点的左子树中
				root = RLRotate(root)
			} else { //新插入节点在右子树根节点的右子树中
				root = RRRotate(root)
			}

		}
	}
	//更新根节点深度
	root.height = max(getHeight(root.left), getHeight(root.right)) + 1
	return root
}

func LLRotate(root *AVLNode) *AVLNode {
	if root == nil {
		return nil
	}
	lSubTree := root.left
	if lSubTree == nil {
		return root
	}
	root.left = lSubTree.right
	//更新之前根节点的深度
	root.height = max(getHeight(root.left), getHeight(root.right)) + 1
	lSubTree.right = root
	//更新新的根节点的深度
	lSubTree.height = max(getHeight(lSubTree.left), getHeight(lSubTree.right)) + 1
	return lSubTree
}

func RRRotate(root *AVLNode) *AVLNode {
	if root == nil {
		return nil
	}
	rSubTree := root.right
	if rSubTree == nil {
		return root
	}
	root.right = rSubTree.left
	//更新之前根节点的深度
	root.height = max(getHeight(root.left), getHeight(root.right)) + 1
	rSubTree.left = root
	//更新新的根节点的深度
	rSubTree.height = max(getHeight(rSubTree.left), getHeight(rSubTree.right)) + 1
	return rSubTree
}

func LRRotate(root *AVLNode) *AVLNode {
	if root == nil {
		return nil
	}
	//对根节点的左子树进行左旋
	root.left = RRRotate(root.left)
	//对整棵树进行右旋
	root = LLRotate(root)
	return root
}

func RLRotate(root *AVLNode) *AVLNode {
	if root == nil {
		return nil
	}
	//对根节点的右子树进行右旋
	root.right = LLRotate(root.right)
	//对整棵树进行左旋
	root = RRRotate(root)
	return root
}

func getHeight(node *AVLNode) uint32 {
	if node == nil {
		return 0
	}
	return node.height
}

func max(a, b uint32) uint32 {
	if a >= b {
		return a
	}
	return b
}

func PrintAVLTree(root *AVLNode) {
	if root == nil {
		fmt.Println("没有节点")
	}
	fmt.Println("根节点: ", root.value)
	if root.left != nil {
		fmt.Println("左子树:")
		PrintAVLTree(root.left)
	}
	if root.right != nil {
		fmt.Println("右子树:")
		PrintAVLTree(root.right)
	}
	fmt.Println("返回上层")
}

func main() {
	var root *AVLNode
	root = Insert(root, 40)
	root = Insert(root, 20)
	root = Insert(root, 50)
	root = Insert(root, 10)
	root = Insert(root, 30)
	root = Insert(root, 5)
	PrintAVLTree(root)
}
