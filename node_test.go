package dom

import (
	"testing"
)

func TestAppendChild(t *testing.T) {
	// Creating node & child
	node := newNode()
	child := node.CloneNode(false)

	// Appending the child
	child = node.AppendChild(child)

	// If node doesn't contain child, test failed
	if contain := node.Contains(child); !contain {
		t.Log("Node must contain the appended child.")
		t.Logf("Node number of child : %v", node.ChildNodes().Length())
		t.Fail()
	}

	// If child parent node is not the same as node, test failed
	if isParent := child.ParentNode().IsSameNode(node); !isParent {
		t.Log("Node must be the parent of the appended child.")
		t.Fail()
	}
}
