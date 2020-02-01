package dom

// NodeList objects are collections of nodes (live)
type NodeList interface {
	append(...Node) NodeList
	ForEach(func(i int, c Node))
	IndexOf(node Node) int
	Item(index int) Node
	Length() int
	Values() []Node
}

var _ NodeList = &BasicNodeList{}

// BasicNodeList is the basic node list
type BasicNodeList []Node

func (nl *BasicNodeList) append(nodes ...Node) NodeList {
	var tmp = BasicNodeList(append(nl.Values(), nodes...))
	return &tmp
}

// ForEach apply the given function for each of
// the Node in the list.
func (nl *BasicNodeList) ForEach(fn func(i int, c Node)) {
	for i, v := range nl.Values() {
		fn(i, v)
	}
}

// IndexOf method return the index of the
// searched node and return -1 if not found.
func (nl *BasicNodeList) IndexOf(searched Node) int {
	for index, node := range nl.Values() {
		if same := searched.IsSameNode(node); same {
			return index
		}
	}

	return -1
}

// Item return a node from the Node list by index
func (nl *BasicNodeList) Item(index int) Node {
	return nl.Values()[index]
}

// Length method return the number of node in the list
func (nl *BasicNodeList) Length() int {
	return len(nl.Values())
}

// Values method returns an iterator allowing to go
// through all values contained in this object.
// https://developer.mozilla.org/en-US/docs/Web/API/NodeList/values
func (nl *BasicNodeList) Values() []Node {
	var tmp = *nl
	return tmp
}
