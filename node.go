package dom

import (
	"errors"
)

// Node interface represent methods of any DOM object
// Strongly inspired of the Node object :
// https://developer.mozilla.org/en-US/docs/Web/API/Node
type Node interface {
	/* ACCESSORS */
	ChildNodes() NodeList    // Read-only
	FirstChild() Node        // Read-only
	LastChild() Node         // Read-only
	NextSibling() Node       // Read-only
	NodeName() string        // Read-only
	NodeType() int           // Read-only
	OwnerDocument() Document // Read-only
	ParentNode() Node        // Read-only
	ParentElement() Element  // Read-only
	PreviousSibling() Node   // Read-only
	TextContent() string
	setNodeName(string)
	setNodeType(int)
	setParentElement(parent Element)
	setParentNode(parent Node)
	SetTextContent(string)
	/* METHODS */
	AppendChild(child Node) Node
	CloneNode(deep bool) Node
	CompareDocumentPosition(other Node) int
	Contains(child Node) bool
	GetRootNodes() Node
	HasChildNodes() bool
	InsertBefore(new Node, reference Node) Node
	IsEqualNode(Node) bool
	IsSameNode(Node) bool
	Normalize()
	RemoveChild(Node) (Node, error)
	ReplaceChild(newChild, oldChild Node) (Node, error)
}

// The CompaerDocumentPosition return values
// are a bitmask with the following values.
const (
	DocumentPositionDisconnected = 1 << iota
	DocumentPositionPreceding
	DocumentPositionFollowing
	DocumentPositionContains
	DocumentPositionContainedBy
	DocumentPositionImplementationSpecific
)

// Node type list
const (
	_ = iota
	ElementNode
	TextNode
	CommentNode
	DocumentNode
)

var _ Node = &BasicNode{}

// BasicNode is the basic struct for Node object
type BasicNode struct {
	childNodes    NodeList
	isConnected   bool
	nodeName      string
	nodeType      int
	parentNode    Node
	parentElement Element
}

func newNode() *BasicNode {
	return &BasicNode{
		childNodes:    &BasicNodeList{},
		isConnected:   false,
		nodeType:      0,
		parentNode:    nil,
		parentElement: nil,
	}
}

// ChildNodes return a node list containing
func (n *BasicNode) ChildNodes() NodeList {
	return n.childNodes
}

// FirstChild method return the first child of the
// current node.
func (n *BasicNode) FirstChild() Node {
	return n.childNodes.Item(0)
}

// LastChild method return the last child of the
// current node.
func (n *BasicNode) LastChild() Node {
	lastIndex := n.childNodes.Length() - 1

	return n.childNodes.Item(lastIndex)
}

// NextSibling - method return the next sibling
// of the current node.
func (n *BasicNode) NextSibling() Node {
	index := n.parentNode.ChildNodes().IndexOf(n)

	return n.parentNode.ChildNodes().Item(index + 1)
}

// NodeName method return a DOMString containing
// the name of the node type.
func (n *BasicNode) NodeName() string {
	// TODO func (n *BasicNode) NodeName() string
	return n.nodeName
}

// NodeType return an integer that identifies
// what the node is.
// https://developer.mozilla.org/en-US/docs/Web/API/Node/nodeType
func (n *BasicNode) NodeType() int {
	// TODO func (n *BasicNode) NodeName() string
	return n.nodeType
}

// OwnerDocument method return the top-level document
// object of the node.
func (n *BasicNode) OwnerDocument() Document {
	// TODO func (n *BasicNode) OwnerDocument() Document
	// https://developer.mozilla.org/en-US/docs/Web/API/Node/ownerDocument
	return n
}

// ParentNode method returns a Node that is the
// parent of this node.
func (n *BasicNode) ParentNode() Node {
	return n.parentNode
}

// ParentElement method returns the parent node.
func (n *BasicNode) ParentElement() Element {
	return n.parentElement
}

// PreviousSibling method return the previous
// sibling of the current node.
func (n *BasicNode) PreviousSibling() Node {
	index := n.parentNode.ChildNodes().IndexOf(n)

	return n.parentNode.ChildNodes().Item(index - 1)
}

// TextContent methode return the textual content
// of an element and all its descendants.
func (n *BasicNode) TextContent() string {
	// TODO func (n *BasicNode) TextContent() string
	// https://developer.mozilla.org/en-US/docs/Web/API/Node/textContent
	return ""
}

func (n *BasicNode) setNodeName(name string) {
	n.nodeName = name
}

func (n *BasicNode) setNodeType(nType int) {
	n.nodeType = nType
}

func (n *BasicNode) setParentElement(parent Element) {
	n.parentElement = parent
}

func (n *BasicNode) setParentNode(parent Node) {
	n.parentNode = parent
}

// SetTextContent methode set the textual content
// of an element and all its descendants.
func (n *BasicNode) SetTextContent(content string) {
	// TODO func (n *BasicNode) SetTextContent(string)
	// https://developer.mozilla.org/en-US/docs/Web/API/Node/textContent
}

// AppendChild methods adds the specified childNode
// argument as the last child to the current node.
// https://developer.mozilla.org/en-US/docs/Web/API/Node/appendChild
func (n *BasicNode) AppendChild(child Node) Node {
	// Child already exist in the document
	// So move it from its current position
	if n.OwnerDocument().Contains(child) {
		// TODO
	}

	// Appending the child.
	n.childNodes = n.childNodes.append(child)
	child.setParentNode(n)

	return child
}

// CloneNode method return a duplicate of the node on
// which this method was called.
// https://developer.mozilla.org/en-US/docs/Web/API/Node/cloneNode
func (n *BasicNode) CloneNode(deep bool) Node {
	clone := newNode()
	clone.setNodeName(n.NodeName())
	clone.setNodeType(n.NodeType())

	// Copy node children
	if deep {
		// TODO copy node children
		for _, child := range n.childNodes.Values() {
			clone.AppendChild(child.CloneNode(true))
		}
	}

	return clone
}

// CompareDocumentPosition method compares the position
// of the given node against another node in any document.
// https://developer.mozilla.org/en-US/docs/Web/API/Node/compareDocumentPosition
func (n *BasicNode) CompareDocumentPosition(other Node) int {
	// TODO func (n *BasicNode) CompareDocumentPosition(other Node) int
	return 0
}

// Contains method returns a Boolean value indicating
// whether a node is a descendant of a given node.
// https://developer.mozilla.org/en-US/docs/Web/API/Node/contains
func (n *BasicNode) Contains(other Node) (contain bool) {
	// Checking first gen child
	for _, childNode := range n.childNodes.Values() {
		if childNode.IsSameNode(other) {
			contain = true
			break
		}
	}

	// Checking child of child and so on
	if !contain {
		for _, childNode := range n.childNodes.Values() {
			if childNode.Contains(other) {
				contain = true
				break
			}
		}
	}

	return contain
}

// GetRootNodes method of the Node interface returns
// the context object's root.
// https://developer.mozilla.org/en-US/docs/Web/API/Node/getRootNode
func (n *BasicNode) GetRootNodes() Node {
	// If parent exist, get parent root
	if n.parentNode != nil {
		return n.parentNode.GetRootNodes()
	}

	// No parent, this parent is the root node.
	return n
}

// HasChildNodes method returns a bool value indicating
// whether the given Node has child nodes or not.
// https://developer.mozilla.org/en-US/docs/Web/API/Node/hasChildNodes
func (n *BasicNode) HasChildNodes() bool {
	return n.childNodes.Length() > 0
}

// InsertBefore method inserts a node before a reference
// node as a child of the node on which this method was called.
// Return the inserted node.
// https://developer.mozilla.org/en-US/docs/Web/API/Node/insertBefore
func (n *BasicNode) InsertBefore(new Node, reference Node) Node {
	// Child already exist in the document
	// So move it from its current position
	if n.OwnerDocument().Contains(new) {
		// TODO InsertBefore child already exist
	}

	// Reference node index
	index := n.childNodes.IndexOf(reference)

	// Reference is not found so we append the new node
	if index == -1 {
		return n.AppendChild(new)
	}

	// Reference node is found let's insert the new node
	// slice of child before the new node
	beforeNew := make([]Node, 0)
	// slice of child after the new node
	afterNew := make([]Node, 0)

	if index > 0 {
		beforeNew = n.childNodes.Values()[:index-1]
		afterNew = n.childNodes.Values()[index:]
	}

	// beforeNew + new + afterNew
	n.childNodes.append(beforeNew...)
	n.childNodes.append(new)
	n.childNodes.append(afterNew...)

	return n.childNodes.Item(index)
}

// IsEqualNode method return whether two nodes are equal.
// https://developer.mozilla.org/en-US/docs/Web/API/Node/isEqualNode
func (n *BasicNode) IsEqualNode(other Node) bool {
	// Checking the list of childrens length
	if n.childNodes.Length() != other.ChildNodes().Length() {
		goto notEqual
	}

	// Check children
	for i, child := range n.childNodes.Values() {
		if otherChild := other.ChildNodes().Item(i); otherChild != nil {
			if !otherChild.IsEqualNode(child) {
				goto notEqual
			}
		}
	}

	// Check nodeType
	if n.NodeType() != other.NodeType() {
		goto notEqual
	}

	return true

notEqual:
	return false
}

// IsSameNode method for Node objects tests whether two
// nodes are the same (reference).
// https://developer.mozilla.org/en-US/docs/Web/API/Node/isSameNode
func (n *BasicNode) IsSameNode(other Node) bool {
	return n == other
}

// Normalize method clean up all the text nodes under
// this element (merge adjacent, remove empty).
// https://developer.mozilla.org/en-US/docs/Web/API/Node/normalize
func (n *BasicNode) Normalize() {
	// TODO func (n *BasicNode) Normalize()
}

// RemoveChild method removes a child node from the DOM
// and returns the removed node.
// https://developer.mozilla.org/en-US/docs/Web/API/Node/removeChild
func (n *BasicNode) RemoveChild(child Node) (Node, error) {
	childIndex := n.childNodes.IndexOf(child)

	// Child not found.
	if err := childIndex == -1; err {
		return child, errors.New("The child to remove must be a direct child")
	}

	// Slice of all the child before the child to remove
	beforeChild := make([]Node, 0)
	// Slice of all the child after the child to remove
	afterChild := make([]Node, 0)

	// If child have previousSibling there is child
	// before the child to remove
	if childIndex > 0 {
		beforeChild = n.childNodes.Values()[:childIndex-1]
	}
	// If child have nextSibling there is child
	// after the child to remove
	if n.childNodes.Length() > (childIndex + 1) {
		afterChild = n.childNodes.Values()[childIndex+1:]
	}

	// Child is removed
	n.childNodes.append(beforeChild...)
	n.childNodes.append(afterChild...)

	return child, nil
}

// ReplaceChild method replaces a child node within the
// given (parent) node.
// https://developer.mozilla.org/en-US/docs/Web/API/Node/replaceChild
func (n *BasicNode) ReplaceChild(newChild, oldChild Node) (Node, error) {
	newChild = n.InsertBefore(newChild, oldChild)
	_, err := n.RemoveChild(oldChild)

	return newChild, err
}
