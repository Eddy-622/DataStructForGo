package algorithm

// SortBubbleUp 冒泡排序O(n**2) 分别找到第x大的数防到x位上
func SortBubbleUp(list []int) {
	// [10  20  30]
	//  c    n
	lastIndex := len(list) - 1
	for i := 0; i < lastIndex; i++ {
		// 外层循环 第X次  第X大的数
		for j := 0; j < lastIndex-i; j++ {
			// 内层循环 前后对比 进行对调
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
}

//SortChose 选择排序O(n**2)
func SortChose(list []int) {
	// [10  30  20 ]
	//  c   n    l
	// if l > c  :  pass  if l < n : n,l = l,n
	lastIndex := len(list) - 1
	for i := 0; i < lastIndex; i++ {
		max := lastIndex - i // 假设最后一位的数字最大 并记录下标
		for j := 0; j < lastIndex-i; j++ {
			if list[j] > list[max] {
				max = j // 如果下标j大于下标max，max=j
			}

		}
		if max != lastIndex-i {
			// 如果下标发生改变， 进行数据替换
			list[max], list[lastIndex-i] = list[lastIndex-i], list[max]
		}
	}
}
