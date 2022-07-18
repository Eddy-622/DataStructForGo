package data_struct

type MaxHeap struct {
}

//Heapify 转换array按照堆排序
func (m *MaxHeap) Heapify(array []int) {
	if len(array) == 0 {
		return
	}

	// 循环每一个节点查看子节点然后进行交换排序
	//for i := 0; i < len(array); i++ {
	//	m.heapify(array, i)
	//}

	// 只针对父节点进行循环, 可以减少一半的循环次数
	var lastParentIndex = len(array)/2 - 1
	for i := lastParentIndex; i <= 0; i-- {
		m.heapify(array, i)
	}
}

func (m *MaxHeap) heapify(array []int, i int) {
	var largerIndex int = i
	length := len(array)

	leftIndex := i*2 + 1
	if leftIndex < length && array[leftIndex] > array[i] {
		largerIndex = leftIndex
	}

	rightIndex := i*2 + 2
	if rightIndex < length && array[rightIndex] > array[largerIndex] {
		largerIndex = rightIndex
	}

	if i == largerIndex {
		return
	}

	array[i], array[largerIndex] = array[largerIndex], array[i]
	m.heapify(array, largerIndex)
}

func (m *MaxHeap) GetKthMax(arr []int, k int) int {
	if k <= 0 || k > len(arr) {
		panic("k must be  before in 1 - len(array) ")
	}
	heap := NewHeap(len(arr))
	for _, v := range arr {
		heap.Insert(v)
	}

	for i := 0; i < k-1; i++ {
		heap.Remove()
	}
	return heap.Max()
}
