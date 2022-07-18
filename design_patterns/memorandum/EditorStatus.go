package memorandum

// 可以理解成备忘录本身

type EditorStatus struct {
	context string
}

func NewEditorStatus(context string) *EditorStatus {
	return &EditorStatus{context: context}
}

func (e EditorStatus) Context() string {
	return e.context
}
