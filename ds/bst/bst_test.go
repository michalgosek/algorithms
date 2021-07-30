package bst_test

import (
	"algorithms/ds/bst"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBST_IsEmpty(t *testing.T) {
	var b bst.BST
	got := b.IsEmpty()
	if got != true {
		t.Errorf("expected to get empty tree; got %t", got)
	}
}
func TestBST_Add(t *testing.T) {
	var b bst.BST
	nodes := []string{"A", "B", "C"}

	for _, n := range nodes {
		got := b.Add(n)
		if got != true {
			t.Errorf("expected to get true after adding node A")
		}
	}
}

func TestBST_Remove(t *testing.T) {
	var b bst.BST

	addOp := b.Add("A")
	if addOp != true {
		t.Fatalf("expected to get true after adding root node A; got %t", addOp)
	}

	size := b.Size()
	if size != 1 {
		t.Errorf("expected to get size equals one, got %d", size)
	}

	// Removing element does not exist
	removeOp := b.Remove("B")
	if removeOp != false {
		t.Errorf("expected to get false after removing non existing node B; got %t", removeOp)
	}

	size = b.Size()
	if size != 1 {
		t.Errorf("expected to get size equals one; got %d", size)
	}

	// Removing elements does exist
	addOp = b.Add("B")
	if addOp != true {
		t.Errorf("expected to get true after adding node B; got %t", addOp)
	}

	size = b.Size()
	if size != 2 {
		t.Errorf("expected to get size equals two; got %d", size)
	}

	removeOp = b.Remove("B")
	if removeOp != true {
		t.Fatalf("expected to get true after removing existing node B; got %t", removeOp)
	}

	size = b.Size()
	if size != 1 {
		t.Errorf("expected to get size equals one after removing node B; got %d", size)
	}

	height := b.Height()
	if size != 1 {
		t.Errorf("expected to get height equals one after removing node B; got %d", size)
	}

	// Removing the root
	removeOp = b.Remove("A")
	if removeOp != true {
		t.Fatalf("expected to get true after root node A; got %t", removeOp)
	}

	size = b.Size()
	if size != 0 {
		t.Errorf("expected to get size equals zero; got %d", size)
	}

	height = b.Height()
	if height != 0 {
		t.Errorf("expected to get height equals zero; got %d", height)
	}
}

func TestBST_Contains(t *testing.T) {
	var b bst.BST
	b.Add("A")
	b.Add("B")
	b.Add("C")
	assert.Equal(t, b.Contains("A"), true)
	assert.Equal(t, b.Contains("C"), true)
	assert.Equal(t, b.Contains("D"), false)
}

func TestBST_Height(t *testing.T) {
	var b bst.BST
	// Tree should look like:
	//        M
	//      J  S
	//    B   N Z
	//  A
	assert.Equal(t, b.Height(), 0)

	// Layer one
	b.Add("M")
	assert.Equal(t, b.Height(), 1)

	// Layer two
	b.Add("J")
	b.Add("S")
	assert.Equal(t, b.Height(), 2)

	// Layer three
	b.Add("B")
	assert.Equal(t, b.Height(), 3)
	b.Add("N")
	assert.Equal(t, b.Height(), 3)
	b.Add("Z")
	assert.Equal(t, b.Height(), 3)

	// Layer four
	b.Add("A")
	assert.Equal(t, b.Height(), 4)
}

func TestBST_Size(t *testing.T) {
	var b bst.BST
	assert.Equal(t, b.Size(), 0)

	b.Add("A")
	assert.Equal(t, b.Size(), 1)
}

func ExampleBST_InOrder() {
	// Tree should look like:
	//        M
	//      J  S
	//    B   N Z
	//  A

	var b bst.BST
	b.Add("M")
	b.Add("J")
	b.Add("S")
	b.Add("B")
	b.Add("N")
	b.Add("Z")
	b.Add("A")
	b.InOrder()
	//Output:
	// A
	// B
	// J
	// M
	// N
	// S
	// Z
}

func ExampleBST_PreOrder() {
	// Tree should look like:
	//        M
	//      J  S
	//    B   N Z
	//  A

	var b bst.BST
	b.Add("M")
	b.Add("J")
	b.Add("S")
	b.Add("B")
	b.Add("N")
	b.Add("Z")
	b.Add("A")
	b.PreOrder()
	//Output:
	// M
	// J
	// B
	// A
	// S
	// N
	// Z
}

func ExampleBST_LevelOrder() {
	// Tree should look like:
	//        M
	//      J  S
	//    B   N Z
	//  A

	var b bst.BST
	b.Add("M")
	b.Add("J")
	b.Add("S")
	b.Add("B")
	b.Add("N")
	b.Add("Z")
	b.Add("A")
	b.LevelOrder()
	//Output:
	// M
	// J
	// S
	// B
	// N
	// Z
	// A
}
