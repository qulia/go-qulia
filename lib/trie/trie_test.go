package trie_test

import (
	"testing"

	"github.com/qulia/go-qulia/lib"

	"github.com/qulia/go-qulia/lib/trie"
	"github.com/stretchr/testify/assert"
)

func TestTrieBasic(t *testing.T) {
	words := []string{"abc", "abcd", "efg", "e", "fg", "hijklm"}
	testTrie := trie.NewTrie(func(_ rune, word []rune, mData lib.Metadata) {
		if _, ok := mData["words"]; !ok {
			mData["words"] = []string{}
		}
		mData["words"] = append(mData["words"].([]string), string(word))
	})
	addWords(words, testTrie)

	checkPrefix(t, testTrie, "ab", []string{"abc", "abcd"})
	checkPrefix(t, testTrie, "e", []string{"efg", "e"})
	checkPrefix(t, testTrie, "hijklm", []string{"hijklm"})
}

func checkPrefix(t *testing.T, testTrie trie.Interface, prefix string, expectedWords []string) {
	mData, ok := testTrie.Search([]rune(prefix))
	assert.True(t, ok)
	wordsWithPrefix := mData["words"].([]string)
	assert.Equal(t, expectedWords, wordsWithPrefix)
}

func TestTriePrefixDoesNotExist(t *testing.T) {
	words := []string{"abc", "abcd", "efg", "e", "fg", "hijklm"}
	testTrie := trie.NewTrie(nil)
	addWords(words, testTrie)

	_, ok := testTrie.Search([]rune("x"))
	assert.False(t, ok)
}

func addWords(words []string, testTrie *trie.Trie) {
	for _, word := range words {
		testTrie.Insert([]rune(word))
	}
}

func TestTrie_Contains(t *testing.T) {
	words := []string{"abc", "abcd", "efg", "e", "fg", "hijklm"}
	testTrie := trie.NewTrie(nil)
	addWords(words, testTrie)

	for _, word := range words {
		assert.True(t, testTrie.Contains([]rune(word)))
	}

	assert.False(t, testTrie.Contains([]rune("")))
	assert.False(t, testTrie.Contains([]rune("ab")))
	assert.False(t, testTrie.Contains([]rune("fg ")))

	testTrie.Insert([]rune(""))
	assert.True(t, testTrie.Contains([]rune("")))
}
