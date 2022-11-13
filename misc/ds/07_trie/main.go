package main

import "fmt"

const AlphabetSize = 26

type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{
		root: &Node{},
	}
}

func (t *Trie) Insert(s string) {
	n := t.root
	for _, c := range s {
		if n.children[c-'a'] == nil {
			n.children[c-'a'] = &Node{}
		}
		n = n.children[c-'a']
	}
	n.isEnd = true
}

func (t *Trie) Search(s string) bool {
	n := t.root
	for _, c := range s {
		if n.children[c-'a'] == nil {
			return false
		}
		n = n.children[c-'a']
	}
	return n.isEnd
}

func main() {
	trie := NewTrie()
	fmt.Println(trie.root)
	trie.Insert("hello")
	fmt.Println("hello is in a trie:", trie.Search("hello"))
	fmt.Println("hellow is not in a trie:", trie.Search("hellow"))
	fmt.Println("ello is not in a trie:", trie.Search("ello"))
}
