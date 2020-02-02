package gom

import ()

// NodeList objects are collections of nodes (live)
// https://developer.mozilla.org/en-US/docs/Web/API/NodeList
type NodeList interface {
	/* GETTERS & SETTERS */
	Length() int
	/* METHODS */
	append(node Node) Node    // private
	appendList(nodes ...Node) // private
	ForEach(func(i int, c Node))
	IndexOf(n Node) int
	Item(index int) Node
	set(index int, node Node) // private
	Values() []Node
}

var _ NodeList = &nodeList{}

type nodeList struct {
	list []Node
}

func newNodeList() *nodeList {
	return &nodeList{}
}

/*****************************************************
 **************** Getters & Setters ******************
 *****************************************************/

// Length method return the number of node in the list
func (nl *nodeList) Length() int {
	return len(nl.Values())
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/

func (nl *nodeList) append(node Node) Node {
	var child = node
	nl.list = append(nl.list, child)
	return child
}

func (nl *nodeList) appendList(nodes ...Node) {
	nl.list = append(nl.list, nodes...)
}

// ForEach apply the given function for each of
// the Node in the list.
func (nl *nodeList) ForEach(fn func(i int, c Node)) {
	for i, v := range nl.Values() {
		fn(i, v)
	}
}

// IndexOf method return the index of the
// searched node and return -1 if not found.
func (nl *nodeList) IndexOf(searched Node) int {
	for index, node := range nl.Values() {
		if same := node.IsSameNode(searched); same {
			return index
		}
	}

	return -1
}

// Item return a node from the Node list by index
func (nl *nodeList) Item(index int) Node {
	if index >= 0 && index < nl.Length() {
		return nl.Values()[index]
	}
	return nil
}

func (nl *nodeList) set(index int, node Node) {
	nl.list[index] = node
}

// Values method returns an iterator allowing to go
// through all values contained in this object.
// https://developer.mozilla.org/en-US/docs/Web/API/NodeList/values
func (nl *nodeList) Values() []Node {
	return nl.list
}
