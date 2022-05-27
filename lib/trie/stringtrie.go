package trie

import "github.com/qulia/go-qulia/lib/set"

type node struct {
	char      byte
	children  map[byte]*node
	terminate bool

	words set.Set[string]
}

type trieImpl struct {
	root *node
}

func newTrieImpl() *trieImpl {
	return &trieImpl{
		root: createRootNode(),
	}
}

func (t *trieImpl) Insert(word string) {
	t.root.insert(word, 0)
}

func (t *trieImpl) Search(prefix string) set.Set[string] {
	if foundAt, ok := t.root.search(prefix, 0); ok {
		return foundAt.words
	}

	return nil
}

func (t *trieImpl) Contains(word string) bool {
	if foundAt, ok := t.root.search(word, 0); ok {
		if foundAt.hasTerminate() {
			return true
		}
	}

	return false
}

func (n *node) insert(word string, index int) {
	n.words.Add(word)
	if index >= len(word) {
		n.checkAndAddTerminate()
		return
	}

	char := word[index]
	if _, ok := n.children[char]; !ok {
		n.children[char] = newNode(char)
	}
	n.children[char].insert(word, index+1)
}

func (n *node) search(prefix string, index int) (*node, bool) {
	if index == len(prefix) {
		return n, true
	}

	char := prefix[index]
	if child, ok := n.children[char]; ok {
		return child.search(prefix, index+1)
	}

	return nil, false
}

func (n *node) checkAndAddTerminate() {
	if !n.hasTerminate() {
		n.children[terminateChar] = createTerminateNode()
	}
}

func (n *node) hasTerminate() bool {
	_, ok := n.children[terminateChar]
	return ok
}

func newNode(char byte) *node {
	n := node{
		char:     char,
		children: make(map[byte]*node),
		words:    set.NewSet[string](),
	}

	return &n
}

func createRootNode() *node {
	n := newNode(rootChar)
	return n
}

func createTerminateNode() *node {
	n := newNode(terminateChar)
	n.terminate = true
	return n
}
