package data_struct

type Heap struct {
	item []int
	size int
}

func NewHeap(itemize int) *Heap {
	return &Heap{item: make([]int, itemize)}
}

func (h *Heap) Insert(v int) {
	if h.IsFull() {
		panic("堆已满")
	}

	h.item[h.size] = v
	h.size++

	h.bubbleUp() // 向上冒泡
}

func (h *Heap) bubbleUp() {
	index := h.size - 1
	for index > 0 && h.item[index] > h.item[h.parent(index)] {
		h.swap(index, h.parent(index))
		index = h.parent(index)
	}
}

// parent 父节点算法 （子节点索引 - 1 ）/ 2
func (h *Heap) parent(index int) int {
	return (index - 1) / 2
}

// leftChild 左子节点算法 (父节点 * 2) + 1
func (h *Heap) leftChild(index int) int {
	return h.item[h.leftChildIndex(index)]
}

// rightChild 右子节点算法 (父节点 * 2) + 2
func (h *Heap) rightChild(index int) int {
	return h.item[h.rightChildIndex(index)]
}

// leftChildIndex 左子节点算法 (父节点 * 2) + 1
func (h *Heap) leftChildIndex(index int) int {
	return (index * 2) + 1
}

// rightChildIndex 右子节点算法 (父节点 * 2) + 2
func (h *Heap) rightChildIndex(index int) int {
	return (index * 2) + 2
}

// 交换两个节点
func (h *Heap) swap(index, parent int) {
	h.item[index], h.item[parent] = h.item[parent], h.item[index]
}

// IsFull 是否满
func (h *Heap) IsFull() bool {
	return h.size == len(h.item)
}

// IsEmpty 是否满
func (h *Heap) IsEmpty() bool {
	return h.size == 0
}

// Remove 删除主节点
func (h *Heap) Remove() int {
	if h.size == 0 {
		return 0
	}
	root := h.item[0]

	h.item[0] = h.item[h.size-1]
	h.size--

	h.bubbleDown()
	return root
}

func (h *Heap) bubbleDown() {
	var index = 0

	// 索引在有效范围内，且父节点不符合要求
	for index <= h.size && !h.isValidParent(index) {
		// 找出最大的子节点索引
		var largerChildIndex = h.largerChildIndex(index)
		// 最大的子节点和父节点交换
		h.swap(index, largerChildIndex)
		// 下一次的检查索引更换为当前循环的最大子节点
		index = largerChildIndex
	}
}

// isValidParent 判断当前节点如果是父节点是否合规：即父节点 < 左右子节点
// 如果不存在左节点： true 【即没有子节点 -> 合规】
// isValid：父节点大于左节点 &&【 父节点大于右节点(如果右节点存在) 】
func (h *Heap) isValidParent(index int) bool {
	if !h.hasLeftChild(index) {
		return true
	}

	var isValid = h.item[index] >= h.leftChild(index)

	if h.hasRightChild(index) {
		isValid = isValid && h.item[index] >= h.rightChild(index)
	}
	return isValid
}

// largerChildIndex：返回最大的子节点索引
// 如果不存在左节点，则返回原节点，不存在左节点即不存在右节点，说明当前节点为最小子节点
// 如果存在左节点但不存在右节点，直接返回左节点， 因为可判断当前节点不合规且没有右节点，即左节点一定大于父节点
// 左右节点同时存在则对比左右节点大小，返回较大的节点
func (h *Heap) largerChildIndex(index int) int {
	if !h.hasLeftChild(index) {
		return index
	}

	if !h.hasRightChild(index) {
		return h.leftChildIndex(index)
	}

	if h.leftChild(index) > h.rightChild(index) {
		return h.leftChildIndex(index)
	} else {
		return h.rightChildIndex(index)
	}

}

func (h *Heap) hasRightChild(index int) bool {
	return h.rightChildIndex(index) <= h.size
}
func (h *Heap) hasLeftChild(index int) bool {
	return h.leftChildIndex(index) <= h.size
}

func (h *Heap) Max() int {
	if h.IsEmpty() {
		panic("堆为空")
	}
	return h.item[0]
}
