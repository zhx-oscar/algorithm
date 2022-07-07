package main

import "fmt"

type TreeNode struct {
	value uint32
	left  *TreeNode
	right *TreeNode
}

func LevelOrderTraverse(root *TreeNode) []uint32 {
	if root == nil {
		return nil
	}
	var valueQueue []uint32
	var curNodeQueue []*TreeNode = []*TreeNode{root}
	var nextNodeQueue []*TreeNode
	for len(curNodeQueue) > 0 {
		for i := 0; i < len(curNodeQueue); i++ {
			node := curNodeQueue[i]
			valueQueue = append(valueQueue, node.value)
			if node.left != nil {
				nextNodeQueue = append(nextNodeQueue, node.left)
			}
			if node.right != nil {
				nextNodeQueue = append(nextNodeQueue, node.right)
			}
		}
		curNodeQueue = nextNodeQueue
		nextNodeQueue = nil
	}
	return valueQueue
}

func queueReverse(queue []uint32) {
	for i, j := 0, len(queue)-1; i < j; i, j = i+1, j-1 {
		queue[i], queue[j] = queue[j], queue[i]
	}
}

func LevelOrderZigZagTraverse(root *TreeNode) []uint32 {
	if root == nil {
		return nil
	}
	var valueQueue, curValueQueue []uint32
	var curNodeQueue []*TreeNode = []*TreeNode{root}
	var nextNodeQueue []*TreeNode
	var reverse bool
	for len(curNodeQueue) > 0 {
		for i := 0; i < len(curNodeQueue); i++ {
			node := curNodeQueue[i]
			curValueQueue = append(curValueQueue, node.value)
			if node.left != nil {
				nextNodeQueue = append(nextNodeQueue, node.left)
			}
			if node.right != nil {
				nextNodeQueue = append(nextNodeQueue, node.right)
			}
		}
		curNodeQueue = nextNodeQueue
		nextNodeQueue = nil
		if reverse {
			queueReverse(curValueQueue)
		}
		reverse = !reverse
		valueQueue = append(valueQueue, curValueQueue...)
		curValueQueue = nil
	}
	return valueQueue
}

func MiddleOrderTraverse(root *TreeNode, values []uint32) []uint32 {
	if root == nil {
		return values
	}
	values = MiddleOrderTraverse(root.left, values)
	values = append(values, root.value)
	values = MiddleOrderTraverse(root.right, values)
	return values
}

func main() {
	root := &TreeNode{
		value: 1,
	}
	root.left = &TreeNode{
		value: 2,
	}
	root.right = &TreeNode{
		value: 3,
	}
	root.left.left = &TreeNode{
		value: 4,
	}
	root.left.right = &TreeNode{
		value: 5,
	}
	root.right.left = &TreeNode{
		value: 6,
	}
	root.right.right = &TreeNode{
		value: 7,
	}
	root.right.left.left = &TreeNode{
		value: 8,
	}
	root.right.left.right = &TreeNode{
		value: 9,
	}
	root.right.right.left = &TreeNode{
		value: 10,
	}
	root.right.right.right = &TreeNode{
		value: 11,
	}
	values := LevelOrderTraverse(root)
	fmt.Println("Level order traverse: ", values)
	values = LevelOrderZigZagTraverse(root)
	fmt.Println("Level order zigzag traverse: ", values)
	values = values[:0]
	values = MiddleOrderTraverse(root, values)
	fmt.Println("Middle order traverse: ", values)
}
