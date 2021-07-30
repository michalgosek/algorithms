package queue

import "errors"

// IntSliceQueue provides queue implementation
// for integer values. Can be used as zero value.
type IntSliceQueue struct {
	data []int
}

// Size returns the size of the queue.
func (q *IntSliceQueue) Size() int {
	return len(q.data)
}

// IsEmpty returns ture if the queue is empty, otherwise
// returns false.
func (q *IntSliceQueue) IsEmpty() bool {
	return len(q.data) == 0
}

// Enqueue adds element to the end of the queue.
// If queue is empty appends it on the beggining.
func (q *IntSliceQueue) Enqueue(v int) bool {
	q.data = append(q.data, v)
	return true
}

// Peek returns the first element of the queue.
func (q *IntSliceQueue) Peek() (int, error) {
	if q.Size() == 0 {
		return -1, ErrEmptyQueue
	}
	return q.data[0], nil
}

// Dequeue returns the first elemet of the queue.
// Resizes the queue by removing returning element.
func (q *IntSliceQueue) Dequeue() (int, error) {
	if q.Size() == 0 {
		return -1, ErrEmptyQueue
	}
	last := q.data[0]
	q.data = q.data[1:]
	return last, nil
}

// ErrEmptyQueue happens during processing peek, dequeue operations on empty queue.
var ErrEmptyQueue = errors.New("operation failed due to zero length of the queue.")
