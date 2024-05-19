package trie

import (
	"github.com/qulia/go-qulia/v2/lib/set"
)

const (
	root          = "root"
	rootChar      = '\u00AE'
	terminate     = "terminate"
	terminateChar = '\u00A5'
)

type Trie interface {
	// Inserts a word into trie
	Insert(word string)

	// Returns words with prefix
	Search(prefix string) set.Set[string]

	// Searches for whole word in the trie,
	// returns true only if the word is added to trie before with all characters
	Contains(word string) bool
}

func NewTrie() Trie {
	return newTrieImpl()
}
