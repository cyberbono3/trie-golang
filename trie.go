package trie

import (
	"container/list"
	"sync"
)


type trieNode struct {
	leaves map[byte]*trieNode
	key    byte
	value  []byte
}

type trie struct {
	mutex sync.RWMutex
	node  *trieNode
	size  int
}

func NewNode(key byte) *trieNode {
	return &trieNode{leaves: make(map[byte]*trieNode), key: key}
}

func NewTrie() *trie {
	return &trie{
		node: &trieNode{leaves: make(map[byte]*trieNode)},
		size: 1,
	}
}



// Insert key value pair in a trie
func (t *trie) Insert(key, value []byte) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	curr := t.node

	for _, k := range key {
		if curr.leaves[k] == nil {
			curr.leaves[k] = NewNode(k)
		}
		curr = curr.leaves[k]
	}

	t.size++

	curr.value = value
}

func (t *trie) Size() int {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return t.size
}

func (t *trie) Search(key []byte) ([]byte, bool) {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	curr := t.node

	for _, k := range key {
		if curr.leaves[k] == nil {
			return nil,false
		}
		curr = curr.leaves[k]
	}
	return curr.value,true
}


func (t *trie) GetAllKeys() []byte {
	var keys []byte
	visited := make(map[*trieNode]bool)

	var dfsKeysFunc func(node *trieNode)
	dfsKeysFunc = func(node *trieNode) {
		if node != nil {
			visited[node] = true
			for k, n := range node.leaves {
				if _, ok := visited[n]; !ok {
					keys = append(keys, k)
					dfsKeysFunc(n)
				}
			}
		}

	}
	dfsKeysFunc(t.node)
	return keys

}
// Bytes is type alis
type Bytes []byte

func (t *trie) GetAllValues() []Bytes {
	queue :=  list.New()
	visited := make(map[*trieNode]bool)
	values := []Bytes{}

	queue.PushBack(t.node)
    // BFS
	for queue.Len() > 0 {
		first := queue.Front()
		queue.Remove(first)

		node := first.Value.(*trieNode)
		visited[node] = true

		if node.value != nil {
			values = append(values, node.value)
		}

		for _, leaf := range node.leaves {
			if _, ok := visited[leaf]; !ok {
				queue.PushBack(leaf)
			}
		}

	}

	return values

}
