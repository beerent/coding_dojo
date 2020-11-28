package linked_list


import (
	"strconv"
)

type SortType int

const (
    ASC = iota
    DESC
)

func add(node *Node, data int) {
	var nextNode Node
	nextNode.data = data

	if node.next != nil {
		nextNode.next = node.next
	}

	node.next = &nextNode
}

func set(node *Node, data int) {
	var nextNode Node
	nextNode.data = data

	node.next = &nextNode
}

func remove(node *Node) {
	node.next = nil
}

func swap(nodeA, nodeB *Node) {
	var tempNode Node

	tempNode.data = nodeA.data
	nodeA.data = nodeB.data
	nodeB.data = tempNode.data
}

func count(node *Node) int {
	count := 0
	countNode := node
	for {
		if countNode == nil {
			break
		}

		count++
		countNode = countNode.next
	}

	return count
}

func sort(node *Node, sortType SortType) {
	switch sortType {
	case ASC:
	    sortAsc(node)
	    break;
	case DESC:
	    sortDesc(node)
	    break;
	}
}

func sortAsc(node *Node) {
	trackingNode := node
	currentNode := node

	for {
		if currentNode == nil {
			break
		}

		swapMade := false

		for {
			if currentNode.next == nil || currentNode.data <= currentNode.next.data {
				break
			}

			swap(currentNode, currentNode.next)
			swapMade = true
			currentNode = currentNode.next
		}

		if !swapMade {
			trackingNode = trackingNode.next
		}
		currentNode = trackingNode
	}
}

func sortDesc(node *Node) {
	sortAsc(node)
	reverse(node)
}

func reverse(node *Node) {
	iterations := count(node)
	for i := iterations; i > 0; i-- {
		currentNode := node
		for j := 1; j < i; j++ {
			swap(currentNode, currentNode.next)
			currentNode = currentNode.next
		}
	}
}

func intToString(i int) string {
	return strconv.Itoa(i)
}

func toString(list Node) string {
	var str string
	str = intToString(list.data)

	for {
		if list.next == nil {
			break
		}

		list = *list.next

		str += ", " + intToString(list.data)
	}

	return str
}