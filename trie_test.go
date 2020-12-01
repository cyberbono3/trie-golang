package trie

import (
	"testing"
	"bytes"
)



func TestNewTrie(t *testing.T) {
	trie := NewTrie()

	if trie.node == nil {
		t.Errorf("trie root is invalid. expected: %v, got: %v", &trieNode{}, trie.node)
	}

	if trie.size != 1 {
		t.Errorf("trie size is invalid. expected: %v, got: %v", 1, trie.size)
	}
}

func TestTrieSize(t *testing.T) {
	trie := NewTrie()

	if trie.size != 1 {
		t.Errorf("trie size is invalid. expected: % v, got: % v", 1, trie.Size())
	}
}

func TestTrieInsert(t *testing.T) {
	trie := NewTrie()
	size := 1

	testMap := map[string][]byte{
		"one":   []byte{1, 2},
		"two":   []byte{2, 3},
		"three": []byte{3, 4},
		"four":  []byte{4, 5},
		"five":  []byte{5, 6},
	}

	for s, b := range testMap {
		trie.Insert([]byte(s), b)
		size++

		if trie.Size() != size {
			t.Errorf("trie size is invalid. expected: %v, got: %v", size, trie.Size())
		}
	}
}

func TestTrieSearch(t *testing.T) {
	trie := NewTrie()

	testMap := map[string][]byte{
		"one":   []byte{1, 2},
		"two":   []byte{2, 3},
		"three": []byte{3, 4},
		"four":  []byte{4, 5},
		"five":  []byte{5, 6},
	}

	for k, v := range testMap {
		trie.Insert([]byte(k), v)
	}
	for k, v := range testMap {
		rv, ok := trie.Search([]byte(k))

		if !ok {
			t.Errorf("unable to find a key, expected: %v, got: %v", true, ok)
		}

		if !bytes.Equal(v, rv){
			t.Errorf("incorrect value for a key %v, expected: %v, got: %v", k, v, rv)
		}
	}

	invalidKey := []byte("invalid key")
	rv, ok := trie.Search(invalidKey)
	if ok {
		t.Errorf("Invalid key %v is not present in a trie, expected: %v, got: %v", invalidKey, false, ok)
	}

	if len(rv) != 0 {
		t.Errorf("invalid value for key %v. expected: %v, got: %v", invalidKey, []byte{}, rv)
	}

}

func TestGetAllValues(t *testing.T) {
	trie := NewTrie()

	vals := trie.GetAllValues()

	if len(vals) != 0 {
		t.Errorf("invalid length of values returned. expected: %v, got %v", 0, len(vals))
	}

	testMap := map[string][]byte{
		"one":   []byte{1, 2},
		"two":   []byte{2, 3},
		"three": []byte{3, 4},
		"four":  []byte{4, 5},
		"five":  []byte{5, 6},
	}

	for s, b := range testMap {
		trie.Insert([]byte(s), b)
	}

	testCases := []map[string]interface{}{
		map[string]interface{}{
			"expectedLen": 5,
			"expectedValues": []Bytes{
				[]byte{1, 2},
				[]byte{2, 3},
				[]byte{3, 4},
				[]byte{4, 5},
				[]byte{5, 6},
			},
		},
	}

	for _, tc := range testCases {
		vals = trie.GetAllValues()

		if len(vals) != tc["expectedLen"].(int) {
			t.Errorf("invalid length of values returned. expected: %v, got: %v",
				tc["expectedLen"].(int),
				len(vals),
			)
		}

		if !byteSliceEq(vals , tc["expectedValues"].([]Bytes) ){
			t.Errorf("missing value from expected list of values. expected: %v, got: %v",
				tc["expectedValues"].([]Bytes),
				vals,
			)
		}
	}
}

func TestGetAllKeys(t *testing.T) {
	trie := NewTrie()

	retrievedKeys := trie.GetAllKeys()

	if len(retrievedKeys) != 0 {
		t.Errorf("invalid length of keys returned. expected: %v, got %v", 0, len(retrievedKeys))
	}

	testMap := map[string][]byte{
		"one":   []byte{1, 2},
		"two":   []byte{2, 3},
		"three": []byte{3, 4},
		"four":  []byte{4, 5},
		"five":  []byte{5, 6},
	}

	for s, b := range testMap {
		trie.Insert([]byte(s), b)
	}

	testCases := []map[string]interface{}{
		map[string]interface{}{
			"expectedLen": 5,
			"expectedKeys": []Bytes{
				Bytes("one"),
				Bytes("two"),
				Bytes("three"),
				Bytes("four"),
				Bytes("five"),
			},
		},
	}

	for _, tc := range testCases {
		retrievedKeys = trie.GetAllKeys()

		if len(retrievedKeys) != tc["expectedLen"].(int) {
			t.Errorf("invalid length of keys returned. expected: %v, got %v",
				tc["expectedLen"].(int),
				len(retrievedKeys),
			)
		}

		if !bytes.Equal(retrievedKeys, tc["expectedKeys"].([]byte)) {
			t.Errorf("missing key from expected list of keys. expected: %v, got %v",
				tc["expectedKeys"].([]byte),
				retrievedKeys ,
			)
		}
	}
}
