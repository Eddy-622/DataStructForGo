package data_struct

type triesNode struct {
	value       rune
	child       map[rune]*triesNode
	isEndOfWord bool
}

type Tries struct {
	root *triesNode
}

func NewTries() *Tries {
	return &Tries{root: newTriesNode(0)}
}

func newTriesNode(value rune) *triesNode {
	return &triesNode{child: make(map[rune]*triesNode),
		value: value}
}

func (t *triesNode) String() string {
	return "value = " + string(t.value)
}

func (t *triesNode) hasChild(ch rune) bool {
	_, ok := t.child[ch]
	return ok
}

func (t *triesNode) addChild(ch rune) {
	t.child[ch] = newTriesNode(ch)
}

func (t *triesNode) getChild(ch rune) *triesNode {
	return t.child[ch]
}

func (t *Tries) Insert(word string) {
	current := t.root
	for _, c := range word {
		if !current.hasChild(c) {
			current.addChild(c)
		}
		current = current.getChild(c)
	}
	current.isEndOfWord = true
}
