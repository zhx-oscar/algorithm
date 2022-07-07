package main

import "fmt"

type RBNode struct {
	value  uint32
	red    bool
	left   *RBNode
	right  *RBNode
	parent *RBNode
}

func Insert(root *RBNode, value uint32) *RBNode {
	if root == nil {
		root = &RBNode{
			value: value,
			red:   false,
		}
		return root
	}
	if value == root.value {
		return root
	}
	//寻找插入位置
	var parent *RBNode
	var current *RBNode = root
	var isLeftChild bool
	for current != nil {
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
	current = &RBNode{
		value:  value,
		parent: parent,
		red:    true,
	}
	if isLeftChild {
		parent.left = current
	} else {
		parent.right = current
	}
	//父节点颜色是红色，需要调整
	if parent.red {
		root = adjustInsert(root, parent)
	}
	return root
}

//current表示当前插入节点的父节点
//返回根节点
func adjustInsert(root, current *RBNode) *RBNode {
	if current == nil || root == nil {
		return nil
	}
	for current.red { //当前插入节点的父节点颜色如果是红色，继续执行逻辑
		//当前插入节点的叔节点
		uncle := getBrother(current)
		//因为current节点是红色，所以它的父节点一定存在
		parent := current.parent
		if uncle != nil && uncle.red { //当前插入节点的叔节点设为红色
			parent.red = true   //将插入节点的祖父节点设为红色
			current.red = false //将插入节点的父节点和叔节点设为黑色
			uncle.red = false
			if parent.parent != nil { //将插入节点的祖父节点作为新的插入节点，检查上层是否满足红黑树性质
				current = parent.parent
			} else { //插入节点的祖父节点是根节点，将其重新改为黑色并返回
				parent.red = false
				return parent
			}
		} else { //当前插入节点的叔节点为黑色或为空
			if current.left != nil { //插入节点是父节点的左孩子
				if parent.left == current { //插入节点的父节点是祖父节点的左孩子
					current.red = false
					parent.red = true
					if parent.parent != nil {
						gparent := parent.parent
						if gparent.left == parent {
							gparent.left = RRotate(parent)
							gparent.left.parent = gparent
						} else {
							gparent.right = RRotate(parent)
							gparent.right.parent = gparent
						}
					} else { //插入节点的祖父节点是根节点，右旋后返回
						return RRotate(parent)
					}
				} else { //插入节点的父节点是祖父节点的右孩子
					current.left.red = false
					parent.red = true
					if parent.parent != nil {
						gparent := parent.parent
						if gparent.left == parent {
							gparent.left = RLRotate(parent)
							gparent.left.parent = gparent
						} else {
							gparent.right = RLRotate(parent)
							gparent.right.parent = gparent
						}
					} else { //插入节点的祖父节点是根节点，右旋并左旋后返回
						return RLRotate(parent)
					}
				}
			} else { //插入节点是父节点的右孩子
				if parent.left == current { //插入节点的父节点是祖父节点的左孩子
					current.right.red = false
					parent.red = true
					if parent.parent != nil {
						gparent := parent.parent
						if gparent.left == parent {
							gparent.left = LRRotate(parent)
							gparent.left.parent = gparent
						} else {
							gparent.right = LRRotate(parent)
							gparent.right.parent = gparent
						}
					} else { //插入节点的祖父节点是根节点，左旋并右旋后返回
						return LRRotate(parent)
					}
				} else { //插入节点的父节点是祖父节点的右孩子
					current.red = false
					parent.red = true
					if parent.parent != nil {
						gparent := parent.parent
						if gparent.left == parent {
							gparent.left = LRotate(parent)
							gparent.left.parent = gparent
						} else {
							gparent.right = LRotate(parent)
							gparent.right.parent = gparent
						}
					} else { //插入节点的祖父节点是根节点，左旋后返回
						return LRotate(parent)
					}
				}
			}
			//退出循环
			break
		}
	}
	return root
}

func Delete(root *RBNode, value uint32) *RBNode {
	if root == nil { //删除根节点
		return nil
	}
	var isLeftChild bool
	var current *RBNode = root
	for current != nil {
		if current.value == value {
			break
		}
		if value < current.value {
			current = current.left
			isLeftChild = true
		} else {
			current = current.right
			isLeftChild = false
		}
	}
	for current != nil {
		if current.left == nil && current.right == nil { //删除节点无子节点
			if current.parent == nil { //删除节点是根节点
				return nil
			} else if current.red { //删除节点为红色
				if isLeftChild {
					current.parent.left = nil
				} else {
					current.parent.right = nil
				}
			} else { //删除节点为黑色
				root = deleteSingleBlack(root, current)
			}
			break
		} else if current.left == nil || current.right == nil { //删除节点有一个子节点，这种情况节点必为黑色，子节点必为红色
			if current.left != nil {
				current.value = current.left.value
				current.left = nil
			} else {
				current.value = current.right.value
				current.right = nil
			}
			break
		} else { //删除节点有两个子节点，处理它的后继节点
			successor := getSuccessor(current)
			current.value = successor.value
			current = successor
		}
	}
	return root
}

//将当前节点的兄弟改为黑色
//返回根节点
func changeBrotherToBlack(root, current *RBNode) *RBNode {
	if current == nil || current.parent == nil {
		return root
	}
	var parent *RBNode = current.parent
	var brother *RBNode
	var isLeftChild bool
	if parent.left == current {
		brother = parent.right
		isLeftChild = true
	} else if parent.right == current {
		brother = parent.left
		isLeftChild = false
	}
	if brother == nil {
		return root
	}
	brother.red = false
	parent.red = true
	gparent := parent.parent
	if gparent != nil {
		if gparent.left == parent {
			if isLeftChild {
				gparent.left = RRotate(parent)
			} else {
				gparent.left = LRotate(parent)
			}
			gparent.left.parent = gparent
		} else if gparent.right == parent {
			if isLeftChild {
				gparent.right = RRotate(parent)
			} else {
				gparent.right = LRotate(parent)
			}
			gparent.right.parent = gparent
		}
	} else {
		if isLeftChild {
			root = RRotate(parent)
		} else {
			root = LRotate(parent)
		}
	}
	return root
}

//删除无子节点的黑色节点
func deleteSingleBlack(root, current *RBNode) *RBNode {
	if root == nil || current == nil {
		return root
	}
	var justRotate bool
	for current != nil {
		brother := getBrother(current)
		if brother == nil {
			return root
		}
		if brother.red { //兄弟节点为红色，将兄弟节点颜色变为黑色
			root = changeBrotherToBlack(root, current)
		}
		parent := current.parent
		if parent == nil {
			return root
		}
		gparent := parent.parent
		if parent.left == current { //兄弟在右边
			if brother.right != nil && brother.right.red { //兄弟节点有右子节点且右子节点为红色
				brother.right.red = false
				brother.red = parent.red
				parent.red = false
				if !justRotate { //判断是否只旋转
					parent.left = nil
				}
				if gparent != nil {
					if gparent.left == parent {
						gparent.left = LRotate(parent)
						gparent.left.parent = gparent
					} else if gparent.right == parent {
						gparent.right = LRotate(parent)
						gparent.right.parent = gparent
					}
				} else {
					return LRotate(parent)
				}
			} else if brother.left != nil && brother.left.red { //兄弟节点有左子节点且左子节点为红色
				brother.left.red = parent.red
				parent.red = false
				if !justRotate { //判断是否只旋转
					parent.left = nil
				}
				if gparent != nil {
					if gparent.left == parent {
						gparent.left = RLRotate(parent)
						gparent.left.parent = gparent
					} else if gparent.right == parent {
						gparent.right = RLRotate(parent)
						gparent.right.parent = gparent
					}
				} else {
					return RLRotate(parent)
				}
			} else { //兄弟节点没有子节点或子节点都为黑色
				if parent.red { //父节点是红色
					brother.red = true
					parent.red = false
					if !justRotate {
						parent.left = nil
					}
				} else { //父节点是黑色，在上层进行平衡操作
					if !justRotate {
						parent.left = nil
					}
					brother.red = true
					current = parent
					justRotate = true
					continue
				}
			}
		} else if parent.right == current { //兄弟在左边
			if brother.left != nil && brother.left.red { //兄弟节点有左子节点且左子节点为红色
				brother.left.red = false
				brother.red = parent.red
				parent.red = false
				if !justRotate {
					parent.right = nil
				}
				if gparent != nil {
					if gparent.left == parent {
						gparent.left = RRotate(parent)
						gparent.left.parent = gparent
					} else if gparent.right == parent {
						gparent.right = RRotate(parent)
						gparent.right.parent = gparent
					}
				} else {
					return RRotate(parent)
				}
			} else if brother.right != nil && brother.right.red { //兄弟节点有右子节点且右子节点为红色
				brother.right.red = parent.red
				parent.red = false
				if !justRotate {
					parent.right = nil
				}
				if gparent != nil {
					if gparent.left == parent {
						gparent.left = LRRotate(parent)
						gparent.left.parent = gparent
					} else if gparent.right == parent {
						gparent.right = LRRotate(parent)
						gparent.right.parent = gparent
					}
				} else {
					return LRRotate(parent)
				}
			} else { //兄弟节点没有子节点或子节点都为黑色
				if parent.red { //父节点是红色
					brother.red = true
					parent.red = false
					if !justRotate {
						parent.right = nil
					}
				} else { //父节点是黑色，在上层进行平衡操作
					if !justRotate {
						parent.right = nil
					}
					brother.red = true
					current = parent
					justRotate = true
					continue
				}
			}
		}
		break
	}
	return root
}

func RRotate(node *RBNode) *RBNode {
	if node == nil {
		return nil
	}
	if node.left == nil {
		return node
	}
	left := node.left
	node.left = left.right
	left.right = node
	left.parent = nil
	node.parent = left
	if node.left != nil {
		node.left.parent = node
	}
	return left
}

func LRotate(node *RBNode) *RBNode {
	if node == nil {
		return nil
	}
	if node.right == nil {
		return node
	}
	right := node.right
	node.right = right.left
	right.left = node
	right.parent = nil
	node.parent = right
	if node.right != nil {
		node.right.parent = node
	}
	return right
}

func RLRotate(node *RBNode) *RBNode {
	if node == nil {
		return nil
	}
	if node.right == nil {
		return node
	}
	node.right = RRotate(node.right)
	node.right.parent = node
	return LRotate(node)
}

func LRRotate(node *RBNode) *RBNode {
	if node == nil {
		return nil
	}
	if node.left == nil {
		return node
	}
	node.left = LRotate(node.left)
	node.left.parent = node
	return RRotate(node)
}

func getSuccessor(node *RBNode) *RBNode {
	if node.right == nil {
		return nil
	}
	successor := node.right
	current := successor.left
	//寻找右子树中最小的子节点
	for current != nil {
		successor = current
		current = current.left
	}
	return successor
}

func getBrother(current *RBNode) *RBNode {
	if current == nil {
		return nil
	}
	if current.parent == nil {
		return nil
	}
	if current.parent.left == current {
		return current.parent.right
	}
	if current.parent.right == current {
		return current.parent.left
	}
	return nil
}

func PrintRBTree(root *RBNode) {
	if root == nil {
		fmt.Println("没有节点")
		return
	}
	fmt.Println("根节点: ", root.value, func() string {
		if root.red {
			return " 红色"
		} else {
			return " 黑色"
		}
	}())
	if root.left != nil {
		fmt.Println("左子树:")
		PrintRBTree(root.left)
	}
	if root.right != nil {
		fmt.Println("右子树:")
		PrintRBTree(root.right)
	}
	fmt.Println("返回上层")
}

func main() {
	fmt.Println("开始添加节点")
	root := Insert(nil, 12)
	root = Insert(root, 1)
	root = Insert(root, 9)
	root = Insert(root, 2)
	root = Insert(root, 0)
	root = Insert(root, 11)
	root = Insert(root, 7)
	root = Insert(root, 19)
	root = Insert(root, 4)
	root = Insert(root, 15)
	root = Insert(root, 18)
	root = Insert(root, 5)
	root = Insert(root, 14)
	root = Insert(root, 13)
	root = Insert(root, 10)
	root = Insert(root, 16)
	fmt.Println("添加节点结束")
	//PrintRBTree(root)
	fmt.Println("开始删除节点")
	root = Delete(root, 12)
	root = Delete(root, 1)
	root = Delete(root, 9)
	root = Delete(root, 2)
	root = Delete(root, 0)
	root = Delete(root, 11)
	root = Delete(root, 7)
	root = Delete(root, 19)
	root = Delete(root, 4)
	root = Delete(root, 15)
	root = Delete(root, 18)
	root = Delete(root, 5)
	root = Delete(root, 14)
	root = Delete(root, 13)
	root = Delete(root, 10)
	root = Delete(root, 16)
	fmt.Println("删除节点结束")
	PrintRBTree(root)
}
