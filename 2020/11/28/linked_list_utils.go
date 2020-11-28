package linked_list


import (
	"fmt"
	"strconv"
)

func add(node *Node, data int) {
	var nextNode Node
	nextNode.data = data

	node.next = &nextNode

	fmt.Println(toString(*node))
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