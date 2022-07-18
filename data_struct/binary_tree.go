package data_struct

import (
	"fmt"
	"math"
)

type treeNode struct {
	LeftChild  *treeNode
	RightChild *treeNode
	Value      int
	height     int
}

type BinaryTree struct {
	Root *treeNode
	Size int
}

// 二叉搜索树的特点：左边节点的值永远小于v 右边节点永远大于v
// 节点左侧的子树所有节点的值都应该小于v

// Find 是否存在v -> bool
func (b *BinaryTree) Find(v int) bool {
	var current = b.Root
	for current != nil {

		if v < current.Value {
			current = current.LeftChild
		} else if v > current.Value {
			current = current.RightChild
		} else {
			return true
		}
	}
	return false
}

func (b *BinaryTree) Insert(v int) int {
	defer func(b *BinaryTree) {
		b.Size++
	}(b)
	n := &treeNode{
		Value: v,
	}
	if b.Root == nil {
		b.Root = n
		return v
	}

	var current = b.Root

	for {
		if v < current.Value {
			if current.LeftChild == nil {
				current.LeftChild = n
				return v
			} else {
				current = current.LeftChild
			}

		} else {
			if current.RightChild == nil {
				current.RightChild = n
				return v
			} else {
				current = current.RightChild
			}
		}
	}

}

func (b *BinaryTree) Values(direction string) []interface{} {
	if direction != "left" && direction != "right" {
		fmt.Println("direction must be left or right")
		return nil
	}

	var ll = LinkedList{}
	if b.Size == 0 {
		return ll.Value()
	}
	current := b.Root
	for current != nil {
		if direction == "left" {
			ll.AddLast(current.Value)
			current = current.LeftChild
		} else if direction == "right" {
			ll.AddLast(current.Value)
			current = current.RightChild
		}

	}

	return ll.Value()
}

// TracersInOrder 自上而下层级遍历
func (b *BinaryTree) TracersInOrder() {
	b.tracersInOrder(b.Root)
}

func (b *BinaryTree) tracersInOrder(treeNode *treeNode) {
	// 终止条件
	if treeNode == nil {
		return
	}
	fmt.Println(treeNode.Value)
	b.tracersInOrder(treeNode.LeftChild)
	b.tracersInOrder(treeNode.RightChild)
}

// Height 返回当前二叉树Root节点高度
func (b *BinaryTree) Height() int {
	return b.height(b.Root)
}

// 计算高度的方法 = max(height(left), height(right)) + 1
func (b *BinaryTree) height(root *treeNode) int {
	if root == nil {
		return -1
	}
	if root.LeftChild == nil && root.RightChild == nil {
		return 0
	}
	return int(math.Max(float64(b.height(root.LeftChild)), float64(b.height(root.RightChild))) + 1)
}

// isLeaf 是否为叶节点，即没有子节点
func (b *BinaryTree) isLeaf(treeNode *treeNode) bool {
	return treeNode.LeftChild == nil && treeNode.RightChild == nil
}

// Equals 判断两个树是否一样
func (b *BinaryTree) Equals(other *BinaryTree) bool {
	return b.equals(b.Root, other.Root)
}

func (b *BinaryTree) equals(leftRoot *treeNode, rightRoot *treeNode) bool {
	if leftRoot == nil && rightRoot == nil {
		return true
	}

	if leftRoot != nil && rightRoot != nil {
		return leftRoot.Value == rightRoot.Value &&
			b.equals(leftRoot.LeftChild, rightRoot.LeftChild) &&
			b.equals(leftRoot.RightChild, rightRoot.RightChild)
	}
	return false
}

// IsBinarySearchTree 判断是否为二叉搜索树，既满足左侧永远小于节点，右侧大于节点
func (b *BinaryTree) IsBinarySearchTree() bool {
	return b.isBinarySearchTree(b.Root, math.MinInt, math.MaxInt)
}

func (b *BinaryTree) isBinarySearchTree(root *treeNode, min int, max int) bool {
	if root == nil {
		return true
	}
	return min < root.Value && root.Value <= max &&
		b.isBinarySearchTree(root.LeftChild, min, root.Value-1) &&
		b.isBinarySearchTree(root.RightChild, root.Value, max)
}

// GetDistanceArray 返回深度为k的节点数据数组
func (b *BinaryTree) GetDistanceArray(k int) []int {
	var arr []int
	arr = b.distanceK(b.Root, k, arr)
	return arr
}

func (b *BinaryTree) distanceK(root *treeNode, k int, arr []int) []int {
	if root == nil {
		return arr
	}

	if k == 0 {
		//fmt.Println(root.Value)
		arr = append(arr, root.Value)
	}
	arr = b.distanceK(root.LeftChild, k-1, arr)
	arr = b.distanceK(root.RightChild, k-1, arr)
	return arr
}

// PrintALL 按照深度升序打印节点数据
func (b *BinaryTree) PrintALL() {
	for i := 0; i <= b.Height(); i++ {
		for _, v := range b.GetDistanceArray(i) {
			fmt.Println(v)
		}
	}
}
