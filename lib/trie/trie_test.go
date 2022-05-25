package trie_test

import (
	"testing"

	"github.com/qulia/go-qulia/lib/set"
	"github.com/qulia/go-qulia/lib/trie"
	"github.com/stretchr/testify/assert"
)

func TestTrieBasic(t *testing.T) {
	words := []string{"abc", "abcd", "efg", "e", "fg", "hijklm"}
	testTrie := trie.NewTrie()
	addWords(words, testTrie)

	checkPrefix(t, testTrie, "ab", []string{"abc", "abcd"})
	checkPrefix(t, testTrie, "e", []string{"efg", "e"})
	checkPrefix(t, testTrie, "hijklm", []string{"hijklm"})
}

func checkPrefix(t *testing.T, testTrie trie.Trie, prefix string, expectedWords []string) {
	result := testTrie.Search(prefix)
	expected := set.NewSet[string]()
	expected.FromSlice(expectedWords)
	assert.NotNil(t, result)
	assert.Equal(t, expected, result)
}

func TestTriePrefixDoesNotExist(t *testing.T) {
	words := []string{"abc", "abcd", "efg", "e", "fg", "hijklm"}
	testTrie := trie.NewTrie()
	addWords(words, testTrie)

	assert.Nil(t, testTrie.Search("x"))
}

func addWords(words []string, testTrie trie.Trie) {
	for _, word := range words {
		testTrie.Insert(word)
	}
}

func TestTrie_Contains(t *testing.T) {
	words := []string{"abc", "abcd", "efg", "e", "fg", "hijklm"}
	testTrie := trie.NewTrie()
	addWords(words, testTrie)

	for _, word := range words {
		assert.True(t, testTrie.Contains(word))
	}

	assert.False(t, testTrie.Contains(""))
	assert.False(t, testTrie.Contains("ab"))
	assert.False(t, testTrie.Contains("fg "))

	testTrie.Insert("")
	assert.True(t, testTrie.Contains(""))
}
