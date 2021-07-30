package queue_test

import (
	"algorithms/ds/queue"
	"errors"
	"testing"
)

func TestQueue_Enqueue(t *testing.T) {
	var q queue.IntSliceQueue

	enqueuOp := q.Enqueue(1)
	if enqueuOp != true {
		t.Errorf("expected to get true after enqueue one; got %t", enqueuOp)
	}
	emptyOp := q.IsEmpty()
	if emptyOp != false {
		t.Errorf("expected to get false after is empty; got %t", emptyOp)
	}
	queueSize := q.Size()
	if queueSize != 1 {
		t.Errorf("expected to get queue size equals one; got %d", queueSize)
	}
}

func TestQueue_Dequeue(t *testing.T) {
	var q queue.IntSliceQueue

	// Attempting to dequeue on zero length queue.
	_, err := q.Dequeue()
	if !errors.Is(err, queue.ErrEmptyQueue) {
		t.Errorf("expected to get err empty queue; got %v", err)
	}

	enqueuOp := q.Enqueue(1)
	if enqueuOp != true {
		t.Errorf("expected to get true after enequeue one; got %t", enqueuOp)
	}
	got, err := q.Dequeue()
	if err != nil {
		t.Errorf("expected to get nil err; got %v", err)
	}
	if got != 1 {
		t.Errorf("expected to get one after dequeue; got %d", got)
	}
}

func TestQueue_Peek(t *testing.T) {
	var q queue.IntSliceQueue
	// Attempting to peek on zero length queue.
	_, err := q.Peek()
	if !errors.Is(err, queue.ErrEmptyQueue) {
		t.Errorf("expected to get err empty queue; got %v", err)
	}

	// Adding single element.
	q.Enqueue(1)
	got, err := q.Peek()
	if err != nil {
		t.Errorf("expected to get nil err; got %v", err)
	}
	if got != 1 {
		t.Errorf("expected to get peek equals one; got %d", got)
	}
}
