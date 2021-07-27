package bst_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBST_IsEmpty(t *testing.T) {
	var b BST
	assert.Equal(t, b.IsEmpty(), true)
}
func TestBST_Add(t *testing.T) {
	var b BST
	assert.Equal(t, b.Add("A"), true)
	assert.Equal(t, b.Add("B"), true)
	assert.Equal(t, b.Add("C"), true)
}

func TestBST_Remove(t *testing.T) {
	var b BST
	b.Add("A")
	assert.Equal(t, b.Size(), 1)
	assert.Equal(t, b.Remove("B"), false)
	assert.Equal(t, b.Size(), 1)

	// Removing elements does exist
	b.Add("B")
	assert.Equal(t, b.Size(), 2)
	assert.Equal(t, b.Remove("B"), true)
	assert.Equal(t, b.Size(), 1)
	assert.Equal(t, b.Height(), 1)

	// Removing the root
	assert.Equal(t, b.Remove("A"), true)
	assert.Equal(t, b.Size(), 0)
	assert.Equal(t, b.Height(), 0)
}

func TestBST_Contains(t *testing.T) {
	var b BST
	b.Add("A")
	b.Add("B")
	b.Add("C")
	assert.Equal(t, b.Contains("A"), true)
	assert.Equal(t, b.Contains("C"), true)
	assert.Equal(t, b.Contains("D"), false)
}

func TestBST_Height(t *testing.T) {
	var b BST
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
	var b BST
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

	var b BST
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

	var b BST
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
