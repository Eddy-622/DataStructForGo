package data_struct

import (
	"math"
)

type AVLTree struct {
	Root *treeNode
	size int
}

func NewAVLTree() *AVLTree {
	return &AVLTree{}
}

// Insert 插入数据
func (a *AVLTree) Insert(v int) {
	newNode := &treeNode{Value: v}
	a.Root = a.insert(newNode, a.Root)
}

// insert 递归方法，更新节点数据同时计算平衡，并维持平衡
func (a *AVLTree) insert(newNode *treeNode, root *treeNode) *treeNode {
	if root == nil {
		return newNode
	}
	if newNode.Value < root.Value {
		root.LeftChild = a.insert(newNode, root.LeftChild)
	} else {
		root.RightChild = a.insert(newNode, root.RightChild)
	}

	// 更新root节点高度
	a.setHeight(root)

	// 检查是否平衡 Balance = left - right
	// < -1 右重
	// > 1  坐重
	//balance := a.balanceFactor(root)
	return a.balance(root)
}

func (a *AVLTree) setHeight(node *treeNode) {
	node.height = int(math.Max(float64(a.height(node.LeftChild)), float64(a.height(node.RightChild)))) + 1
}

func (a *AVLTree) balance(root *treeNode) *treeNode {
	// 旋转方式
	//     情况1                        情况2                         情况3               情况4 ...
	//  10                           10                                10
	//     20      -1                    30  1                      5      -1
	//         30                    20                                 8
	// 左旋 leftRotate(10)          rightRotate(30)                leftRotate(5)
	//                             leftRotate(10)                 rightRotate(10)

	if a.isLeftHeavy(root) {
		if a.height(root.LeftChild) < 0 { // 情况4
			root.LeftChild = a.rotateLeft(root.LeftChild)
		}
		return a.rotateRight(root)

	} else if a.isRightHeavy(root) {
		if a.height(root.RightChild) > 0 { // 情况2
			root.LeftChild = a.rotateLeft(root.RightChild)
		}
		return a.rotateLeft(root)
	}
	return root
}

func (a *AVLTree) rotateLeft(root *treeNode) *treeNode {
	newRoot := root.RightChild
	root.RightChild = newRoot.LeftChild
	newRoot.LeftChild = root

	a.setHeight(root)
	a.setHeight(newRoot)
	return newRoot
}

func (a *AVLTree) rotateRight(root *treeNode) *treeNode {
	newRoot := root.LeftChild
	root.LeftChild = newRoot.RightChild
	newRoot.RightChild = root

	a.setHeight(root)
	a.setHeight(newRoot)
	return newRoot
}

func (a *AVLTree) height(node *treeNode) int {
	if node == nil {
		return -1
	}
	return node.height
}

func (a *AVLTree) balanceFactor(node *treeNode) int {
	return a.height(node.LeftChild) - a.height(node.RightChild)
}

func (a *AVLTree) isLeftHeavy(node *treeNode) bool {
	return a.balanceFactor(node) > 1
}

func (a *AVLTree) isRightHeavy(node *treeNode) bool {
	return a.balanceFactor(node) < -1
}

func (a *AVLTree) IsEmpty() bool {
	return a.size == 0
}
