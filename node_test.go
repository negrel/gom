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

	// Check that node contains the child
	if contain := node.Contains(child); !contain {
		t.Log("Node must contain the appended child.")
		t.Logf("Node number of child : %v", node.ChildNodes().Length())
		t.Fail()
	}

	// Check that child parent is node
	if isParent := child.ParentNode().IsSameNode(node); !isParent {
		t.Log("Node must be the parent of the appended child. (pointer must be the same)")
		t.Logf("Child parent node pointer address : %p", child.ParentNode())
		t.Logf("Node pointer address              : %p", node)
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

	// Checking that clone is equal to node
	if equal := clone.IsEqualNode(node); !equal {
		t.Log("Clone must be equal to node. (deep clone)")
		t.Logf("Clone child node : %v", clone.ChildNodes())
		t.Logf("Node child node  : %v", node.ChildNodes())
		t.Fail()
	}

	// Checking that clone doesn't point to node
	if same := clone.IsSameNode(node); same {
		t.Log("Clone must not point to the same reference than node.")
		t.Logf("Clone pointers address : %p", clone)
		t.Logf("Node pointers address  : %p", node)
		t.Fail()
	}
}

func TestCompareDocumentPosition(t *testing.T) {
	// TODO func TestCompareDocumentPosition(t *testing.T)
}

func TestContains(t *testing.T) {
	node := newNode()
	child1 := newNode()
	child2 := child1.CloneNode(false)

	child1 = node.AppendChild(child1)

	// Checking that child1 is equal child2
	if equal := child1.IsEqualNode(child2); !equal {
		t.Log("Child2 must be a clone of child1.")
		t.Fail()
	}

	// Checking that node doesn't contain child2
	if contain := node.Contains(child2); contain {
		t.Log("Node must not contain child2. (pointer must be different)")
		t.Logf("Node childrens pointer address : %v", node.ChildNodes().Values())
		t.Logf("Child2 pointer address         : %p", child2)
		t.Fail()
	}

	child2 = child1.AppendChild(child2)

	// Checking that child1 contain child2
	if contain := child1.Contains(child2); !contain {
		t.Log("Child1 must contain child2 (direct child). (pointer must be the same)")
		t.Logf("Child1 childrens pointer address : %v", child1.ChildNodes().Values())
		t.Logf("Child2 pointer address           : %p", child2)
		t.Fail()
	}

	// Checking that node contain child2 (child of child1)
	if contain := node.Contains(child2); !contain {
		t.Log("Node must contain child2 (direct child of child1). (pointer must be the same)")
		t.Logf("Child1 childrens pointer address : %v", child1.ChildNodes().Values())
		t.Logf("Child2 pointer address           : %p", child2)
		t.Fail()
	}
}

func TestGetRootNode(t *testing.T) {
	node := newNode()
	child1 := newNode()
	child2 := newNode()

	child1 = node.AppendChild(child1)

	// Checking that node root is node
	if same := node.GetRootNode().IsSameNode(node); !same {
		t.Log("Node root node must be node itself. (pointer must be the same)")
		t.Logf("Node root pointer address : %p", node.GetRootNode())
		t.Logf("Node pointer address      : %p", node)
		t.Fail()
	}

	// Checking that child root is node
	if same := child1.GetRootNode().IsSameNode(node); !same {
		t.Log("Node must be the root node of child1. (pointer must be the same)")
		t.Logf("Child1 root pointer address : %p", child1.GetRootNode())
		t.Logf("Node pointer address        : %p", node)
		t.Fail()
	}

	// Checking that child2 root is not node
	if same := child2.GetRootNode().IsSameNode(node); same {
		t.Log("Node must not be the root node of child2. (pointer must not be the same)")
		t.Logf("Child2 root pointer address : %p", child2.GetRootNode())
		t.Logf("Node pointer address        : %p", node)
		t.Fail()
	}
}

func TestHasChildNodes(t *testing.T) {
	node := newNode()
	child := newNode()
	clone := node.CloneNode(false)

	node.AppendChild(child)

	// Checking that node has child.
	if hasChild := node.HasChildNodes(); !hasChild {
		t.Log("Node have child. (node.HasChildNodes must return true)")
		t.Fail()
	}

	// Checking that clone has not child.
	if hasChild := clone.HasChildNodes(); hasChild {
		t.Log("Clone haven't any child. (clone.HasChildNodes must return false)")
		t.Fail()
	}
}
