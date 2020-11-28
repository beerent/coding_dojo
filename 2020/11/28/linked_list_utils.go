package linked_list


import (
	"strconv"
)

func add(node *Node, data int) {
	var nextNode Node
	nextNode.data = data

	if node.next != nil {
		nextNode.next = node.next
	}

	node.next = &nextNode
}

//func reverse(list LinkedList) LinkedList {
//	return list
//}

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