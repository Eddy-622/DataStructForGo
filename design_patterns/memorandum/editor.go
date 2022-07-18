package memorandum

// 备忘录模式用来创造用来可撤销操作的类

// 备忘录调度类

type Editor struct {
	content string
	history *History
}

func NewEditor() *Editor {
	return &Editor{
		history: NewHistory(),
	}
}

func (e *Editor) Content() string {
	return e.content
}

func (e *Editor) SetContent(content string) {
	e.content = content
	e.createStatus()
}

func (e *Editor) createStatus() {
	status := NewEditorStatus(e.content)
	e.history.Push(status)
}

func (e *Editor) Back() {
	content := e.history.Pop()
	if c, ok := content.(*EditorStatus); ok {
		e.content = c.context
	}
	if content == nil {
		e.content = ""
	}
}
