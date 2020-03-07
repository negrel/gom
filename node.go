package gom

import (
	e "github.com/negrel/gom/exception"
)

/* NOTE Node missing props & methods (OFFICIAL DOM) :
 * ** Props **
 * baseURI
 * baseURIObject
 * all obsolete or non-standardized props
 * ** Methods **
 * isDefaultNameSpace
 * lookupPrefix
 * all obsolete or non-standardized methods
 */

// Node is an interface and does not exist
// as node (abstract). It is used embedded
// by all nodes.
// (Document, DocumentType, Element, Text,
// and Comment).
// https://developer.mozilla.org/en-US/docs/Web/API/Node
// https://dom.spec.whatwg.org/#node
type Node interface {
	/* Private */
	apply(func(self Node))
	setParentElement(parent Element)
	setParentNode(parent Node)
	/* GETTERS & SETTERS (props) */
	ChildNodes() NodeList
	FirstChild() Node
	LastChild() Node
	NextSibling() Node
	NodeName() string
	NodeType() NodeType
	OwnerDocument() Document
	ParentNode() Node
	ParentElement() Element
	PreviousSibling() Node
	TextContent() string
	SetOwnerDocument(doc Document)
	SetTextContent(content string)
	/* METHODS */
	AppendChild(child Node) Node
	CloneNode(deep bool) Node
	CompareDocumentPosition(other Node) int
	Contains(other Node) bool
	GetRootNode() Node
	HasChildNodes() bool
	InsertBefore(new, reference Node) Node
	IsEqualNode(other Node) bool
	IsSameNode(other Node) bool
	Normalize()
	RemoveChild(child Node) (Node, e.Exception)
	ReplaceChild(newChild, oldChild Node) e.Exception
}

var _ Node = &node{}

type node struct {
	childNodes    NodeList
	isConnected   bool
	parentNode    Node
	parentElement Element
	document      Document
}

// The CompareDocumentPosition return values
// are a bitmask with the following values.
const (
	DocumentPositionDisconnected = 1 << iota
	DocumentPositionPreceding
	DocumentPositionFollowing
	DocumentPositionContains
	DocumentPositionContainedBy
	DocumentPositionImplementationSpecific
)

// NodeType is the type of a Node.
type NodeType = uint32

// Node type list
const (
	_ = iota
	ElementNode
	AttributeNode
	TextNode
	CommentNode
	DocumentNode
	DocumentTypeNode
)

func newNode() Node {
	return &node{
		childNodes:    newNodeList(),
		isConnected:   false,
		parentNode:    nil,
		parentElement: nil,
		document:      nil,
	}
}

// apply the function to the node and all is descendant.
func (n *node) apply(fn func(node Node)) {
	// apply to the node itself
	fn(n)

	// apply to all the children
	n.childNodes.ForEach(func(_ int, child Node) {
		child.apply(fn)
	})
}

func (n *node) setParentElement(parent Element) {
	n.parentElement = parent
}

func (n *node) setParentNode(parent Node) {
	n.parentNode = parent
}

/*****************************************************
 **************** Getters & Setters ******************
 *****************************************************/
// ANCHOR Getters & Setters

// ChildNodes return a node list containing
func (n *node) ChildNodes() NodeList {
	return n.childNodes
}

// FirstChild method return the first child of the
// current node.
func (n *node) FirstChild() Node {
	return n.childNodes.Item(0)
}

// LastChild method return the last child of the
// current node.
func (n *node) LastChild() Node {
	lastIndex := n.childNodes.Length() - 1

	return n.childNodes.Item(lastIndex)
}

// NextSibling - method return the next sibling
// of the current node.
func (n *node) NextSibling() Node {
	index := n.parentNode.ChildNodes().IndexOf(n)

	return n.parentNode.ChildNodes().Item(index + 1)
}

// NodeName method return a DOMString containing
// the name of the node type.
func (n *node) NodeName() string {
	// https://dom.spec.whatwg.org/#dom-node-nodename
	return ""
}

// NodeType return an integer that identifies
// what the node is.
// https://developer.mozilla.org/en-US/docs/Web/API/Node/nodeType
func (n *node) NodeType() NodeType {
	return 0
}

// OwnerDocument method return the top-level document
// object of the node.
func (n *node) OwnerDocument() Document {
	// TODO func (n *node) OwnerDocument() Document
	// https://developer.mozilla.org/en-US/docs/Web/API/Node/ownerDocument
	return n.document
}

// ParentNode method returns a node that is the
// parent of this node.
func (n *node) ParentNode() Node {
	return n.parentNode
}

// ParentElement method returns the parent node.
func (n *node) ParentElement() Element {
	return n.parentElement
}

// PreviousSibling method return the previous
// sibling of the current node.
func (n *node) PreviousSibling() Node {
	index := n.parentNode.ChildNodes().IndexOf(n)

	return n.parentNode.ChildNodes().Item(index - 1)
}

// TextContent methode return the textual content
// of an element and all its descendants.
func (n *node) TextContent() string {
	// TODO func (n *node) TextContent() string
	// https://dom.spec.whatwg.org/#dom-node-textcontent
	return ""
}

func (n *node) SetOwnerDocument(doc Document) {
	n.document = doc
}

// SetTextContent methode set the textual content
// of an element and all its descendants.
func (n *node) SetTextContent(content string) {
	// TODO func (n *node) SetTextContent(string)
	// https://dom.spec.whatwg.org/#dom-node-textcontent
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// AppendChild methods adds the specified childNode
// argument as the last child to the current node.
// https://developer.mozilla.org/en-US/docs/Web/API/Node/appendChild
func (n *node) AppendChild(child Node) Node {
	// Child already exist in the tree
	// So move it from its current position
	if n.GetRootNode().Contains(child) {
		// TODO
	}

	// Appending the child.
	child = n.childNodes.append(child)
	child.setParentNode(n)

	return child
}

// CloneNode method return a duplicate of the node on
// which this method was called. Set the deep argument
// to true if you want the childs to be cloned
// https://developer.mozilla.org/en-US/docs/Web/API/Node/cloneNode
func (n *node) CloneNode(deep bool) Node {
	clone := newNode()

	// Copy node children
	if deep {
		for _, child := range n.childNodes.Values() {
			clone.AppendChild(child.CloneNode(true))
		}
	}

	return clone
}

// CompareDocumentPosition method compares the position
// of the given node against another node in any document.
// https://developer.mozilla.org/en-US/docs/Web/API/Node/compareDocumentPosition
func (n *node) CompareDocumentPosition(other Node) int {
	// TODO func (n *node) CompareDocumentPosition(other Node) int
	// https://dom.spec.whatwg.org/#dom-node-comparedocumentposition
	return 0
}

// Contains method returns a Boolean value indicating
// whether a node is a descendant of a given node.
// https://developer.mozilla.org/en-US/docs/Web/API/Node/contains
func (n *node) Contains(other Node) (contain bool) {
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

// GetRootNode method of the node interface returns
// the context object's root.
// https://developer.mozilla.org/en-US/docs/Web/API/Node/getRootNode
func (n *node) GetRootNode() Node {
	// If parent exist, get parent root
	if n.parentNode != nil {
		return n.parentNode.GetRootNode()
	}

	// No parent, this node is the root node.
	return n
}

// HasChildNodes method returns a bool value indicating
// whether the given node has child nodes or not.
// https://developer.mozilla.org/en-US/docs/Web/API/Node/hasChildNodes
func (n *node) HasChildNodes() bool {
	return n.childNodes.Length() > 0
}

// InsertBefore method inserts a node before a reference
// node as a child (or append the node if the reference
// node is not a direct child) of the node on which this
// method was called.
// Return the inserted node.
// https://developer.mozilla.org/en-US/docs/Web/API/Node/insertBefore
func (n *node) InsertBefore(new Node, reference Node) Node {
	// Child already exist in the tree
	// So move it from its current position
	if n.GetRootNode().Contains(new) {
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
	n.childNodes.appendList(beforeNew...)
	n.childNodes.append(new)
	n.childNodes.appendList(afterNew...)

	return n.childNodes.Item(index)
}

// IsEqualNode method return whether two nodes are equal.
// https://developer.mozilla.org/en-US/docs/Web/API/Node/isEqualNode
func (n *node) IsEqualNode(other Node) bool {
	if other == nil {
		goto notEqual
	}

	// Check nodeType
	if n.NodeType() != other.NodeType() {
		goto notEqual
	}

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

	return true

notEqual:
	return false
}

// IsSameNode method for node objects tests whether two
// nodes are the same (reference).
// https://developer.mozilla.org/en-US/docs/Web/API/Node/isSameNode
func (n *node) IsSameNode(other Node) bool {
	return n == other
}

// Normalize method clean up all the text nodes under
// this element (merge adjacent, remove empty).
// https://developer.mozilla.org/en-US/docs/Web/API/Node/normalize
func (n *node) Normalize() {
	// TODO func (n *node) Normalize()
	// https://dom.spec.whatwg.org/#dom-node-normalize
}

// RemoveChild method removes a child node from the DOM
// and returns the removed node.
// https://developer.mozilla.org/en-US/docs/Web/API/Node/removeChild
func (n *node) RemoveChild(child Node) (Node, e.Exception) {
	childIndex := n.childNodes.IndexOf(child)

	// Child not found.
	if err := childIndex == -1; err {
		return child,
			e.New(e.NotFoundError, "The node to be removed is not a child of this node.")
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

	// Appending slice of nodes child before and after
	// the child to remove.
	n.childNodes.appendList(beforeChild...)
	n.childNodes.appendList(afterChild...)

	// Removing parent of the child
	child.setParentNode(nil)

	return child, nil
}

// ReplaceChild method replaces a child node within the
// given (parent) node.
// https://developer.mozilla.org/en-US/docs/Web/API/Node/replaceChild
func (n *node) ReplaceChild(newChild, oldChild Node) e.Exception {

	index := n.childNodes.IndexOf(oldChild)

	if index == -1 {
		return e.New(e.NotFoundError, "The node to be replaced is not a child of this node.")
	}

	n.childNodes.set(index, newChild)

	return nil
}
