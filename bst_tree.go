package main

import "fmt"

type BSTNode struct {
	value uint32
	left  *BSTNode
	right *BSTNode
}

func Insert(root *BSTNode, value uint32) *BSTNode {
	if root == nil {
		root = &BSTNode{
			value: value,
		}
	}
	if value < root.value {
		root.left = Insert(root.left, value)
	} else if value > root.value {
		root.right = Insert(root.right, value)
	}
	return root
}

func Find(root *BSTNode, value uint32) *BSTNode {
	if root == nil {
		return nil
	}
	current := root
	for current != nil {
		if value == current.value {
			return current
		} else if value < current.value {
			current = current.left
		} else {
			current = current.right
		}
	}
	return nil
}

func PreOrderTraverse(root *BSTNode, values []uint32) []uint32 {
	if root == nil {
		return values
	}
	values = append(values, root.value)
	values = PreOrderTraverse(root.left, values)
	values = PreOrderTraverse(root.right, values)
	return values
}

func InOrderTraverse(root *BSTNode, values []uint32) []uint32 {
	if root == nil {
		return values
	}
	values = InOrderTraverse(root.left, values)
	values = append(values, root.value)
	values = InOrderTraverse(root.right, values)
	return values
}

func PostOrderTraverse(root *BSTNode, values []uint32) []uint32 {
	if root == nil {
		return values
	}
	values = PostOrderTraverse(root.left, values)
	values = PostOrderTraverse(root.right, values)
	values = append(values, root.value)
	return values
}

func Minimum(root *BSTNode) *BSTNode {
	if root == nil {
		return nil
	}
	current := root.left
	parent := root
	for current != nil {
		parent = current
		current = current.left
	}
	return parent
}

func Maximum(root *BSTNode) *BSTNode {
	if root == nil {
		return nil
	}
	current := root.right
	parent := root
	for current != nil {
		parent = current
		current = current.right
	}
	return parent
}

func Delete(root *BSTNode, value uint32) *BSTNode {
	if root == nil {
		return root
	}
	var isLeftChild bool
	var current *BSTNode = root
	var parent *BSTNode
	for current != nil {
		if current.value == value {
			break
		}
		if value < current.value {
			parent = current
			current = current.left
			isLeftChild = true
		} else {
			parent = current
			current = current.right
			isLeftChild = false
		}
	}
	//????????????????????????
	if current == nil {
		return root
	}
	//????????????
	if current.left == nil && current.right == nil {
		//????????????????????????????????????
		if parent == nil {
			return nil
		}
		if isLeftChild {
			parent.left = nil
		} else {
			parent.right = nil
		}
	}
	//??????????????????
	if current.left != nil && current.right == nil {
		//????????????????????????????????????
		if parent == nil {
			root = current.left
			return root
		}
		if isLeftChild {
			parent.left = current.left
		} else {
			parent.right = current.left
		}
	} else if current.right != nil && current.left == nil {
		//????????????????????????????????????
		if parent == nil {
			root = current.right
			return root
		}
		if isLeftChild {
			parent.left = current.right
		} else {
			parent.right = current.right
		}
	}
	//??????????????????
	if current.left != nil && current.right != nil {
		//??????????????????
		successor := getSuccessor(current)
		//????????????????????????????????????
		if parent == nil {
			root = successor
			return root
		}
		if isLeftChild {
			parent.left = successor
		} else {
			parent.right = successor
		}
	}
	return root
}

func getSuccessor(node *BSTNode) *BSTNode {
	if node.right == nil {
		return nil
	}
	successor := node.right
	successorParent := node
	current := successor.left
	//??????????????????????????????????????????
	for current != nil {
		successorParent = successor
		successor = current
		current = current.left
	}
	//?????????????????????????????????????????????????????????????????????????????????
	if successorParent != node {
		successorParent.left = successor.right
		successor.right = node.right
	}
	//?????????????????????????????????????????????
	successor.left = node.left
	return successor
}

func PrintBSTTree(root *BSTNode) {
	if root == nil {
		fmt.Println("????????????")
	}
	fmt.Println("?????????: ", root.value)
	if root.left != nil {
		fmt.Println("?????????:")
		PrintBSTTree(root.left)
	}
	if root.right != nil {
		fmt.Println("?????????:")
		PrintBSTTree(root.right)
	}
	fmt.Println("????????????")
}

func main() {
	var root *BSTNode
	root = Insert(root, 4)
	root = Insert(root, 2)
	root = Insert(root, 1)
	root = Insert(root, 3)
	root = Insert(root, 6)
	root = Insert(root, 5)
	root = Insert(root, 8)
	root = Insert(root, 7)
	root = Delete(root, 1)
	root = Delete(root, 8)
	root = Delete(root, 6)
	root = Delete(root, 4)
	PrintBSTTree(root)
	node := Find(root, 5)
	if node != nil {
		fmt.Printf("node:%d found\n", node.value)
	}
	values := make([]uint32, 0, 7)
	values = PreOrderTraverse(root, values)
	fmt.Println("PreOrder: ", values)
	values = values[:0]
	values = InOrderTraverse(root, values)
	fmt.Println("InOrder: ", values)
	values = values[:0]
	values = PostOrderTraverse(root, values)
	fmt.Println("PostOrder: ", values)
	min := Minimum(root)
	fmt.Println("Min: ", min.value)
	max := Maximum(root)
	fmt.Println("Max: ", max.value)
}
