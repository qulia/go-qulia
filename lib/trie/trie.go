package trie

import "github.com/qulia/go-qulia/lib"

const (
	root     = "root"
	rootChar = '.'
)

type UpdateMDataFunc func(char uint8, word string, metadata lib.Metadata)

type Interface interface {
	Insert(word string)
	Search(prefix string) (lib.Metadata, bool)
}

type node struct {
	char     uint8
	children map[uint8]*node
	mData    lib.Metadata
}

type Trie struct {
	root            *node
	updateMDataFunc UpdateMDataFunc
}

func NewTrie(updateMDataFunc UpdateMDataFunc) *Trie {
	t := Trie{
		root:            createRootNode(),
		updateMDataFunc: updateMDataFunc,
	}
	return &t
}

func (t *Trie) Insert(word string) {
	t.root.insert(word, 0, t.updateMDataFunc)
}

func (t *Trie) Search(prefix string) (lib.Metadata, bool) {
	return t.root.search(prefix, 0)
}

func (n *node) insert(word string, index int, updateMDataFunc UpdateMDataFunc) {
	if index >= len(word) {
		return
	}

	char := word[index]
	if _, ok := n.children[char]; !ok {
		n.children[char] = newNode(char)
	}

	n.children[char].checkAndCallUpdateMDataFunc(updateMDataFunc, word)
	n.children[char].insert(word, index+1, updateMDataFunc)
}

func (n *node) checkAndCallUpdateMDataFunc(updateMDataFunc UpdateMDataFunc, word string) {
	if updateMDataFunc != nil {
		updateMDataFunc(n.char, word, n.mData)
	}
}

func (n *node) search(prefix string, index int) (lib.Metadata, bool) {
	if index == len(prefix) {
		return n.mData, true
	}

	char := prefix[index]
	if child, ok := n.children[char]; ok {
		return child.search(prefix, index+1)
	}

	return nil, false
}

func newNode(char uint8) *node {
	n := node{
		char:     char,
		mData:    lib.Metadata{},
		children: make(map[uint8]*node),
	}

	return &n
}

func createRootNode() *node {
	n := newNode(rootChar)
	n.mData[root] = true
	return n
}
