package memorandum

import "testing"

// 备忘录模式用来创建可撤销操作的类
// editor status history 分别负责各自的功能

func TestMem(t *testing.T) {
	m := NewEditor()
	m.SetContent("a")
	//m.SetContent("b")
	//m.SetContent("c")

	m.Back()
	//m.SetContent("d")
	t.Log(m.Content())

}
