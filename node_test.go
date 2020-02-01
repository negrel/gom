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

func TestCloneNode(t *testing.T) {
	node := newNode()
	child := newNode()
	child2 := newNode()

	node.AppendChild(child2)
	child = node.AppendChild(child)

	// Clone the node but not his childs
	clone := node.CloneNode(false)

	// Clone must not be equal (different childs)
	if equal := clone.IsEqualNode(node); equal {
		t.Log("Clone must not be equal to node. (different child)")
		t.Logf("Clone child node : %v", clone.ChildNodes())
		t.Logf("Node child node  : %v", node.ChildNodes())
		t.Fail()
	}

	clone = node.CloneNode(true)

	// Clone must be equal (deep clone)
	if equal := clone.IsEqualNode(node); !equal {
		t.Log("Clone must be equal to node. (deep clone)")
		t.Logf("Clone child node : %v", clone.ChildNodes())
		t.Logf("Node child node  : %v", node.ChildNodes())
		t.Fail()
	}

	// Clone must not be the same
	if same := clone.IsSameNode(node); same {
		t.Log("Clone must not point to the same reference than node.")
		t.Logf("Clone pointers address : %p", clone)
		t.Logf("Node pointers address  : %p", node)
		t.Fail()
	}
}
