package bst

import (
	"fmt"
	"math"
)

type Node struct {
	value string
	left  *Node
	right *Node
}

/*
BST is a tree is unidirected graph which staisfies any of
following defiintions:
	** An acylic connected graph
	** A connected graph with N nodes and N-1 edges
	** A graph in which any two vertices are connected
  by exactly one path

Used:
	* Implementation of some map and set ADTs
	* Red Black Trees
	* AVL Trees
	* Splay Trees
	* etc...

	* Used in the implementation of binary heaps
	* Syntax tree (used by compiler and calculators)
	* Treap - a probabilistic DS (uses a radnomized BST)

	Operation	| 	Avg		|	Worst   |
	---------------------------------------------------------
	Insert  	|	O(log(n))	|	O(n)	|
	Delele		|	O(log(n))   	|       O(n)    |
	Remove      	|   	O(log(n))	|       O(n)    |
	Search      	|	O(log(n))   	|       O(n)    |
*/
type BST struct {
	nodeCount int
	root      *Node
}

func (b *BST) IsEmpty() bool {
	return b.nodeCount == 0
}

func (b *BST) Size() int {
	return b.nodeCount
}

func (b *BST) Add(value string) bool {
	if b.Contains(value) {
		return false
	}

	b.root = b.add(b.root, value)
	b.nodeCount++
	return true
}

func (b *BST) add(node *Node, value string) *Node {
	if node == nil {
		node = &Node{value: value}
	}
	if value > node.value {
		node.right = b.add(node.right, value)
	}
	if value < node.value {
		node.left = b.add(node.left, value)
	}
	return node
}

func (b *BST) Contains(value string) bool {
	return b.contains(b.root, value)
}

func (b *BST) contains(node *Node, value string) bool {
	if node == nil {
		return false
	}
	if value > node.value {
		return b.contains(node.right, value)
	}
	if value < node.value {
		return b.contains(node.left, value)
	}
	return true
}

func (b *BST) Remove(value string) bool {
	if b.Contains(value) {
		b.root = b.remove(b.root, value)
		b.nodeCount--
		return true
	}
	return false
}

func (b *BST) remove(node *Node, value string) *Node {
	if node == nil {
		return nil
	}
	if value > node.value {
		node.right = b.remove(node.right, value)
	} else if value < node.value {
		node.left = b.remove(node.left, value)
	} else {
		if node.left == nil {
			return node.right
		}
		if node.right == nil {
			return node.left
		}
		tmp := b.findMax(node.left)
		node.value = tmp.value
		node.right = b.remove(node.right, tmp.value)
	}
	return node
}

func (b *BST) Height() int {
	return b.height(b.root)
}

func (b *BST) height(node *Node) int {
	if node == nil {
		return 0
	}
	return int(math.Max(float64(b.height(node.left)), float64(b.height(node.right)))) + 1
}

func (b *BST) findMin(node *Node) *Node {
	if node.left == nil {
		return node
	}
	return b.findMin(node.left)
}

func (b *BST) findMax(node *Node) *Node {
	if node.right == nil {
		return node
	}
	return b.findMax(node.right)
}

func (b *BST) PreOrder() {
	b.preOrder(b.root)
}

func (b *BST) preOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.value)
	b.preOrder(node.left)
	b.preOrder(node.right)
}

func (b *BST) InOrder() {
	b.inOrder(b.root)
}

func (b *BST) inOrder(node *Node) {
	if node == nil {
		return
	}
	b.inOrder(node.left)
	fmt.Println(node.value)
	b.inOrder(node.right)
}
