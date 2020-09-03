package trie

import (
	"container/list"
	"sync"
)

type TrieNode struct {
	leaves map[byte]*TrieNode
	key    byte
	value  []byte
}

type Trie struct {
	mutex sync.RWMutex
	node  *TrieNode
	size  int
}

func NewNode(key byte) *TrieNode {
	return &TrieNode{leaves: make(map[byte]*TrieNode), key: key}
}

func NewTrie() *Trie {
	return &Trie{
		node: &TrieNode{leaves: make(map[byte]*TrieNode)},
		size: 1,
	}
}

// Insert key value pair in a tree
func (t *Trie) Insert(key, value []byte) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	curr := t.node

	if key != nil {
		for _, k := range key {
			if curr.leaves[k] == nil {
				curr.leaves[k] = NewNode(k)
			}
			curr = curr.leaves[k]
		}
	}

	curr.value = value
}

func (t *Trie) Size() int {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return t.size
}

func (t *Trie) Search(key []byte) ([]byte, bool) {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	curr := t.node

	for _, k := range key {
		if curr.leaves[k] == nil {
			return nil, false
		}
		curr = curr.leaves[k]
	}
	return curr.value, true
}

func GetKeysByDFS(node *TrieNode, visited map[*TrieNode]bool, key, keys []byte) {
	if node != nil {
		keyList := append(key, node.key)
		visited[node] = true

		if node.value != nil {
			fullKeyList := make([]byte, len(keyList))

			copy(fullKeyList, keyList)

			// we ignore the first byte which is root key
			keys = append(keys, fullKeyList[1:])
		}
		for _, leaf := range node.leaves {
			if _, ok := visited[leaf]; !ok {
				GetKeysByDFS(leaf, visited, keyList, keys)
			}
		}
	}
}

func (t *Trie) GetAllKeys() []byte {
	keys := []byte{}
	visited := make(map[*TrieNode]bool)
	GetKeysByDFS(t.node, visited, []byte{}, keys)
	return keys

}

func (t *Trie) GetAllValues() []byte {
	q = list.New()
	visited := make(map[*TrieNode]bool)
	values := []byte{}

	q.PushBack(t.node)

	for q.Len() > 0 {
		item := q.Front()
		q.Remove(item)

		node := item.Value.(*TrieNode)
		visited[node] = true

		if node.value != nil {
			values = append(values, node.value)
		}

		for _, leaf := range node.leaves {
			if _, ok := visited[leaf]; !ok {
				q.PushBack(leaf)
			}
		}

	}

	return values

	// ToDO GetPrefixKeys and GetPrefixValues
}
