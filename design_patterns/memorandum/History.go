package memorandum

import "container/list"

type History struct {
	historyList *list.List
	nowStatus   *EditorStatus
}

func NewHistory() *History {
	return &History{historyList: list.New()}
}

func (h *History) Push(s *EditorStatus) {
	if h.nowStatus != nil {
		h.historyList.PushBack(h.nowStatus)
	}
	h.nowStatus = s
}

func (h *History) Pop() any {
	if h.IsEmpty() {
		return nil
	}
	lastItem := h.historyList.Back()
	return h.historyList.Remove(lastItem)
}

func (h *History) IsEmpty() bool {
	return h.historyList.Len() == 0
}
