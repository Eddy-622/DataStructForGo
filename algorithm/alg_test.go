package algorithm

import (
	"testing"
)

var l = RandomIntList(10)

func TestRandomIntList(t *testing.T) {
	t.Log(RandomIntList(10))
}

func TestSortBubbleUp(t *testing.T) {
	SortBubbleUp(l)
	t.Log(l)
}

func TestSortChose(t *testing.T) {
	SortChose(l)
	t.Log(l)
}

func TestMap(t *testing.T) {

	m := make(map[string]int)
	m["1"] = 1
	t.Log(m)

}
