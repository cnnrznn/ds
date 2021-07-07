package trie

import "sync"

type Trie struct {
	Root *Node
}

type Node struct {
	Rune     string
	Data     interface{}
	Children map[string]*Node
	sync.Mutex
}

func newNode(data string) *Node {
	return &Node{
		Rune:     data,
		Children: make(map[string]*Node),
	}
}

func New() *Trie {
	return &Trie{
		Root: newNode(""),
	}
}

func (t *Trie) Insert(prefix string, callback func(*Node)) {
	node := t.Root
	for _, r := range prefix {
		callback(node)
		if _, ok := node.Children[string(r)]; !ok {
			node.Children[string(r)] = &Node{
				Rune:     string(r),
				Children: make(map[string]*Node),
			}
		}
		node = node.Children[string(r)]
	}
	callback(node)
}

func (t *Trie) Find(prefix string) *Node {
	node := t.Root
	for _, r := range prefix {
		if n, ok := node.Children[string(r)]; ok {
			node = n
		} else {
			return nil
		}
	}
	return node
}
