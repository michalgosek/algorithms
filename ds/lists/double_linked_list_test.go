package lists_test

import (
	"algorithms/ds/lists"
	"testing"
)

func TestShouldReturnZeroAfterCallingSizeOnEmptyList(t *testing.T) {
	var l lists.DoubleLinkedList
	size := l.Size()
	if size != 0 {
		t.Fatal("expected to get size equals 0")
	}
}

func TestShouldReturnTrueAfterCallingIsEmptyOnEmptyList(t *testing.T) {
	var l lists.DoubleLinkedList
	got := l.IsEmpty()
	if got != true {
		t.Fatal("expected to get empty list")
	}
}

func TestShouldAddFirstElementToTheList(t *testing.T) {
	var l lists.DoubleLinkedList
	l.AddFirst(3)
	got := l.Size()
	if got != 1 {
		t.Fatal("expected to get size equals 1")
	}
	l.AddFirst(5)
	got = l.Size()
	if got != 2 {
		t.Fatal("expected to get size equals 2")
	}
}

func TestShouldAddLastElementToTheList(t *testing.T) {
	var l lists.DoubleLinkedList

	l.AddLast(3)
	got := l.Size()
	if got != 1 {
		t.Fatal("expected to get size equals 1")
	}
	l.AddLast(5)
	got = l.Size()
	if got != 2 {
		t.Fatal("expected to get size equals 2")
	}
}

func TestShouldPanicAfterRemoveFirstWhenEmpty(t *testing.T) {
	var l lists.DoubleLinkedList
	assertPanic(t, l.RemoveFirst)
}

func assertPanic(t *testing.T, f func()) {
	defer func() { recover() }()
	f()
	t.Fatal("expected to panic after calling test func")
}
