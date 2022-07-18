package algorithm

import (
	"math/rand"
	"time"
)

//RandomIntList 随机生成n个0-100的切片
func RandomIntList(n int) []int {
	list := make([]int, n)
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		list[i] = rand.Intn(100)
	}
	return list
}
