package trie_test

import (
	"testing"

	"github.com/qulia/go-qulia/lib"
	"github.com/qulia/go-qulia/lib/trie"
	"github.com/stretchr/testify/assert"
)

func TestTrieBasic(t *testing.T) {
	words := []string{"abc", "abcd", "efg", "e", "fg", "hijklm"}
	testTrie := trie.NewTrie(func(_ uint8, word string, mData lib.Metadata) {
		if _, ok := mData["words"]; !ok {
			mData["words"] = []string{}
		}
		mData["words"] = append(mData["words"].([]string), word)
	})

	for _, word := range words {
		testTrie.Insert(word)
	}

	checkPrefix(t, testTrie, "ab", []string{"abc", "abcd"})
	checkPrefix(t, testTrie, "e", []string{"efg", "e"})
	checkPrefix(t, testTrie, "hijklm", []string{"hijklm"})
}

func checkPrefix(t *testing.T, testTrie trie.Interface, prefix string, expectedWords []string) {
	mData, ok := testTrie.Search(prefix)
	assert.True(t, ok)
	wordsWithPrefix := mData["words"].([]string)
	assert.Equal(t, expectedWords, wordsWithPrefix)
}
