package linked_list

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var node Node
	node.data = 1

	add(&node, 2)
	add(&node, 3)
	add(&node, 4)

	listAsString := toString(node)

	if listAsString != "1, 4, 3, 2" {
		t.Errorf("listAsString() expected '1, 4, 3, 2', found '" + listAsString + "'")
	}
}

func TestSet(t *testing.T) {
	var node Node
	node.data = 1

	var nextNode Node
	nextNode.data = 2
	node.next = &nextNode

	set(&node, 3)

	listAsString := toString(node)

	if listAsString != "1, 3" {
		t.Errorf("listAsString() expected '1, 3', found '" + listAsString + "'")
	}
}

func TestCount(t *testing.T) {
	var node Node
	node.data = 4

	var node2 Node
	node2.data = 2
	node.next = &node2

	var node3 Node
	node3.data = 1
	node2.next = &node3

	var node4 Node
	node4.data = 3
	node3.next = &node4

	count := count(&node)

	if count != 4 {
		t.Errorf("listAsString() expected '4', found '" + intToString(count) + "'")
	}
}

func TestRemove(t *testing.T) {
	var node Node
	node.data = 1

	var nextNode Node
	nextNode.data = 2
	node.next = &nextNode

	remove(&node)

	listAsString := toString(node)

	if listAsString != "1" {
		t.Errorf("listAsString() expected '1', found '" + listAsString + "'")
	}
}

func TestSwap (t *testing.T) {
	var node Node
	node.data = 4

	var node2 Node
	node2.data = 2
	node.next = &node2

	var node3 Node
	node3.data = 1
	node2.next = &node3

	var node4 Node
	node4.data = 3
	node3.next = &node4

	swap(&node2, &node4)

	listAsString := toString(node)

	if listAsString != "4, 3, 1, 2" {
		t.Errorf("listAsString() expected '4, 3, 1, 2', found '" + listAsString + "'")
	}
}

func TestReverse(t *testing.T) {
	var node Node
	node.data = 4

	var node2 Node
	node2.data = 2
	node.next = &node2

	var node3 Node
	node3.data = 1
	node2.next = &node3

	var node4 Node
	node4.data = 3
	node3.next = &node4

	reverse(&node)

	listAsString := toString(node)

	if listAsString != "3, 1, 2, 4" {
		t.Errorf("TestReverse() expected '3, 1, 2, 4', found '" + listAsString + "'")
	}
}

func TestSortAsc(t *testing.T) {
	var node Node
	node.data = 4

	var node2 Node
	node2.data = 2
	node.next = &node2

	var node3 Node
	node3.data = 1
	node2.next = &node3

	var node4 Node
	node4.data = 3
	node3.next = &node4

	sortAsc(&node)

	listAsString := toString(node)

	if listAsString != "1, 2, 3, 4" {
		t.Errorf("listAsString() expected '1, 2, 3, 4', found '" + listAsString + "'")
	}
}

func TestSortDesc(t *testing.T) {
	var node Node
	node.data = 4

	var node2 Node
	node2.data = 2
	node.next = &node2

	var node3 Node
	node3.data = 1
	node2.next = &node3

	var node4 Node
	node4.data = 3
	node3.next = &node4

	sortDesc(&node)

	listAsString := toString(node)

	if listAsString != "4, 3, 2, 1" {
		t.Errorf("listAsString() expected '4, 3, 2, 1', found '" + listAsString + "'")
	}
}