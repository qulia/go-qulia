package trie

import (
	"github.com/qulia/go-qulia/lib"
)

const (
	root          = "root"
	rootChar      = '\u00AE'
	terminate     = "terminate"
	terminateChar = '\u00A5'
)

type UpdateMDataFunc func(char rune, word []rune, metadata lib.Metadata)

type Interface interface {
	// Inserts a word into trie
	Insert(word []rune)

	// Searches for prefix in the trie, indicated by bool
	// if exist, returns metadata stored corresponding to last char
	Search(prefix []rune) (lib.Metadata, bool)

	// Searches for whole word in the trie,
	// returns true only if the word is added to trie before with all characters
	Contains(word []rune) bool
}

type node struct {
	char     rune
	children map[rune]*node
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

func (t *Trie) Insert(word []rune) {
	t.root.insert(word, 0, t.updateMDataFunc)
}

func (t *Trie) Search(prefix []rune) (lib.Metadata, bool) {
	if foundAt, ok := t.root.search(prefix, 0); ok {
		return foundAt.mData, true
	}

	return nil, false
}

func (t *Trie) Contains(word []rune) bool {
	if foundAt, ok := t.root.search(word, 0); ok {
		if foundAt.hasTerminate() {
			return true
		}
	}

	return false
}

func (n *node) insert(word []rune, index int, updateMDataFunc UpdateMDataFunc) {
	if index >= len(word) {
		n.checkAndAddTerminate()
		return
	}

	char := word[index]
	if _, ok := n.children[char]; !ok {
		n.children[char] = newNode(char)
	}

	n.children[char].checkAndCallUpdateMDataFunc(updateMDataFunc, word)
	n.children[char].insert(word, index+1, updateMDataFunc)
}

func (n *node) checkAndCallUpdateMDataFunc(updateMDataFunc UpdateMDataFunc, word []rune) {
	if updateMDataFunc != nil {
		updateMDataFunc(n.char, word, n.mData)
	}
}

func (n *node) search(prefix []rune, index int) (*node, bool) {
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

func newNode(char rune) *node {
	n := node{
		char:     char,
		mData:    lib.Metadata{},
		children: make(map[rune]*node),
	}

	return &n
}

func createRootNode() *node {
	n := newNode(rootChar)
	n.mData[root] = true
	return n
}

func createTerminateNode() *node {
	n := newNode(terminateChar)
	n.mData[terminate] = true
	return n
}
