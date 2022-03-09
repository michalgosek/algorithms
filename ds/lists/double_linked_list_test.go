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

func TestGetNodeValueAt(t *testing.T) {
	var l lists.DoubleLinkedList
	l.AddFirst(1)
	l.AddFirst(2)
	l.AddFirst(3)
	l.AddLast(8)

	got := l.ValueAt(0)
	if got != 3 {
		t.Fatalf("expected to get value equals 3; got %d", got)
	}
	got = l.ValueAt(1)
	if got != 2 {
		t.Fatalf("expected to get value equals 2; got %d", got)
	}
	got = l.ValueAt(2)
	if got != 1 {
		t.Fatalf("expected to get value equals 1; got %d", got)
	}
	got = l.ValueAt(3)
	if got != 8 {
		t.Fatalf("expected to get value equals 8; got %d", got)
	}
}

func TestAddAtElementToTheList(t *testing.T) {
	var l lists.DoubleLinkedList
	l.AddAt(0, 1)
	l.AddAt(1, 2)
	l.AddAt(1, 3)
	l.AddAt(2, 4)
	l.AddAt(1, 8)

	got := l.Size()
	if got != 5 {
		t.Fatalf("expected to get length equals 5; got %d", got)
	}

	expectedValues := []int{1, 8, 3, 4, 2}
	for i, v := range expectedValues {
		got := l.ValueAt(i)
		if got != v {
			t.Fatalf("expected to get %d at %d index", v, i)
		}
	}
}

func TestShouldPanicAfterRemoveFirstWhenEmpty(t *testing.T) {
	var l lists.DoubleLinkedList
	assertPanic(t, l.RemoveFirst)
}

func TestShouldPanicAfterAddingAtNegativeIndex(t *testing.T) {
	var l lists.DoubleLinkedList
	defer func() {
		recover()
	}()
	l.AddAt(-1, 0)
	t.Fatal("expected to panic after calling func with negative idx")
}

func assertPanic(t *testing.T, f func()) {
	defer func() { recover() }()
	f()
	t.Fatal("expected to panic after calling test func")
}
