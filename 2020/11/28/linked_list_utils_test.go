package linked_list

import (
	//"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	var node Node
	node.data = 1

	add(&node, 2)

	listAsString := toString(node)

	if listAsString != "1, 2" {
		t.Errorf("listAsString() expected '1, 2', found '" + listAsString + "'")
	}
}