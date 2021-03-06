package gom

import (
	"sort"

	e "github.com/negrel/gom/exception"
)

// NamedNodeMap interface represents a collection
// of Attr objects.
// https://developer.mozilla.org/en-US/docs/Web/API/NamedNodeMap
// https://dom.spec.whatwg.org/#interface-namednodemap
type NamedNodeMap interface {
	/* Private */
	getNamedItem(string) (Attr, bool)
	/* GETTERS & SETTERS (props) */
	Length() int
	/* METHODS */
	GetNamedItem(string) Attr
	Item(int) Attr
	SetNamedItem(Attr)
	RemoveNamedItem(string) (Attr, e.Exception)
	Values() []Attr // Not part of DOM specification
}

var _ NamedNodeMap = &namedNodeMap{}

type namedNodeMap struct {
	dict map[string]Attr
}

func (n *namedNodeMap) getNamedItem(name string) (Attr, bool) {
	attr, ok := n.dict[name]
	return attr, ok
}

/*****************************************************
 **************** Getters & Setters ******************
 *****************************************************/
// ANCHOR Getters & Setters

func (n *namedNodeMap) Length() int {
	return len(n.dict)
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// GetNamedItem return the attribute corresponding to
// the given name.
func (n *namedNodeMap) GetNamedItem(name string) Attr {
	return n.dict[name]
}

// Item returns the Attr at the given index, or null if
// the index is higher or equal to the number of nodes.
// (Slower than GetNamedItem)
func (n *namedNodeMap) Item(index int) Attr {
	return n.Values()[index]
}

// SetNamedItem Replaces, or adds, the Attr identified
// in the map by the given name.
func (n *namedNodeMap) SetNamedItem(attr Attr) {
	// Set the new attribute value
	n.dict[attr.Name()] = attr
}

// RemoveNamedItem remove the specified attribute.
func (n *namedNodeMap) RemoveNamedItem(name string) (Attr, e.Exception) {
	attr := n.dict[name]

	// Check if attribute exist
	if attr == nil {
		return nil, e.New(e.NotFoundError, "The attr to be removed is not part of this element")
	}

	delete(n.dict, name)

	return attr, nil
}

// Values return an iterable array of attributes.
func (n *namedNodeMap) Values() []Attr {
	arr := make([]Attr, n.Length())

	for _, value := range n.dict {
		arr = append(arr, value)
	}

	// Alphabetical sort
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Name() < arr[j].Name()
	})

	return arr
}
