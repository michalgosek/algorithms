package lists

type node struct {
	val         interface{}
	left, right *node
}

type DoubleLinkedList struct {
	length int
	root   *node
}

func (d *DoubleLinkedList) AddFirst(v int) {
	if d.root == nil {
		d.root = &node{
			left:  nil,
			right: nil,
			val:   v,
		}
	} else {
		node := node{
			left:  nil,
			right: nil,
			val:   v,
		}
		head := d.root
		d.root = &node
		head.left = &node
		d.root.right = head
	}
	d.length++
}

func (d *DoubleLinkedList) AddLast(v int) {
	if d.length == 0 {
		d.AddFirst(v)
		return
	}

	curr := d.root
	for curr.right != nil {
		curr = curr.right
	}

	node := &node{
		left:  nil,
		right: nil,
		val:   v,
	}

	curr.right = node
	node.left = curr
	d.length++
}

func (d *DoubleLinkedList) Size() int {
	return d.length
}

func (d *DoubleLinkedList) IsEmpty() bool {
	return d.length == 0
}

func (d *DoubleLinkedList) RemoveFirst() {
	if d.root == nil {
		panic("no elements to remove!")
	}
}
