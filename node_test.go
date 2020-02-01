package dom

import (
	"testing"
)

func TestAppendChild(t *testing.T) {
	node := newNode()
	child := node.CloneNode(false)

	child = node.AppendChild(child)

	if contain := node.Contains(child); !contain {
		t.Log("Node must contain the appended child.")
		t.Logf("Node number of child : %v", node.ChildNodes().Length())
		t.Fail()
	}

	if isParent := child.ParentNode().IsSameNode(node); !isParent {
		t.Log("Node must be the parent of the appended child.")
		t.Fail()
	}
}
